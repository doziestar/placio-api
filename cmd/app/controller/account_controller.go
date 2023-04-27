package controller

//https://docs.gofiber.io/guide/validation

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"placio-app/Dto"
	"placio-app/errors"
	"placio-app/middleware"
	_ "placio-app/models"
	"placio-app/service"
	"placio-app/utility"
	"placio-pkg/logger"
)

type AccountController struct {
	store   service.IAccountService
	user    service.IUser
	utility utility.IUtility
}

func NewAccountController(store service.IAccountService, utility utility.IUtility) *AccountController {
	return &AccountController{store: store, utility: utility}
}

//
//func requestLogger() gin.HandlerFunc {
//	return func(c *gin.Context) error {
//		fmt.Printf("Request: %s %s\n", c.Method(), c.Path())
//		return c.Next()
//	}
//}

func (c *AccountController) RegisterRoutes(app *gin.RouterGroup) {
	//app.Use(requestLogger())
	accountGroup := app.Group("/accounts")
	accountGroup.GET("/", middleware.Verify("user"), utility.Use(c.getUserAccount))
	accountGroup.POST("/create-account", utility.Use(c.createAccount))
	accountGroup.POST("/:accountId/switch-account/", middleware.Verify("user"), utility.Use(c.switchAccount))
	accountGroup.POST("/:accountId/make-default/", middleware.Verify("user"), utility.Use(c.makeAccountDefault))
	accountGroup.POST("/add-account", middleware.Verify("user"), utility.Use(c.addAccount)) // add account to owner
	accountGroup.POST("/plan", middleware.Verify("owner"), utility.Use(c.plan))
	accountGroup.PATCH("/plan", middleware.Verify("owner"), utility.Use(c.updatePlan))
	accountGroup.GET("/get-user-accounts", middleware.Verify("user"), utility.Use(c.getAccounts))
	accountGroup.GET("/get-user-active-account", middleware.Verify("user"), utility.Use(c.getUserActiveAccount))
	accountGroup.GET("/:accountId", middleware.Verify("user"), utility.Use(c.getUserAccount))
	accountGroup.PATCH("/card", middleware.Verify("owner"), utility.Use(c.updateInvoice))
	accountGroup.GET("/invoice", middleware.Verify("owner"), utility.Use(c.getInvoice))
	accountGroup.GET("/plans", middleware.Verify("owner"), utility.Use(c.getPlans))
	accountGroup.GET("/subscription", middleware.Verify("owner"), utility.Use(c.getSubscription))
	accountGroup.POST("/upgrade", middleware.Verify("owner"), utility.Use(c.upgradePlan))
	accountGroup.DELETE("/", middleware.Verify("owner"), utility.Use(c.deleteAccount))
	accountGroup.POST("/:id/follow", middleware.Verify("user"), utility.Use(c.followAccount))
	accountGroup.POST("/:id/unfollow", middleware.Verify("user"), utility.Use(c.unfollowAccount))
	accountGroup.GET("/:id/followers", middleware.Verify("user"), utility.Use(c.getFollowers))
	accountGroup.GET("/:id/following", middleware.Verify("user"), utility.Use(c.getFollowing))

}

// CreateAccount creates a new user account and assigns the user to the account.
// The function performs the following steps:
// 1. Parses the incoming request body into a SignUpDto.
// 2. Validates the input data.
// 3. Checks if the user has already registered an account.
// 4. Creates a new account and assigns the user to it.
// 5. Sends a welcome email to the user.
// 6. Authenticates the user (currently commented out).
//
// @Summary Create a new account
// @Description Create a new account and assign the user to it
// @Tags Account
// @Accept json
// @Produce json
// @Param SignUpDto body Dto.SignUpDto true "Sign Up Data"
// @Success 200 {object} Dto.UserResponse "Successfully created account"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 403 {object} Dto.ErrorDTO "Forbidden"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/accounts/create-account [post]
func (c *AccountController) createAccount(ctx *gin.Context) error {
	fmt.Println("Entering createAccount function")
	data := new(Dto.SignUpDto)

	logger.Info(context.Background(), fmt.Sprintf("data: %v", data))
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	log.Println("CreateAccount", data)

	// validate input
	if err := utility.Validate(data.Email, data.Name, data.Password, data.Username); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	// confirm_password field is a dummy field to prevent bot signups
	if data.ConfirmPassword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Confirm password is required"})
		return errors.ErrInvalid
	}

	response, err := c.store.CreateUserAccount(data, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, response)
	return nil
}

// SwitchAccount switches the user to a different account.
// The function performs the following steps:
// 1. Parses the incoming request body into a SwitchAccountDto.
// 2. Validates the input data.
// 3. Checks if the user is a member of the account.
// 4. Switches the user to the account.
//
// @Summary Switch to a different account
// @Description Switch to a different account
// @Tags Account
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} Dto.UserResponse "Successfully switched account"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 403 {object} Dto.ErrorDTO "Forbidden"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/accounts/{accountId}/switch-account [post]
func (c *AccountController) switchAccount(ctx *gin.Context) error {
	accountId := ctx.Param("accountId")
	userId, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return nil
	}
	response, err := c.store.SwitchUserAccount(accountId, userId.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	data, err := utility.RemoveSensitiveFields(response)

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully switched account", "data": data, "status": "success"})
	return nil
}

// Plan godoc
// @Summary subscribe to a plan
// @Description subscribe to a plan
// @Tags Account
// @Accept json
// @Produce json
// @Param plan body string true "plan"
// @Param token body string true "token"
// @Param stripe.customer.id body string true "stripe.customer.id"
// @Param stripe.subscription.id body string true "stripe.subscription.id"
// @Param stripe.subscription.price body string true "stripe.subscription.price"
// @Success 200 {object} models.Account "account"
// @Failure 400 {object} Dto.ErrorDTO "inputError": "plan", "message": "Plan is required"
// @Failure 500 {object} Dto.ErrorDTO "error": "Internal Server Error"
// @Router /accounts/plan [post]
func (c *AccountController) plan(ctx *gin.Context) error {
	//	data := new(struct {
	//		Plan   string              `json:"plan"`
	//		Token  *stripe.TokenParams `json:"token,omitempty"`
	//		Stripe struct {
	//			Customer struct {
	//				ID string `json:"id,omitempty"`
	//			} `json:"customer,omitempty"`
	//			Subscription struct {
	//				ID    string `json:"id,omitempty"`
	//				Price int64  `json:"price,omitempty"`
	//			} `json:"subscription,omitempty"`
	//		} `json:"stripe,omitempty"`
	//	})
	//
	//	if err := c.BodyParser(data); err != nil {
	//		return c.Status(fiber.StatusBadRequest).JSON(gin.H{
	//			"error": "Bad request",
	//		})
	//	}
	//
	//	if data.Plan == "" {
	//		return c.Status(fiber.StatusBadRequest).JSON(gin.H{
	//			"error": "Plan is required",
	//		})
	//	}
	//
	//	// check the plan exists
	//	plan, ok := settings.Plans[data.Plan]
	//	if !ok {
	//		return c.Status(fiber.StatusBadRequest).JSON(gin.H{
	//			"error": "Plan doesn't exist",
	//		})
	//	}
	//
	//	accountData, err := account.GET(c)
	//	if err != nil {
	//		return c.Status(fiber.StatusNotFound).JSON(gin.H{
	//			"error": "No account with that ID",
	//		})
	//	}
	//
	//	// process stripe subscription for non-free accounts
	//	// if a 2-factor payment hasn't occurred, create the stripe subscription
	//	if data.Plan != "free" {
	//		if data.Stripe == (struct {
	//			Customer struct {
	//				ID string "json:\"id,omitempty\""
	//			} "json:\"customer,omitempty\""
	//			Subscription struct {
	//				ID    string "json:\"id,omitempty\""
	//				Price int64  "json:\"price,omitempty\""
	//			} "json:\"subscription,omitempty\""
	//		}{}) {
	//			if data.Token == nil || data.Token.ID == "" {
	//				return c.Status(fiber.StatusBadRequest).JSON(gin.H{
	//					"error": "Please enter your credit card details",
	//				})
	//			}
	//
	//			// create a stripe customer and subscribe them to a plan
	//			customer, err := stripe.CustomerCreate(
	//				accountData.OwnerEmail,
	//				data.Token.ID,
	//			)
	//			if err != nil {
	//				return c.Status(fiber.StatusInternalServerError).JSON(gin.H{
	//					"error": err.Error(),
	//				})
	//			}
	//			data.Stripe.Customer.ID = customer.ID
	//
	//			subscription, err := stripe.CustomerSubscribe(
	//				customer.ID,
	//				data.Plan,
	//			)
	//			if err != nil {
	//				return c.Status(fiber.StatusInternalServerError).JSON(gin.H{
	//					"error": err.Error(),
	//				})
	//			}
	//			data.Stripe.Subscription.ID = subscription.ID
	//
	//			// check for an incomplete payment that requires 2-factor authentication
	//			if subscription.LatestInvoice.PaymentIntent.Status == "requires_action" {
	//				log.Println("Stripe payment requires further action")
	//
	//				return c.Status(fiber.StatusOK).JSON(gin.H{
	//					"requires_payment_action": true,
	//					"customer": gin.H{
	//						"id": data.Stripe.Customer.ID,
	//					},
	//					"subscription": gin.H{
	//						"id":    data.Stripe.Subscription.ID,
	//						"price": data.Stripe.Subscription.Price,
	//					},
	//					"client_secret": subscription.LatestInvoice.PaymentIntent.ClientSecret,
	//				})
	//			}
	//		}
	//
	//		// stripe info hasn't been passed back as part of 2
	//		if data.Stripe.Customer.ID == "" || data.Stripe.Subscription.ID == "" {
	//			return c.Status(fiber.StatusBadRequest).JSON(gin.H{
	//				"error": "Stripe customer or subscription ID is missing",
	//			})
	//		}
	//
	//		// update the stripe subscription
	//		subscription, err := stripe.CustomerUpdateSubscription(
	//			data.Stripe.Customer.ID,
	//			data.Stripe.Subscription.ID,
	//			data.Plan,
	//		)
	//		if err != nil {
	//			return c.Status(fiber.StatusInternalServerError).JSON(gin.H{
	//				"error": err.Error(),
	//			})
	//		}
	//		data.Stripe.Subscription.ID = subscription.ID
	//	}
	//
	//	// update the account plan
	//	if err := account.UpdatePlan(accountData.ID, data.Plan, data.Stripe.Customer.ID, data.Stripe.Subscription.ID); err != nil {
	//		return c.Status(fiber.StatusInternalServerError).JSON(gin.H{
	//			"error": "Internal Server Error",
	//		})
	//	}
	//
	//	c.Status(fiber.StatusOK).JSON(gin.H{
	//		"success": true,
	//	})
	//
	return nil
}

func SendMail(to, template string, content map[string]interface{}) error {
	// implementation of mail.send()
	return nil
}

// UpdatePlan godoc
// @Summary Update account plan
// @Description Update account plan
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param plan body string true "Plan ID"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /accounts/{id}/plan [put]
func (c *AccountController) updatePlan(ctx *gin.Context) error {
	//	data := new(struct {
	//		ID   string `json:"id"`
	//		Plan string `json:"plan"`
	//	})
	//	if err := c.BodyParser(data); err != nil {
	//		return err
	//	}
	//
	//	// implementation of utility.validate()
	//	// return an error if the request body is invalid
	//
	//	accountID := ""
	//	if c.Locals("permission") == "master" {
	//		accountID = data.ID
	//	} else {
	//		accountID = c.Locals("account").(string)
	//	}
	//
	//	plan := Plan{}
	//	for _, p := range plans {
	//		if p.ID == data.Plan {
	//			plan = p
	//			break
	//		}
	//	}
	//	if plan.ID == "" {
	//		return fmt.Errorf("No plan with that ID")
	//	}
	//
	//	accountData, err := GetAccountData(accountID)
	//	if err != nil {
	//		return err
	//	}
	//
	//	// user is upgrading from paid to free,
	//	// direct them to the upgrade view
	//	if accountData.Plan == "free" && plan.ID != "free" {
	//		if c.Locals("permission") == "master" {
	//			return fiber.NewError(fiber.StatusForbidden, "The account holder will need to enter their card details and upgrade to a paid plan.")
	//		} else {
	//			return c.Status(fiber.StatusPaymentRequired).JSON(gin.H{
	//				"message": "Please upgrade your account",
	//				"plan":    plan.ID,
	//			})
	//		}
	//	}
	//
	//	if plan.ID == "free" {
	//		// user is downgrading - cancel the stripe subscription
	//		if accountData.StripeSubscriptionID != "" {
	//			subscription, err := sub.GET(accountData.StripeSubscriptionID, nil)
	//			if err != nil {
	//				return err
	//			}
	//
	//			err = updateAccountData(accountID, &AccountData{
	//				StripeSubscriptionID: "",
	//				Plan:                 plan.ID,
	//			})
	//			if err != nil {
	//				return err
	//			}
	//
	//			if subscription.Status != "canceled" {
	//				_, err := sub.Cancel(accountData.StripeSubscriptionID, nil)
	//				if err != nil {
	//					return err
	//				}
	//			}
	//		}
	//	} else {
	//		// user is switching to a different paid plan
	//		if accountData.StripeSubscriptionID != "" {
	//			// check for an incomplete payment that requires 2-factor authentication
	//			subscription, err := sub.GET(accountData.StripeSubscriptionID, nil)
	//			if err != nil {
	//				return err
	//			}
	//
	//			if subscription.LatestInvoice.PaymentIntent.Status == "requires_action" {
	//				return c.Status(fiber.StatusOK).JSON(gin.H{
	//					"requires_payment_action": true,
	//					"customer": gin.H{
	//						"id": accountData.StripeCustomerID,
	//					},
	//					"subscription": gin.H{
	//						"id":    accountData.StripeSubscriptionID,
	//						"price": plan.Price,
	//					},
	//					"client_secret": subscription.LatestInvoice.PaymentIntent.ClientSecret,
	//				})
	//			}
	//
	//			// update the stripe subscription
	//			subscription, err = sub.Update(accountData.StripeSubscriptionID, &stripe.SubscriptionParams{
	//				Items: []*stripe.SubscriptionItemsParams{
	//					{
	//						ID:   subscription.Items.Data[0].ID,
	//						Plan: stripe.String(plan.ID),
	//					},
	//				},
	//			})
	//			if err != nil {
	//				return err
	//			}
	//
	//			err = updateAccountData(accountID, &AccountData{
	//				StripeSubscriptionID: subscription.ID,
	//				Plan:                 plan.ID,
	//			})
	//			if err != nil {
	//				return err
	//			}
	//
	//			// send an email to the account holder
	//			err = sendMail(accountData.Email, "plan-upgraded", gin.H{
	//				"plan": plan.Name,
	//			})
	//			if err != nil {
	//				return err
	//			}
	//		}
	//	}
	//
	return nil
}

// GetAccounts godoc
// @Summary GET account
// @Description GET account
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} models.Account "Account"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /accounts/get-user-accounts [get]
func (c *AccountController) getAccounts(ctx *gin.Context) error {

	accounts, err := c.store.GetAccounts(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	ctx.JSON(fiber.StatusOK, accounts)
	return nil
	//	// implementation of utility.validate()
	//	// return an error if the request body is invalid
	//
	//	// get the account data
	//	accountData, err := GetAccountData(c.Locals("account").(string))
	//	if err != nil {
	//		return err
	//	}
	//
	//	// get the account
	//	account, err := GetAccount(accountData.AccountID)
	//	if err != nil {
	//		return err
	//	}
	//
	//	// get the plan
	//	plan := Plan{}
	//	for _, p := range plans {
	//		if p.ID == accountData.Plan {
	//			plan = p
	//			break
	//		}
	//	}
	//
	//	// get the subscription
	//	subscription := Subscription{}
	//	if accountData.StripeSubscriptionID != "" {
	//		s, err := sub.GET(accountData.StripeSubscriptionID, nil)
	//		if err != nil {
	//			return err
	//		}
	//
	//		subscription = Subscription{
	//			ID:        s.ID,
	//			Status:    s.Status,
	//			Plan:      s.Items.Data[0].Plan.ID,
	//			Price:     s.Items.Data[0].Plan.Amount,
	//			Quantity:  s.Items.Data[0].Quantity,
	//			StartDate: s.StartDate,
	//			EndDate:   s.EndDate,
	//		}
	//	}
	//
	//	// get the invoices
	//	invoices := []Invoice{}
	//	if accountData.StripeCustomerID != "" {
	//		i := invoice.List(&stripe.InvoiceListParams{
	//			Customer: stripe.String(accountData.StripeCustomerID),
	//			Limit:    stripe.Int64(100),
	//		})
	//		for i.Next() {
	//			invoice := i.Invoice()
	//			invoices = append(invoices, Invoice{
	//				ID:          invoice.ID,
	//				Amount:      invoice.AmountPaid,
	//				Date:        invoice.Date,
	//				PeriodStart: invoice.PeriodStart,
	//				PeriodEnd:   invoice.PeriodEnd,
	//			})
	//		}
	//		if err := i.Err(); err != nil {
	//			return err
	//		}
	//	}
	//
	//	// get the payment methods
	//	paymentMethods := []PaymentMethod{}
	//	if accountData.StripeCustomerID != "" {
	//		pm := paymentmethod.List(&stripe.PaymentMethodListParams{
	//			Customer: stripe.String(accountData.StripeCustomerID),
	//			Type:     stripe.String("card"),
	//		})
	//		for pm.Next() {
	//			paymentMethod := pm.PaymentMethod()
	//			paymentMethods = append(paymentMethods, PaymentMethod{
	//				ID:   paymentMethod.ID,
	//				Card: paymentMethod.Card.Brand + " " + paymentMethod.Card.Last4,
	//			})
	//		}
	//		if err := pm.Err(); err != nil {
	//			return err
	//		}
	//	}
	//
	//	// get the payment intents
	//	paymentIntents := []PaymentIntent{}
	//	if accountData.StripeCustomerID != "" {
	//		pi := paymentintent.List(&stripe.PaymentIntentListParams{
	//			Customer: stripe.String(accountData.StripeCustomerID),
	//			Limit:    stripe.Int64(100),
	//		})
	//		for pi.Next() {
	//			paymentIntent := pi.PaymentIntent()
	//			paymentIntents = append(paymentIntents, PaymentIntent{
	//				ID:     paymentIntent.ID,
	//				Amount: paymentIntent.Amount,
	//				Status: paymentIntent.Status,
	//				Date:   paymentIntent.Created,
	//			})
	//		}
	//		if err := pi.Err(); err != nil {
	//			return err
	//		}
	//	}
	//
	//	// get the charges
	//	charges := []Charge{}
	//	if accountData.StripeCustomerID != "" {
	//		ch := charge.List(&stripe.ChargeListParams{
	//			Customer: stripe.String(accountData.StripeCustomerID),
	//			Limit:    stripe.Int64(100),
	//		})
	//		for ch.Next() {
	//			charge := ch.Charge()
	//			charges = append(charges, Charge{
	//				ID:     charge.ID,
	//				Amount: charge.Amount,
	//				Status: charge.Status,
	//				Date:   charge.Created,
	//			})
	//		}
	//		if err := ch.Err(); err != nil {
	//			return err
	//		}
	//	}
	//
	ctx.JSON(http.StatusOK, gin.H{
		"account": account,
	})

	return nil
}

// GetAccount godoc
// @Summary GET account
// @Description GET account
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} models.Account "Account"
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO "Error"
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/get-user-active-account [get]
func (c *AccountController) getUserActiveAccount(ctx *gin.Context) error {
	log.Println("GET user active account")
	account, err := c.store.GetAccount(ctx)
	if err != nil {
		return fiber.ErrNotFound
	}

	ctx.JSON(http.StatusOK, gin.H{
		"account": account,
	})
	return nil
}

// UpdateInvoice godoc
// @Summary Update invoice
// @Description Update invoice
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Param body body Dto.Invoice true "Invoice"
// @Success 200 {object} gin.H
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/invoice/{id} [put]
func (c *AccountController) updateInvoice(ctx *gin.Context) error {
	return nil
}

// GetInvoice godoc
// @Summary GET invoice
// @Description GET invoice
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/invoice/{id} [get]
func (c *AccountController) getInvoice(ctx *gin.Context) error {
	return nil
}

// GetPlans godoc
// @Summary GET plans
// @Description GET plans
// @Tags Account
// @Accept json
// @Produce json
// @Success 200 {object} gin.H
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/plans [get]
func (c *AccountController) getPlans(ctx *gin.Context) error {
	return nil
}

// GetSubscription godoc
// @Summary GET subscription
// @Description GET subscription
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/subscription/{id} [get]
func (c *AccountController) getSubscription(ctx *gin.Context) error {
	return nil
}

// UpgradePlan godoc
// @Summary Upgrade plan
// @Description Upgrade plan
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Param body body Dto.Subscription true "Subscription"
// @Success 200 {object} gin.H
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/subscription/{id} [put]
func (c *AccountController) upgradePlan(ctx *gin.Context) error {
	return nil
}

// CancelSubscription godoc
// @Summary Cancel subscription
// @Description Cancel subscription
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/subscription/{id} [delete]
func (c *AccountController) cancelSubscription(ctx *gin.Context) error {
	return nil
}

// DeleteAccount godoc
// @Summary Delete account
// @Description Delete account
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/{id} [delete]
func (c *AccountController) deleteAccount(ctx *gin.Context) error {
	return nil
}

// AddAccount godoc
// @Summary Add account
// @Description Add account
// @Tags Account
// @Accept json
// @Produce json
// @Param body body Dto.AddAccountDto true "Account"
// @Success 200 {object} models.Account
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/add-account [post]
func (c *AccountController) addAccount(ctx *gin.Context) error {
	var accountDto *Dto.AddAccountDto
	if err := ctx.BindJSON(&accountDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
		})
		return err
	}

	// Validate request
	//if err := c.utility.Validate(accountDto, "accountName, accountType"); err != nil {
	//	return ctx.Status(400).JSON(gin.H{
	//		"status":  "error",
	//		"message": err.Error(),
	//	})
	//}

	// Create account
	accountData, err := c.store.CreateBusinessAccount(accountDto, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Account created successfully",
		"data":    accountData,
	})

	return nil
}

// GetAccount godoc
// @Summary GET account
// @Description GET account
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} models.Account "Account"
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/{accountId} [get]
func (c *AccountController) getUserAccount(ctx *gin.Context) error {
	accountId := ctx.Param("accountId")
	if accountId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
		})
		return errors.New("account id is required")
	}

	ctx.Set("accountId", accountId)
	ctx.Set("accountId", accountId)

	// GET account
	accountData, err := c.store.GetAccount(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Account retrieved successfully",
		"data":    accountData,
	})

	return nil
}

// MakeAccountDefault godoc
// @Summary Make account default
// @Description Make account default
// @Tags Account
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} Dto.UserAccountResponse
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/{accountId}/make-default [put]
func (c *AccountController) makeAccountDefault(ctx *gin.Context) error {
	accountId := ctx.Param("accountId")
	if accountId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
		})
		return errors.New("account id is required")
	}

	// Make account default
	account, err := c.store.MakeAccountDefault(accountId, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Account made default successfully",
		"data":    account,
	})

	return nil
}

// @Summary Follow an account
// @Description Follow another account by ID
// @Tags Followers
// @Accept json
// @Produce json
// @Param id path string true "Follower Account ID"
// @Param following_id formData string true "Following Account ID"
// @Success 200 {object} map[string]string "Successfully followed account"
// @Failure 400 {object} map[string]string "Bad Request"
// @Router /api/v1/accounts/{id}/follow [post]
func (c *AccountController) followAccount(ctx *gin.Context) error {
	followerID := ctx.Param("id")
	followingID := ctx.Param("following_id")

	err := c.store.Follow(followerID, followingID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully followed account"})
	return nil
}

// @Summary Unfollow an account
// @Description Unfollow another account by ID
// @Tags Followers
// @Accept json
// @Produce json
// @Param id path string true "Follower Account ID"
// @Param following_id formData string true "Following Account ID"
// @Success 200 {object} map[string]string "Successfully unfollowed account"
// @Failure 400 {object} map[string]string "Bad Request"
// @Router /api/v1/accounts/{id}/unfollow [post]
func (c *AccountController) unfollowAccount(ctx *gin.Context) error {
	followerID := ctx.Param("id")
	followingID := ctx.Param("following_id")

	err := c.store.Unfollow(followerID, followingID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully unfollowed account"})
	return nil
}

// @Summary GET followers
// @Description GET the list of accounts following the specified account
// @Tags Followers
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {array} models.Account "Successfully retrieved followers"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/v1/accounts/{id}/followers [get]
func (c *AccountController) getFollowers(ctx *gin.Context) error {
	accountID := ctx.Param("id")

	followers, err := c.store.ListFollowers(accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"followers": followers})
	return nil
}

// @Summary GET following
// @Description GET the list of accounts the specified account is following
// @Tags Followers
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {array} models.Account "Successfully retrieved following accounts"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/v1/accounts/{id}/following [get]
func (c *AccountController) getFollowing(ctx *gin.Context) error {
	accountID := ctx.Param("id")

	following, err := c.store.ListFollowing(accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"following": following})
	return nil
}
