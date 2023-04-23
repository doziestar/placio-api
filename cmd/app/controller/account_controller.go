package controller

//https://docs.gofiber.io/guide/validation

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"placio-app/Dto"
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

func requestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Printf("Request: %s %s\n", c.Method(), c.Path())
		return c.Next()
	}
}

func (c *AccountController) RegisterRoutes(app fiber.Router) {
	//app.Use(requestLogger())
	accountGroup := app.Group("/accounts")
	accountGroup.Get("/", middleware.Verify("user"), utility.Use(c.getUserAccount))
	accountGroup.Post("/create-account", utility.Use(c.createAccount))
	accountGroup.Post("/:accountId/switch-account/", middleware.Verify("user"), utility.Use(c.switchAccount))
	accountGroup.Post("/:accountId/make-default/", middleware.Verify("user"), utility.Use(c.makeAccountDefault))
	accountGroup.Post("/add-account", middleware.Verify("user"), utility.Use(c.addAccount)) // add account to owner
	accountGroup.Post("/plan", middleware.Verify("owner"), utility.Use(c.plan))
	accountGroup.Patch("/plan", middleware.Verify("owner"), utility.Use(c.updatePlan))
	accountGroup.Get("/get-user-accounts", middleware.Verify("user"), utility.Use(c.getAccounts))
	accountGroup.Get("/get-user-active-account", middleware.Verify("user"), utility.Use(c.getUserActiveAccount))
	accountGroup.Get("/:accountId", middleware.Verify("user"), utility.Use(c.getUserAccount))
	accountGroup.Patch("/card", middleware.Verify("owner"), utility.Use(c.updateInvoice))
	accountGroup.Get("/invoice", middleware.Verify("owner"), utility.Use(c.getInvoice))
	accountGroup.Get("/plans", middleware.Verify("owner"), utility.Use(c.getPlans))
	accountGroup.Get("/subscription", middleware.Verify("owner"), utility.Use(c.getSubscription))
	accountGroup.Post("/upgrade", middleware.Verify("owner"), utility.Use(c.upgradePlan))
	accountGroup.Delete("/", middleware.Verify("owner"), utility.Use(c.deleteAccount))
	accountGroup.Post("/:id/follow", middleware.Verify("user"), utility.Use(c.followAccount))
	accountGroup.Post("/:id/unfollow", middleware.Verify("user"), utility.Use(c.unfollowAccount))
	accountGroup.Get("/:id/followers", middleware.Verify("user"), utility.Use(c.getFollowers))
	accountGroup.Get("/:id/following", middleware.Verify("user"), utility.Use(c.getFollowing))

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
func (c *AccountController) createAccount(ctx *fiber.Ctx) error {
	fmt.Println("Entering createAccount function")
	data := new(Dto.SignUpDto)

	logger.Info(context.Background(), fmt.Sprintf("data: %v", data))
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}
	log.Println("CreateAccount", data)

	// validate input
	if err := utility.Validate(data.Email, data.Name, data.Password, data.Username); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// confirm_password field is a dummy field to prevent bot signups
	if data.ConfirmPassword == "" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You need to confirm your password",
		})
	}

	response, err := c.store.CreateUserAccount(data, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
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
func (c *AccountController) switchAccount(ctx *fiber.Ctx) error {
	accountId := ctx.Params("accountId")
	userId := ctx.Locals("user").(string)
	response, err := c.store.SwitchUserAccount(accountId, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	data, err := utility.RemoveSensitiveFields(response)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully switched account",
		"data":    data,
		"status":  "success",
	})
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
func (c *AccountController) plan(ctx *fiber.Ctx) error {
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
	//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//			"error": "Bad request",
	//		})
	//	}
	//
	//	if data.Plan == "" {
	//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//			"error": "Plan is required",
	//		})
	//	}
	//
	//	// check the plan exists
	//	plan, ok := settings.Plans[data.Plan]
	//	if !ok {
	//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//			"error": "Plan doesn't exist",
	//		})
	//	}
	//
	//	accountData, err := account.Get(c)
	//	if err != nil {
	//		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
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
	//				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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
	//				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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
	//				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//					"error": err.Error(),
	//				})
	//			}
	//			data.Stripe.Subscription.ID = subscription.ID
	//
	//			// check for an incomplete payment that requires 2-factor authentication
	//			if subscription.LatestInvoice.PaymentIntent.Status == "requires_action" {
	//				log.Println("Stripe payment requires further action")
	//
	//				return c.Status(fiber.StatusOK).JSON(fiber.Map{
	//					"requires_payment_action": true,
	//					"customer": fiber.Map{
	//						"id": data.Stripe.Customer.ID,
	//					},
	//					"subscription": fiber.Map{
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
	//			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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
	//			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//				"error": err.Error(),
	//			})
	//		}
	//		data.Stripe.Subscription.ID = subscription.ID
	//	}
	//
	//	// update the account plan
	//	if err := account.UpdatePlan(accountData.ID, data.Plan, data.Stripe.Customer.ID, data.Stripe.Subscription.ID); err != nil {
	//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//			"error": "Internal Server Error",
	//		})
	//	}
	//
	//	c.Status(fiber.StatusOK).JSON(fiber.Map{
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
func (c *AccountController) updatePlan(ctx *fiber.Ctx) error {
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
	//			return c.Status(fiber.StatusPaymentRequired).JSON(fiber.Map{
	//				"message": "Please upgrade your account",
	//				"plan":    plan.ID,
	//			})
	//		}
	//	}
	//
	//	if plan.ID == "free" {
	//		// user is downgrading - cancel the stripe subscription
	//		if accountData.StripeSubscriptionID != "" {
	//			subscription, err := sub.Get(accountData.StripeSubscriptionID, nil)
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
	//			subscription, err := sub.Get(accountData.StripeSubscriptionID, nil)
	//			if err != nil {
	//				return err
	//			}
	//
	//			if subscription.LatestInvoice.PaymentIntent.Status == "requires_action" {
	//				return c.Status(fiber.StatusOK).JSON(fiber.Map{
	//					"requires_payment_action": true,
	//					"customer": fiber.Map{
	//						"id": accountData.StripeCustomerID,
	//					},
	//					"subscription": fiber.Map{
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
	//			err = sendMail(accountData.Email, "plan-upgraded", fiber.Map{
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
// @Summary Get account
// @Description Get account
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
func (c *AccountController) getAccounts(ctx *fiber.Ctx) error {

	accounts, err := c.store.GetAccounts(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"accounts": accounts,
	})
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
	//		s, err := sub.Get(accountData.StripeSubscriptionID, nil)
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
	return ctx.JSON(fiber.Map{
		//"account": account,
		"status": "success",
		//			"plan":             plan,
	})
}

// GetAccount godoc
// @Summary Get account
// @Description Get account
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
func (c *AccountController) getUserActiveAccount(ctx *fiber.Ctx) error {
	log.Println("Get user active account")
	account, err := c.store.GetAccount(ctx)
	if err != nil {
		return fiber.ErrNotFound
	}

	return ctx.JSON(fiber.Map{
		"account": account,
	})
}

// UpdateInvoice godoc
// @Summary Update invoice
// @Description Update invoice
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Param body body Dto.Invoice true "Invoice"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/invoice/{id} [put]
func (c *AccountController) updateInvoice(ctx *fiber.Ctx) error {
	return nil
}

// GetInvoice godoc
// @Summary Get invoice
// @Description Get invoice
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/invoice/{id} [get]
func (c *AccountController) getInvoice(ctx *fiber.Ctx) error {
	return nil
}

// GetPlans godoc
// @Summary Get plans
// @Description Get plans
// @Tags Account
// @Accept json
// @Produce json
// @Success 200 {object} fiber.Map
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/plans [get]
func (c *AccountController) getPlans(ctx *fiber.Ctx) error {
	return nil
}

// GetSubscription godoc
// @Summary Get subscription
// @Description Get subscription
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/subscription/{id} [get]
func (c *AccountController) getSubscription(ctx *fiber.Ctx) error {
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
// @Success 200 {object} fiber.Map
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/subscription/{id} [put]
func (c *AccountController) upgradePlan(ctx *fiber.Ctx) error {
	return nil
}

// CancelSubscription godoc
// @Summary Cancel subscription
// @Description Cancel subscription
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/subscription/{id} [delete]
func (c *AccountController) cancelSubscription(ctx *fiber.Ctx) error {
	return nil
}

// DeleteAccount godoc
// @Summary Delete account
// @Description Delete account
// @Tags Account
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 404 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /accounts/{id} [delete]
func (c *AccountController) deleteAccount(ctx *fiber.Ctx) error {
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
func (c *AccountController) addAccount(ctx *fiber.Ctx) error {
	var accountDto *Dto.AddAccountDto
	if err := ctx.BodyParser(&accountDto); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
		})
	}

	// Validate request
	//if err := c.utility.Validate(accountDto, "accountName, accountType"); err != nil {
	//	return ctx.Status(400).JSON(fiber.Map{
	//		"status":  "error",
	//		"message": err.Error(),
	//	})
	//}

	// Create account
	accountData, err := c.store.CreateBusinessAccount(accountDto, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   accountData,
	})
}

// GetAccount godoc
// @Summary Get account
// @Description Get account
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
func (c *AccountController) getUserAccount(ctx *fiber.Ctx) error {
	accountId := ctx.Params("accountId")
	if accountId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Account ID is required",
		})
	}

	ctx.Locals("accountId", accountId)
	ctx.Set("accountId", accountId)

	// Get account
	accountData, err := c.store.GetAccount(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   accountData,
	})
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
func (c *AccountController) makeAccountDefault(ctx *fiber.Ctx) error {
	accountId := ctx.Params("accountId")
	if accountId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Account ID is required",
		})
	}

	// Make account default
	account, err := c.store.MakeAccountDefault(accountId, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   account,
	})
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
func (c *AccountController) followAccount(ctx *fiber.Ctx) error {
	followerID := ctx.Params("id")
	followingID := ctx.FormValue("following_id")

	err := c.store.Follow(followerID, followingID)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"status": "success"})
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
func (c *AccountController) unfollowAccount(ctx *fiber.Ctx) error {
	followerID := ctx.Params("id")
	followingID := ctx.FormValue("following_id")

	err := c.store.Unfollow(followerID, followingID)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"status": "success"})
}

// @Summary Get followers
// @Description Get the list of accounts following the specified account
// @Tags Followers
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {array} models.Account "Successfully retrieved followers"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/v1/accounts/{id}/followers [get]
func (c *AccountController) getFollowers(ctx *fiber.Ctx) error {
	accountID := ctx.Params("id")

	followers, err := c.store.ListFollowers(accountID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"followers": followers})
}

// @Summary Get following
// @Description Get the list of accounts the specified account is following
// @Tags Followers
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {array} models.Account "Successfully retrieved following accounts"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/v1/accounts/{id}/following [get]
func (c *AccountController) getFollowing(ctx *fiber.Ctx) error {
	accountID := ctx.Params("id")

	following, err := c.store.ListFollowing(accountID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"following": following})
}
