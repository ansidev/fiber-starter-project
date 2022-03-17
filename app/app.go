package main

import (
	"database/sql"
	"fmt"
	authorController "github.com/ansidev/fiber-starter-project/author/controller/http"
	"github.com/ansidev/fiber-starter-project/config"
	"github.com/ansidev/fiber-starter-project/constant"
	gormPkg "github.com/ansidev/fiber-starter-project/pkg/gorm"
	"github.com/ansidev/fiber-starter-project/pkg/log"
	postController "github.com/ansidev/fiber-starter-project/post/controller/http"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	appEnv string
	sqlDb  *sql.DB
	gormDb *gorm.DB
)

func init() {
	log.InitLogger("console")

	appEnv = os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = constant.DefaultProdEnv
	}

	if appEnv == constant.DefaultProdEnv {
		config.LoadConfig("/app", constant.DefaultProdConfig, &config.AppConfig)
	} else {
		config.LoadConfig(".", constant.DefaultDevConfig, &config.AppConfig)
	}
}

func main() {
	// Flush log buffer if necessary
	defer log.Sync()

	app := fiber.New()

	// Default route
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(map[string]string{
			"name":        constant.AppName,
			"version":     constant.AppVersion,
			"releaseDate": constant.AppReleaseDate,
		})
	})

	initInfrastructureServices()
	initControllers(app)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(initAddress()); err != nil {
			log.Fatal(err)
		}
	}()

	// Create channel to signify a signal being sent
	exit := make(chan os.Signal, 1)
	// When an interrupt or termination signal is sent, notify the channel
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	// Block the main thread until an interrupt is received
	<-exit
	log.Info("Gracefully shutting down...")
	_ = app.Shutdown()

	_ = sqlDb.Close()
	log.Info("Closed database channel")

	log.Info("Server exiting")
}

func initInfrastructureServices() {
	sqlDb = InitSqlClient(config.AppConfig.SqlDbConfig)
	dialector := postgres.New(postgres.Config{
		Conn:                 sqlDb,
		PreferSimpleProtocol: true,
	})
	gormDb = gormPkg.InitGormDb(dialector)
}

func initControllers(app *fiber.App) {
	authorService := InitAuthorService(gormDb)
	authorController.NewAuthorController(app, authorService)

	postService := InitPostService(gormDb)
	postController.NewPostController(app, postService)
}

func initAddress() string {
	return fmt.Sprintf("%s:%d", config.AppConfig.Host, config.AppConfig.Port)
}
