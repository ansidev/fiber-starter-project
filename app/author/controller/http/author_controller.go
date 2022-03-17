package http

import (
	"net/http"
	"strconv"

	"github.com/ansidev/fiber-starter-project/author/service"
	"github.com/ansidev/fiber-starter-project/pkg/log"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func registerRoutes(app *fiber.App, authorController *AuthorController) {
	v1 := app.Group("/author/v1")

	v1.Get("/authors/:id", authorController.GetAuthor)
}

func NewAuthorController(app *fiber.App, authorService service.IAuthorService) {
	controller := &AuthorController{authorService}
	registerRoutes(app, controller)
}

type AuthorController struct {
	authorService service.IAuthorService
}

func (ctrl *AuthorController) GetAuthor(ctx *fiber.Ctx) error {
	authorIdParam := ctx.Params("id")

	authorId, err := strconv.ParseInt(authorIdParam, 10, 64)

	if err != nil {
		log.Errorz("Invalid author id", zap.String("author_id", authorIdParam))
	}

	author, err := ctrl.authorService.GetByID(authorId)

	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(map[string]string{
				"error": err.Error(),
			})
	}

	return ctx.Status(http.StatusOK).JSON(author)
}
