package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/service"
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
		searchRoute.GET("/users", ss.searchUsers)
		searchRoute.GET("/places", ss.searchPlaces)
		searchRoute.GET("/events", ss.searchEvents)
		searchRoute.GET("/businesses", ss.searchBusinesses)
	}
}

func (ss *SearchController) searchUsers(ctx *gin.Context) {
	searchText := ctx.Query("searchText")
	users, err := ss.searchService.SearchUsers(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (ss *SearchController) searchPlaces(ctx *gin.Context) {
	searchText := ctx.Query("searchText")
	places, err := ss.searchService.SearchPlaces(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, places)
}

func (ss *SearchController) searchEvents(ctx *gin.Context) {
	searchText := ctx.Query("searchText")
	events, err := ss.searchService.SearchEvents(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func (ss *SearchController) searchBusinesses(ctx *gin.Context) {
	searchText := ctx.Query("searchText")
	businesses, err := ss.searchService.SearchBusinesses(ctx, searchText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, businesses)
}
