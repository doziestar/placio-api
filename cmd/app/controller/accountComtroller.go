package controller

func Create(c *fiber.Ctx) error {
	data := new(struct {
		Email           string `json:"email"`
		Name            string `json:"name"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
		Plan            string `json:"plan"`
	})

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	// validate input
	if err := validate(data.Email, data.Name, data.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// confirm_password field is a dummy field to prevent bot signups
	if data.ConfirmPassword != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Registration denied",
		})
	}

	// check if user has already registered an account
	userData, err := user.Get(data.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

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
			if err := user.SavePassword(userData.ID, data.Password); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Internal Server Error",
				})
			}
		}

		c.Locals("duplicate_user", duplicateUser)
		c.Locals("has_password", hasPassword)
	}

	// create the account
	accountData, err := account.Create(data.Plan)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	c.Locals("account_id", accountData.ID)

	// create the user and assign to account
	userData, err = user.Create(data.Email, data.Name, accountData.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	if err := user.AddToAccount(userData.ID, accountData.ID, "owner"); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	// send welcome email
	if err := mail.Send(userData.Email, "new-account", map[string]interface{}{
		"name": userData.Name,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	// authenticate the user
	return authController.Signup(c)
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

func Plan(c *fiber.Ctx) error {
	data := new(struct {
		Plan   string      `json:"plan"`
		Token  *stripe.TokenParams `json:"token,omitempty"`
		Stripe struct {
			Customer struct {
				ID string `json:"id,omitempty"`
			} `json:"customer,omitempty"`
			Subscription struct {
				ID    string `json:"id,omitempty"`
				Price int64  `json:"price,omitempty"`
			} `json:"subscription,omitempty"`
		} `json:"stripe,omitempty"`
	})

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	if data.Plan == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Plan is required",
		})
	}

	// check the plan exists
	plan, ok := settings.Plans[data.Plan]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Plan doesn't exist",
		})
	}

	accountData, err := account.Get(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No account with that ID",
		})
	}

	// process stripe subscription for non-free accounts
	// if a 2-factor payment hasn't occurred, create the stripe subscription
	if data.Plan != "free" {
		if data.Stripe == (struct {
			Customer struct {
				ID string "json:\"id,omitempty\""
			} "json:\"customer,omitempty\""
			Subscription struct {
				ID    string "json:\"id,omitempty\""
				Price int64  "json:\"price,omitempty\""
			} "json:\"subscription,omitempty\""
		}{}) {
			if data.Token == nil || data.Token.ID == "" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Please enter your credit card details",
				})
			}

			// create a stripe customer and subscribe them to a plan
			customer, err := stripe.CustomerCreate(
				accountData.OwnerEmail,
				data.Token.ID,
			)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			data.Stripe.Customer.ID = customer.ID

			subscription, err := stripe.CustomerSubscribe(
				customer.ID,
				data.Plan,
			)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			data.Stripe.Subscription.ID = subscription.ID

			// check for an incomplete payment that requires 2-factor authentication
			if subscription.LatestInvoice.PaymentIntent.Status == "requires_action" {
				log.Println("Stripe payment requires further action")

				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"requires_payment_action": true,
					"customer": fiber.Map{
						"id": data.Stripe.Customer.ID,
					},
					"subscription": fiber.Map{
						"id":    data.Stripe.Subscription.ID,
						"price": data.Stripe.Subscription.Price,
					},
					"client_secret": subscription.LatestInvoice.PaymentIntent.ClientSecret,
				})
			}
		}

		// stripe info hasn't been passed back as part of 2
		if data.Stripe.Customer.ID == "" || data.Stripe.Subscription.ID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Stripe customer or subscription ID is missing",
			})
		}

		// update the stripe subscription
		subscription, err := stripe.CustomerUpdateSubscription(
			data.Stripe.Customer.ID,
			data.Stripe.Subscription.ID,
			data.Plan,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		data.Stripe.Subscription.ID = subscription.ID
	}

	// update the account plan
	if err := account.UpdatePlan(accountData.ID, data.Plan, data.Stripe.Customer.ID, data.Stripe.Subscription.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})

	return nil
}

func getAccountData(accountID string) (*AccountData, error) {
	// implementation of account.get()
	// return an error if the account does not exist
	return nil, fmt.Errorf("Account does not exist")
}

func updateAccountData(accountID string, data *AccountData) error {
	// implementation of account.update()
	return nil
}

func sendMail(to, template string, content map[string]interface{}) error {
	// implementation of mail.send()
	return nil
}

func updatePlan(c *fiber.Ctx) error {
	data := new(struct {
		ID   string `json:"id"`
		Plan string `json:"plan"`
	})
	if err := c.BodyParser(data); err != nil {
		return err
	}

	// implementation of utility.validate()
	// return an error if the request body is invalid

	accountID := ""
	if c.Locals("permission") == "master" {
		accountID = data.ID
	} else {
		accountID = c.Locals("account").(string)
	}

	plan := Plan{}
	for _, p := range plans {
		if p.ID == data.Plan {
			plan = p
			break
		}
	}
	if plan.ID == "" {
		return fmt.Errorf("No plan with that ID")
	}

	accountData, err := getAccountData(accountID)
	if err != nil {
		return err
	}

	// user is upgrading from paid to free,
	// direct them to the upgrade view
	if accountData.Plan == "free" && plan.ID != "free" {
		if c.Locals("permission") == "master" {
			return fiber.NewError(fiber.StatusForbidden, "The account holder will need to enter their card details and upgrade to a paid plan.")
		} else {
			return c.Status(fiber.StatusPaymentRequired).JSON(fiber.Map{
				"message": "Please upgrade your account",
				"plan":    plan.ID,
			})
		}
	}

	if plan.ID == "free" {
		// user is downgrading - cancel the stripe subscription
		if accountData.StripeSubscriptionID != "" {
			subscription, err := sub.Get(accountData.StripeSubscriptionID, nil)
			if err != nil {
				return err
			}

			err = updateAccountData(accountID, &AccountData{
				StripeSubscriptionID: "",
				Plan:                plan.ID,
			})
			if err != nil {
				return err
			}

			if subscription.Status != "canceled" {
				_, err := sub.Cancel(accountData.StripeSubscriptionID, nil)
				if err != nil {
					return err
				}
			}
		}
	} else {
		// user is switching to a different paid plan
		if accountData.StripeSubscriptionID != "" {
			// check for an incomplete payment that requires 2-factor authentication
			subscription, err := sub.Get(accountData.StripeSubscriptionID, nil)
			if err != nil {
				return err
			}

			if subscription.LatestInvoice.PaymentIntent.Status == "requires_action" {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"requires_payment_action": true,
					"customer": fiber.Map{
						"id": accountData.StripeCustomerID,
					},
					"subscription": fiber.Map{
						"id":    accountData.StripeSubscriptionID,
						"price": plan.Price,
					},
					"client_secret": subscription.LatestInvoice.PaymentIntent.ClientSecret,
				})
			}

			// update the stripe subscription
			subscription, err = sub.Update(accountData.StripeSubscriptionID, &stripe.SubscriptionParams{
				Items: []*stripe.SubscriptionItemsParams{
					{
						ID:   subscription.Items.Data[0].ID,
						Plan: stripe.String(plan.ID),
					},
				},
			})
			if err != nil {
				return err
			}

			err = updateAccountData(accountID, &AccountData{
				StripeSubscriptionID: subscription.ID,
				Plan:                plan.ID,
			})
			if err != nil {
				return err
			}

			// send an email to the account holder
			err = sendMail(accountData.Email, "plan-upgraded", fiber.Map{
				"plan": plan.Name,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}