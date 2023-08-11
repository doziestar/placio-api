package media

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/utility"
)

type MediaController struct {
	mediaService MediaService
}

func NewMediaController(mediaService MediaService) *MediaController {
	return &MediaController{mediaService: mediaService}
}

func (mc *MediaController) RegisterRoutes(router *gin.RouterGroup) {
	mediaRouter := router.Group("/media")
	{
		//mediaRouter.Get("/", mc.getAllMedia)
		//mediaRouter.GET("/:id", utility.Use(mc.getMedia))
		mediaRouter.POST("/", utility.Use(mc.uploadMedia))
		//mediaRouter.Put("/:id", mc.updateMedia)
		//mediaRouter.POST("/:id", utility.Use(mc.deleteMedia))
	}
}

// @Summary Upload media
// @Description Upload a media file (image, gif, or video) for a post
// @Tags Media
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Media file"
// @Param postID formData string true "Post ID"
// @Success 201 {object} ent.Media "Successfully uploaded media"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/media/ [post]
func (mc *MediaController) uploadMedia(ctx *gin.Context) error {

	form, err := ctx.MultipartForm()
	if err != nil {

		return err
	}
	files, ok := form.File["files"]
	if !ok || len(files) == 0 {

		return fmt.Errorf("file is required")
	}

	uploadedMedia, err := mc.mediaService.UploadFiles(ctx, files)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusCreated, uploadedMedia)
	return nil
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
//func (c *MediaController) getMedia(ctx *gin.Context) error {
//	mediaID := ctx.Param("mediaID")
//
//	media, err := c.mediaService.GetMedia(mediaID)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"error": "Internal Server Error",
//		})
//		return err
//	}
//
//	if media == nil {
//		ctx.JSON(http.StatusNotFound, gin.H{
//			"error": "Media Not Found",
//		})
//		return nil
//	}
//
//	ctx.JSON(http.StatusOK, media)
//	return nil
//}

// @Summary Delete media
// @Description Delete media by its ID
// @Tags Media
// @Produce json
// @Param mediaID path string true "Media ID"
// @Success 204 "Successfully deleted media"
// @Failure 404 {object} Dto.ErrorDTO "Media Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/media/{mediaID} [delete]
//func (c *MediaController) deleteMedia(ctx *gin.Context) error {
//	mediaID := ctx.Param("mediaID")
//
//	err := c.mediaService.DeleteMedia(mediaID)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"error": "Internal Server Error",
//		})
//		return err
//	}
//
//	ctx.JSON(http.StatusNoContent, nil)
//	return nil
//}

// @Summary List all media for a post
// @Description Retrieve all media for the specified post
// @Tags Media
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {array} models.Media "Successfully retrieved media"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/media/post/{postID} [get]
//func (c *MediaController) listMedia(ctx *fiber.Ctx) error {
//	postID := ctx.Params("postID")
//	mediaList, err := c.mediaService.ListMedia(postID)
//	if err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "Internal Server Error",
//		})
//	}
//
//	return ctx.JSON(mediaList)
//}

//# Pull the pre-built Redoc Docker image
//FROM redocly/redoc
//
//# Copy your local swagger.json file to a directory inside the Docker image
//COPY ./app/swagger.json /usr/share/nginx/html/swagger.json
//
//# Set the environment variable for your OpenAPI definition URL
//ENV PAGE_TITLE="Placio API Documentation"
//ENV PAGE_FAVICON="https://res.cloudinary.com/placio/image/upload/v1686763351/fpv01oen8dx3g7uyiezl.png"
//ENV SPEC_URL="/swagger.json"
//ENV PORT=80
//
//# Expose port 80 for the app
//EXPOSE 80

//# --- Build stage ---
//FROM node:16-alpine AS build-stage
//
//WORKDIR /app
//
//# Install the Redocly CLI
//RUN npm install -g @redocly/cli@latest
//
//# Copy documentation files
//COPY ./docs ./docs
//
//ENV NODE_OPTIONS="--max-old-space-size=16384"
//# Bundle the documentation into a static OpenAPI file
//RUN redocly bundle ./docs/app/swagger.yaml -o ./docs/swagger.yaml
//
//# Generate the documentation HTML file
//RUN redocly build-docs ./docs/swagger.yaml --output ./docs/redoc-static.html
//
//# --- Production stage ---
//FROM nginx:alpine AS production-stage
//
//# Remove default nginx page
//RUN rm -rf /usr/share/nginx/html/*
//
//# Copy the static HTML file to the nginx html directory
//COPY --from=build-stage /app/docs/redoc-static.html /usr/share/nginx/html/index.html
//
//EXPOSE 80
//
//CMD ["nginx", "-g", "daemon off;"]
