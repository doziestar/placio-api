package smartGym

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/ent"
	"placio-pkg/middleware"
)

type SmartFitnessController struct {
	smartFitnessService ISmartFitness
}

// NewSmartFitnessController creates a new controller for handling fitness-related requests.
func NewSmartFitnessController(smartFitnessService ISmartFitness) *SmartFitnessController {
	return &SmartFitnessController{smartFitnessService: smartFitnessService}
}

// RegisterRoutes sets up the routes for gym, trainer, member, and subscription endpoints.
func (c *SmartFitnessController) RegisterRoutes(router, routerWithAuth *gin.RouterGroup) {
	trainerRouter := router.Group("/fitness/trainers")
	trainerRouterWithAuth := routerWithAuth.Group("/fitness/trainers")
	{
		trainerRouterWithAuth.POST("/", middleware.ErrorMiddleware(c.createTrainer))
		trainerRouter.GET("/", middleware.ErrorMiddleware(c.getTrainers))
		trainerRouter.GET("/:trainerId", middleware.ErrorMiddleware(c.getTrainerByID))
		trainerRouterWithAuth.PUT("/:trainerId", middleware.ErrorMiddleware(c.updateTrainer))
		trainerRouterWithAuth.DELETE("/:trainerId", middleware.ErrorMiddleware(c.deleteTrainer))
	}

	memberRouter := router.Group("/fitness/members")
	memberRouterWithAuth := routerWithAuth.Group("/fitness/members")
	{
		memberRouterWithAuth.POST("/", middleware.ErrorMiddleware(c.createMember))
		memberRouter.GET("/", middleware.ErrorMiddleware(c.getMembers))
		memberRouter.GET("/:memberId", middleware.ErrorMiddleware(c.getMemberByID))
		memberRouterWithAuth.PUT("/:memberId", middleware.ErrorMiddleware(c.updateMember))
		memberRouterWithAuth.DELETE("/:memberId", middleware.ErrorMiddleware(c.deleteMember))
	}

	subscriptionRouter := router.Group("/fitness/subscriptions")
	subscriptionRouterWithAuth := routerWithAuth.Group("/fitness/subscriptions")
	{
		subscriptionRouterWithAuth.POST("/", middleware.ErrorMiddleware(c.createSubscription))
		subscriptionRouter.GET("/", middleware.ErrorMiddleware(c.getSubscriptions))
		subscriptionRouter.GET("/:subscriptionId", middleware.ErrorMiddleware(c.getSubscriptionByID))
		subscriptionRouterWithAuth.PUT("/:subscriptionId", middleware.ErrorMiddleware(c.updateSubscription))
		subscriptionRouterWithAuth.DELETE("/:subscriptionId", middleware.ErrorMiddleware(c.deleteSubscription))
	}
}

// createTrainer creates a new trainer.
// @Summary Create a new trainer
// @Description Create a new trainer
// @Tags trainers
// @Accept json
// @Produce json
// @Param trainer body Trainer true "Trainer object"
// @Success 201 {object} Trainer
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /trainers [post]
func (c *SmartFitnessController) createTrainer(ctx *gin.Context) error {
	var trainerDto ent.Trainer
	if err := ctx.ShouldBindJSON(&trainerDto); err != nil {
		return err
	}

	trainer, err := c.smartFitnessService.CreateTrainer(ctx, &trainerDto)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusCreated, trainer)
	return nil
}

func (c *SmartFitnessController) getTrainers(ctx *gin.Context) error {
	trainers, err := c.smartFitnessService.GetTrainers(ctx)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, trainers)
	return nil
}

func (c *SmartFitnessController) getTrainerByID(ctx *gin.Context) error {
	trainerId := ctx.Param("trainerId")
	trainer, err := c.smartFitnessService.GetTrainerByID(ctx, trainerId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, trainer)
	return nil
}

func (c *SmartFitnessController) updateTrainer(ctx *gin.Context) error {
	var trainerDto ent.Trainer
	if err := ctx.ShouldBindJSON(&trainerDto); err != nil {
		return err
	}

	trainerId := ctx.Param("trainerId")
	trainer, err := c.smartFitnessService.UpdateTrainer(ctx, trainerId, &trainerDto)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, trainer)
	return nil
}

func (c *SmartFitnessController) deleteTrainer(ctx *gin.Context) error {
	trainerId := ctx.Param("trainerId")
	err := c.smartFitnessService.DeleteTrainer(ctx, trainerId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusNoContent, nil)
	return nil
}

// createSubscription creates a new subscription.
func (c *SmartFitnessController) createSubscription(ctx *gin.Context) error {
	var subscriptionDto ent.Subscription
	if err := ctx.ShouldBindJSON(&subscriptionDto); err != nil {
		return err
	}

	subscription, err := c.smartFitnessService.CreateSubscription(ctx, &subscriptionDto)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusCreated, subscription)
	return nil
}

// getSubscriptions retrieves all subscriptions.
func (c *SmartFitnessController) getSubscriptions(ctx *gin.Context) error {
	subscriptions, err := c.smartFitnessService.GetSubscriptions(ctx)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, subscriptions)
	return nil
}

// getSubscriptionByID retrieves a subscription by its ID.
func (c *SmartFitnessController) getSubscriptionByID(ctx *gin.Context) error {
	subscriptionId := ctx.Param("subscriptionId")
	subscription, err := c.smartFitnessService.GetSubscriptionByID(ctx, subscriptionId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, subscription)
	return nil
}

// updateSubscription updates a subscription.
func (c *SmartFitnessController) updateSubscription(ctx *gin.Context) error {
	var subscriptionDto ent.Subscription
	if err := ctx.ShouldBindJSON(&subscriptionDto); err != nil {
		return err
	}

	subscriptionId := ctx.Param("subscriptionId")
	subscription, err := c.smartFitnessService.UpdateSubscription(ctx, subscriptionId, &subscriptionDto)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, subscription)
	return nil
}

// deleteSubscription deletes a subscription.
func (c *SmartFitnessController) deleteSubscription(ctx *gin.Context) error {
	subscriptionId := ctx.Param("subscriptionId")
	err := c.smartFitnessService.DeleteSubscription(ctx, subscriptionId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusNoContent, nil)
	return nil
}

// createMember creates a new gym member.
func (c *SmartFitnessController) createMember(ctx *gin.Context) error {
	var memberDto ent.User
	if err := ctx.ShouldBindJSON(&memberDto); err != nil {
		return err
	}

	member, err := c.smartFitnessService.CreateMember(ctx, &memberDto)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusCreated, member)
	return nil
}

// getMembers retrieves all members.
func (c *SmartFitnessController) getMembers(ctx *gin.Context) error {
	members, err := c.smartFitnessService.GetMembers(ctx)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, members)
	return nil
}

// getMemberByID retrieves a specific member by their ID.
func (c *SmartFitnessController) getMemberByID(ctx *gin.Context) error {
	memberId := ctx.Param("memberId")
	member, err := c.smartFitnessService.GetMemberByID(ctx, memberId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, member)
	return nil
}

// updateMember updates a member.
func (c *SmartFitnessController) updateMember(ctx *gin.Context) error {
	var memberDto ent.User
	if err := ctx.ShouldBindJSON(&memberDto); err != nil {
		return err
	}

	memberId := ctx.Param("memberId")
	member, err := c.smartFitnessService.UpdateMember(ctx, memberId, &memberDto)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, member)
	return nil
}

// deleteMember deletes a member.
func (c *SmartFitnessController) deleteMember(ctx *gin.Context) error {
	memberId := ctx.Param("memberId")
	err := c.smartFitnessService.DeleteMember(ctx, memberId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusNoContent, nil)
	return nil
}
