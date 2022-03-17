package http

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"

	"github.com/ansidev/fiber-starter-project/pkg/log"
	"github.com/ansidev/fiber-starter-project/post/service"
	"go.uber.org/zap"
)

func registerRoutes(app *fiber.App, postController *PostController) {
	v1 := app.Group("/post/v1")

	v1.Get("/posts/:id", postController.GetPost)
}

func NewPostController(app *fiber.App, postService service.IPostService) {
	controller := &PostController{postService}
	registerRoutes(app, controller)
}

type PostController struct {
	postService service.IPostService
}

func (ctrl *PostController) GetPost(ctx *fiber.Ctx) error {
	postIdParam := ctx.Params("id")

	postId, err := strconv.ParseInt(postIdParam, 10, 64)

	if err != nil {
		log.Errorz("Invalid post id", zap.String("post_id", postIdParam))
	}

	post, err := ctrl.postService.GetByID(postId)

	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(map[string]string{
				"error": err.Error(),
			})
	}

	return ctx.Status(http.StatusOK).JSON(post)
}
