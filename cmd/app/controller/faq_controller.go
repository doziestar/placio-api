package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"
)

type FAQController struct {
	service service.FAQService
}

func NewFAQController(service service.FAQService) *FAQController {
	return &FAQController{service: service}
}

func (fc *FAQController) RegisterRoutes(router, routerWithoutAuth *gin.RouterGroup) {
	faqRouter := router.Group("/faq")
	//faqRouterWithoutAuth := routerWithoutAuth.Group("/faq")
	{
		faqRouter.POST("business/:businessID", utility.Use(fc.createFAQ))
		faqRouter.GET("/:faqID", utility.Use(fc.getFAQ))
		faqRouter.PATCH("/:faqID", utility.Use(fc.updateFAQ))
		faqRouter.DELETE("/:faqID", utility.Use(fc.deleteFAQ))
		faqRouter.POST("/:faqID/place/:placeID", utility.Use(fc.associateFAQWithPlace))
		faqRouter.POST("/:faqID/event/:eventID", utility.Use(fc.associateFAQWithEvent))
		faqRouter.GET("/business/:businessID", utility.Use(fc.getFAQsByBusiness))
	}
}

// @Summary Create a new FAQ
// @ID create-faq
// @Tags FAQ
// @Produce json
// @Param businessID path string true "Business ID"
// @Param faqData body Dto.FAQDto true "FAQ Data"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Create a new FAQ for a specific Business
// @Success 201 {object} Dto.FAQDto
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /faq/business/{businessID} [post]
func (fc *FAQController) createFAQ(c *gin.Context) error {
	businessID := c.Param("businessID")

	var faqData Dto.FAQDto
	if err := c.ShouldBindJSON(&faqData); err != nil {

		return nil
	}

	faq, err := fc.service.CreateFAQ(c, businessID, &faqData)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusCreated, faq)
	return nil
}

// @Summary Get FAQ by ID
// @Description Retrieve FAQ by ID
// @ID get-FAQ
// @Tags FAQ
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param faqID path string true "FAQ ID"
// @Success 200 {object} ent.FAQ
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /faq/{faqID} [get]
func (fc *FAQController) getFAQ(c *gin.Context) error {
	faqID := c.Param("faqID")

	faq, err := fc.service.GetFAQ(c, faqID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, faq)
	return nil
}

// @Summary Get FAQs by Business ID
// @Description Retrieve FAQs by Business ID
// @ID get-FAQs-by-business
// @Tags FAQ
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param businessID path string true "Business ID"
// @Success 200 {object} []ent.FAQ
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /faq/business/{businessID} [get]
func (fc *FAQController) getFAQsByBusiness(c *gin.Context) error {
	businessID := c.Param("businessID")

	faqs, err := fc.service.GetFAQsByBusiness(c, businessID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, faqs)
	return nil
}

// @Summary Update FAQ by ID
// @Description Update FAQ by ID
// @ID update-FAQ
// @Tags FAQ
// @Accept  json
// @Produce  json
// @Param faqID path string true "FAQ ID"
// @Param Authorization header string true "Bearer token"
// @Param faqData body Dto.FAQDto true "FAQ Data"
// @Success 200 {object} ent.FAQ
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Router /faq/{faqID} [patch]
func (fc *FAQController) updateFAQ(c *gin.Context) error {
	faqID := c.Param("faqID")

	var faqData map[string]interface{}
	if err := c.ShouldBindJSON(&faqData); err != nil {

		return nil
	}

	faq, err := fc.service.UpdateFAQ(c, faqID, faqData)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, faq)
	return nil
}

// @Summary Delete FAQ by ID
// @Description Delete FAQ by ID
// @ID delete-FAQ
// @Tags FAQ
// @Accept  json
// @Produce  json
// @Param faqID path string true "FAQ ID"
// @Success 200 {object} string "FAQ deleted successfully"
// @Param Authorization header string true "Bearer token"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /faq/{faqID} [delete]
func (fc *FAQController) deleteFAQ(c *gin.Context) error {
	faqID := c.Param("faqID")

	err := fc.service.DeleteFAQ(c, faqID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "FAQ deleted successfully", ""))
	return nil
}

// @Summary Associate FAQ with Place
// @Description Associate a FAQ with a Place by their IDs
// @ID associate-FAQ-place
// @Tags FAQ
// @Accept  json
// @Produce  json
// @Param faqID path string true "FAQ ID"
// @Param placeID path string true "Place ID"
// @Success 200 {object} string "FAQ associated with place successfully"
// @Param Authorization header string true "Bearer token"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /faq/{faqID}/place/{placeID} [post]
func (fc *FAQController) associateFAQWithPlace(c *gin.Context) error {
	faqID := c.Param("faqID")
	placeID := c.Param("placeID")

	err := fc.service.AssociateFAQWithPlace(c, faqID, placeID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "FAQ associated with place successfully", ""))
	return nil
}

// @Summary Associate FAQ with Event
// @Description Associate a FAQ with an Event by their IDs
// @ID associate-FAQ-event
// @Tags FAQ
// @Accept  json
// @Produce  json
// @Param faqID path string true "FAQ ID"
// @Param eventID path string true "Event ID"
// @Success 200 {object} string "FAQ associated with event successfully"
// @Param Authorization header string true "Bearer token"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /faq/{faqID}/event/{eventID} [post]
func (fc *FAQController) associateFAQWithEvent(c *gin.Context) error {
	faqID := c.Param("faqID")
	eventID := c.Param("eventID")

	err := fc.service.AssociateFAQWithEvent(c, faqID, eventID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "FAQ associated with event successfully", ""))
	return nil
}
