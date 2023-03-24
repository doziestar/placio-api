package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"placio-app/Dto"
	"placio-app/database"
	"placio-app/models"
	"placio-pkg/logger"

	"github.com/gofiber/fiber/v2"
)

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
// @Tags accounts
// @Accept json
// @Produce json
// @Param SignUpDto body Dto.SignUpDto true "Sign Up Data"
// @Success 200 {object} models.Account "Successfully created account"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /api/v1/accounts [post]
func CreateAccount(c *fiber.Ctx) error {
	data := new(Dto.SignUpDto)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}
	log.Println("CreateAccount", data)

	// validate input
	if err := validate(data.Email, data.Name, data.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// confirm_password field is a dummy field to prevent bot signups
	if data.ConfirmPassword == "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You need to confirm your password",
		})
	}

	user := new(models.User)

	// check if user has already registered an account
	userData, _ := user.GetByEmail(data.Email, database.DB)

	//if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	//	// continue if user doesn't exist
	//} else {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Internal Server Error",
	//	})
	//}

	logger.Info(context.Background(), fmt.Sprintf("userData: %v", userData))

	if userData != nil {
		// user already owns an account
		if userData.Permission == "owner" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"inputError": "email",
				"message":    "You have already registered an account",
			})
		}

		// flag for authController to notify onboarding ui
		// that the user's existing account was used
		duplicateUser := true
		hasPassword := userData.HasPassword

		// save the new password if it exists and user doesn't have one
		if !hasPassword && data.Password != "" {
			if err := user.SavePassword(userData.UserID, data.Password); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Internal Server Error",
				})
			}
		}

		c.Locals("duplicate_user", duplicateUser)
		c.Locals("has_password", hasPassword)
	}
	//permission := func() string {
	//	if userData != nil {
	//		return userData.Permission
	//	}
	//	return "owner"
	//}()
	logger.Info(context.Background(), "CreateAccount")

	//account := new(models.Account)

	// create the user and assign to account
	newUser, err := user.CreateUser(*data, c, database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	//// create the account
	//accountData, err := account.CreateAccount(newUser.UserID, permission, database.DB)
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Internal Server Error",
	//	})
	//}
	//c.Locals("account_id", accountData.ID)
	//
	//if err := user.AddToAccount(userData.UserID, accountData.ID, "owner"); err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Internal Server Error",
	//	})
	//}

	mail := new(models.EmailContent)
	// send welcome email
	if err := mail.Send(newUser.Email, "new-account", userData.ToJson()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	// authenticate the user
	//return authController.Signup(c)
	return nil
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

//func Plan(c *fiber.Ctx) error {
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
//	return nil
//}

func SendMail(to, template string, content map[string]interface{}) error {
	// implementation of mail.send()
	return nil
}

//func UpdatePlan(c *fiber.Ctx) error {
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
//	return nil
//}
