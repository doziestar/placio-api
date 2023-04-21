package controller

import (
	"github.com/gofiber/fiber/v2"
	_ "placio-app/Dto"
	_ "placio-app/models"
	"placio-app/service"
)

type MediaController struct {
	mediaService service.MediaService
}

func NewMediaController(mediaService service.MediaService) *MediaController {
	return &MediaController{mediaService: mediaService}
}

func (mc *MediaController) RegisterRoutes(router fiber.Router) {
	mediaRouter := router.Group("/media")
	{
		//mediaRouter.Get("/", mc.getAllMedia)
		mediaRouter.Get("/:id", mc.getMedia)
		//mediaRouter.Post("/", mc.createMedia)
		//mediaRouter.Put("/:id", mc.updateMedia)
		mediaRouter.Delete("/:id", mc.deleteMedia)
	}
}

// @Summary Upload media
// @Description Upload a media file (image, gif, or video) for a post
// @Tags Media
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Media file"
// @Param postID formData string true "Post ID"
// @Success 201 {object} models.Media "Successfully uploaded media"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/media/ [post]
func (c *MediaController) uploadMedia(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	postID := ctx.FormValue("postID")
	if postID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Post ID is required",
		})
	}

	uploadedMedia, err := c.mediaService.UploadMedia(file, postID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(uploadedMedia)
}

// @Summary Get media
// @Description Retrieve media by its ID
// @Tags Media
// @Produce json
// @Param mediaID path string true "Media ID"
// @Success 200 {object} models.Media "Successfully retrieved media"
// @Failure 404 {object} Dto.ErrorDTO "Media Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/media/{mediaID} [get]
func (c *MediaController) getMedia(ctx *fiber.Ctx) error {
	mediaID := ctx.Params("mediaID")

	media, err := c.mediaService.GetMedia(mediaID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	if media == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Media Not Found",
		})
	}

	return ctx.JSON(media)
}

// @Summary Delete media
// @Description Delete media by its ID
// @Tags Media
// @Produce json
// @Param mediaID path string true "Media ID"
// @Success 204 "Successfully deleted media"
// @Failure 404 {object} Dto.ErrorDTO "Media Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/media/{mediaID} [delete]
func (c *MediaController) deleteMedia(ctx *fiber.Ctx) error {
	mediaID := ctx.Params("mediaID")

	err := c.mediaService.DeleteMedia(mediaID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusNoContent).SendString("")
}

// @Summary List all media for a post
// @Description Retrieve all media for the specified post
// @Tags Media
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {array} models.Media "Successfully retrieved media"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/media/post/{postID} [get]
func (c *MediaController) listMedia(ctx *fiber.Ctx) error {
	postID := ctx.Params("postID")
	mediaList, err := c.mediaService.ListMedia(postID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.JSON(mediaList)
}