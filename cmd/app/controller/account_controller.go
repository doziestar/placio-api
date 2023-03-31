package controller

//https://docs.gofiber.io/guide/validation

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"placio-app/Dto"
	"placio-app/middleware"
	"placio-app/service"
	"placio-app/utility"
	"placio-pkg/logger"
)

type AccountController struct {
	store service.IAccountService
}

func NewAccountController(store service.IAccountService) *AccountController {
	return &AccountController{store: store}
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
	accountGroup.Post("/create-account", utility.Use(c.createAccount))
	accountGroup.Post("/plan", middleware.Verify("owner"), utility.Use(c.plan))
	accountGroup.Patch("/plan", middleware.Verify("owner"), utility.Use(c.updatePlan))
	accountGroup.Get("/", middleware.Verify("owner"), utility.Use(c.getAccounts))
	accountGroup.Get("/card", middleware.Verify("owner"), utility.Use(c.getAccount))
	accountGroup.Patch("/card", middleware.Verify("owner"), utility.Use(c.updateInvoice))
	accountGroup.Get("/invoice", middleware.Verify("owner"), utility.Use(c.getInvoice))
	accountGroup.Get("/plans", middleware.Verify("owner"), utility.Use(c.getPlans))
	accountGroup.Get("/subscription", middleware.Verify("owner"), utility.Use(c.getSubscription))
	accountGroup.Post("/upgrade", middleware.Verify("owner"), utility.Use(c.upgradePlan))
	accountGroup.Delete("/", middleware.Verify("owner"), utility.Use(c.deleteAccount))

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
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/v1/accounts [post]
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
	if err := validate(data.Email, data.Name, data.Password); err != nil {
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

func validate(email string, name string, password string) error {
	if email == "" {
		return errors.New("Email is required")
	}
	if name == "" {
		return errors.New("Name is required")
	}
	if password == "" {
		return errors.New("Password is required")
	}
	if len(password) < 8 {
		return errors.New("Password must be at least 8 characters")
	}

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
// @Failure 400 {object} map[string]string "inputError": "plan", "message": "Plan is required"
// @Failure 500 {object} map[string]string "error": "Internal Server Error"
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
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 404 {object} map[string]interface{} "Not Found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
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
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /accounts [get]
func (c *AccountController) getAccounts(ctx *fiber.Ctx) error {
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
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /accounts/{id} [get]
func (c *AccountController) getAccount(ctx *fiber.Ctx) error {
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
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
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
// @Failure 400 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
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
// @Failure 400 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
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
// @Failure 400 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
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
// @Failure 400 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
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
// @Failure 400 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
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
// @Failure 400 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /accounts/{id} [delete]
func (c *AccountController) deleteAccount(ctx *fiber.Ctx) error {
	return nil
}
