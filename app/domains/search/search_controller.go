package search

import (
	"net/http"
	_ "placio-app/Dto"
	"placio-app/ent"
	_ "placio-app/ent"
	"placio-app/utility"
	"sync"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
	searchService SearchService
}

func NewSearchController(service SearchService) SearchController {
	return SearchController{searchService: service}
}

func (ss *SearchController) RegisterRoutes(route *gin.RouterGroup) {
	searchRoute := route.Group("/search")
	{
		searchRoute.GET("/", utility.Use(ss.search))
		searchRoute.GET("/db", utility.Use(ss.searchDB))
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

		return err
	}
	businesses, err := ss.searchService.SearchBusinesses(ctx, searchText)
	if err != nil {

		return err
	}
	events, err := ss.searchService.SearchEvents(ctx, searchText)
	if err != nil {

		return err
	}
	places, err := ss.searchService.SearchPlaces(ctx, searchText)
	if err != nil {

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

// @Summary Search DB
// @Description Search for Users, Places, Events, and Businesses
// @Tags Search
// @Accept json
// @Produce json
// @Param type query string false "Type of search - user, place, event, business"
// @Param searchText query string true "Text to search for"
// @Success 200 {object} Dto.SearchResponses "Successfully found search results"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/search/db [get]
func (ss *SearchController) searchDB(ctx *gin.Context) error {
	searchType := ctx.Query("type")
	searchText := ctx.Query("searchText")

	// Define a struct to encapsulate search results and possible error
	type searchResult struct {
		Users      []*ent.User
		Businesses []*ent.Business
		Events     []*ent.Event
		Places     []*ent.Place
		Err        error
	}

	// Make searches concurrently
	resultsChan := make(chan *searchResult, 4) // Buffer to avoid blocking

	searchAll := searchType == ""
	var wg sync.WaitGroup // Use wait group to wait for all goroutines to finish

	// Define a closure for concurrent searches
	performSearch := func(searchType string) {
		defer wg.Done() // Ensure the wait group is signaled upon completion
		result := &searchResult{}

		switch searchType {
		case "user":
			result.Users, result.Err = ss.searchService.SearchUsersDB(ctx, searchText)
		case "business":
			result.Businesses, result.Err = ss.searchService.SearchBusinessesDB(ctx, searchText)
		case "event":
			result.Events, result.Err = ss.searchService.SearchEventsDB(ctx, searchText)
		case "place":
			result.Places, result.Err = ss.searchService.SearchPlacesDB(ctx, searchText)
		}

		resultsChan <- result // Send the result to the channel
	}

	// Start concurrent searches depending on the searchType
	if searchAll || searchType == "user" {
		wg.Add(1)
		go performSearch("user")
	}
	if searchAll || searchType == "business" {
		wg.Add(1)
		go performSearch("business")
	}
	if searchAll || searchType == "event" {
		wg.Add(1)
		go performSearch("event")
	}
	if searchAll || searchType == "place" {
		wg.Add(1)
		go performSearch("place")
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Collect the results
	var users []*ent.User
	var businesses []*ent.Business
	var events []*ent.Event
	var places []*ent.Place

	for result := range resultsChan {
		if result.Err != nil {
			return result.Err
		}
		users = append(users, result.Users...)
		businesses = append(businesses, result.Businesses...)
		events = append(events, result.Events...)
		places = append(places, result.Places...)
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

		return err
	}
	ctx.JSON(http.StatusOK, businesses)
	return nil
}
