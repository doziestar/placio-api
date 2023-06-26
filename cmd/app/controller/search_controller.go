package controller

import (
	"net/http"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
	searchService service.SearchService
}

func NewSearchController(service service.SearchService) SearchController {
	return SearchController{searchService: service}
}

func (ss *SearchController) RegisterRoutes(route *gin.RouterGroup) {
	searchRoute := route.Group("/search")
	{
		searchRoute.GET("/", utility.Use(ss.search))
		searchRoute.GET("/users", utility.Use(ss.searchUsers))
		searchRoute.GET("/places", utility.Use(ss.searchPlaces))
		searchRoute.GET("/events", utility.Use(ss.searchEvents))
		searchRoute.GET("/businesses", utility.Use(ss.searchBusinesses))
	}
}

// @Summary Full-text search
// @Description Search for Users, Places, Events, and Businesses
// @Tags Search
// @Accept json
// @Produce json
// @Param searchText query string true "Text to search for"
// @Success 200 {object} Dto.SearchResponse "Successfully found search results"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/search [get]
func (ss *SearchController) search(ctx *gin.Context) error {
	searchText := ctx.Query("searchText")
	users, err := ss.searchService.SearchUsers(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	businesses, err := ss.searchService.SearchBusinesses(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	events, err := ss.searchService.SearchEvents(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	places, err := ss.searchService.SearchPlaces(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users":      users,
		"businesses": businesses,
		"events":     events,
		"places":     places,
	})
	return nil
}

// @Summary Search for Users
// @Description Search for Users by a given search text
// @Tags Search
// @Accept json
// @Produce json
// @Param searchText query string true "Text to search for"
// @Success 200 {object} []ent.User "Successfully found search results"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/search/users [get]
func (ss *SearchController) searchUsers(ctx *gin.Context) error {
	searchText := ctx.Query("searchText")
	users, err := ss.searchService.SearchUsers(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, users)
	return nil
}

// @Summary Search for Places
// @Description Search for Places by a given search text
// @Tags Search
// @Accept json
// @Produce json
// @Param searchText query string true "Text to search for"
// @Success 200 {object} []ent.Place "Successfully found search results"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/search/places [get]
func (ss *SearchController) searchPlaces(ctx *gin.Context) error {
	searchText := ctx.Query("searchText")
	places, err := ss.searchService.SearchPlaces(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, places)
	return nil
}

// @Summary Search for Events
// @Description Search for Events by a given search text
// @Tags Search
// @Accept json
// @Produce json
// @Param searchText query string true "Text to search for"
// @Success 200 {object} []ent.Event "Successfully found search results"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/search/events [get]
func (ss *SearchController) searchEvents(ctx *gin.Context) error {
	searchText := ctx.Query("searchText")
	events, err := ss.searchService.SearchEvents(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, events)
	return nil
}

// @Summary Search for Businesses
// @Description Search for Businesses by a given search text
// @Tags Search
// @Accept json
// @Produce json
// @Param searchText query string true "Text to search for"
// @Success 200 {object} []ent.Business "Successfully found search results"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/search/businesses [get]
func (ss *SearchController) searchBusinesses(ctx *gin.Context) error {
	searchText := ctx.Query("searchText")
	businesses, err := ss.searchService.SearchBusinesses(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, businesses)
	return nil
}
