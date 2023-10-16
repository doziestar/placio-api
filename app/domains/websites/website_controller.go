package websites

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/ent"
	"placio-app/utility"
	"placio-pkg/middleware"
)

type WebsiteController struct {
	websiteService IWebsite
}

func NewWebsiteController(websiteService IWebsite) *WebsiteController {
	return &WebsiteController{
		websiteService: websiteService,
	}
}

func (w *WebsiteController) RegisterRoutes(router *gin.RouterGroup) {
	websiteRouter := router.Group("/website")
	{
		websiteRouter.GET("/:businessID", middleware.ErrorMiddleware(w.getBusinessWebsite))
		websiteRouter.POST("/:businessID", middleware.ErrorMiddleware(w.createBusinessWebsite))
		websiteRouter.PATCH("/:businessID", middleware.ErrorMiddleware(w.updateBusinessWebsite))
		websiteRouter.GET("/verify/:domainName", middleware.ErrorMiddleware(w.verifyDomainName))
	}
}

func (w *WebsiteController) getBusinessWebsite(c *gin.Context) error {
	// get the business id
	businessID := c.Param("businessID")
	// get the website
	website, err := w.websiteService.GetBusinessWebsite(c, businessID)
	if err != nil {
		return err
	}
	// return the website
	c.JSON(http.StatusOK, utility.ProcessResponse(website))
	return nil
}

func (w *WebsiteController) verifyDomainName(c *gin.Context) error {

	// get the domain name
	domainName := c.Param("domainName")
	// get the website

	domainNameExists, err := w.websiteService.VerifyDomainName(c, domainName)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(domainNameExists))
	return nil
}

func (w *WebsiteController) createBusinessWebsite(c *gin.Context) error {
	// get the business id
	businessID := c.Param("businessID")
	var websiteData *ent.Website

	if err := c.ShouldBind(&websiteData); err != nil {
		return err
	}
	// get the website
	website, err := w.websiteService.CreateBusinessWebsite(c, businessID, websiteData)
	if err != nil {

		return err
	}
	// return the website
	c.JSON(http.StatusOK, utility.ProcessResponse(website))
	return nil
}

func (w *WebsiteController) updateBusinessWebsite(c *gin.Context) error {
	// get the business id
	businessID := c.Param("businessID")
	var websiteData *ent.Website

	if err := c.ShouldBind(&websiteData); err != nil {
		return err
	}
	// get the website
	website, err := w.websiteService.UpdateBusinessWebsite(c, businessID, websiteData)
	if err != nil {
		return err
	}
	// return the website
	c.JSON(http.StatusOK, utility.ProcessResponse(website))
	return nil
}
