package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID                         string `gorm:"primaryKey"`
	Permission                 string
	AccountType                string
	AccountID                  string
	Onboarded                  bool
	Interests                  []string `gorm:"type:text;default:[]"`
	UserID                     string   `gorm:"column:user_id"`
	Plan                       string
	Active                     bool
	StripeCustomerID           string
	StripeSubscriptionID       string
	StripeSubscriptionStatus   string
	PayStackSubscriptionID     string
	PayStackSubscriptionStatus string
	PayStackCustomerID         string
}

/*
 * account.create()
 * create a new account and return the account id
 */
 func (a *Account) Create() error {
	db := getDB()

	a.ID = uuid.New().String()
	a.Name = "My Account"
	a.Active = true
	a.DateCreated = time.Now()

	result := db.Create(&a)
	return result.Error
}

/*
 * account.get()
 * get an account by email or id
 */
func (a *Account) Get(id string) error {
	db := getDB()

	result := db.Where("id = ?", id).First(&a)
	if result.Error != nil {
		return result.Error
	}

	// get owner info
	var userData struct {
		Name  string
		Email string
	}
	result = db.Table("user").
		Select("name, email").
		Where("account_id = ? AND (permission = ? OR permission = ?)",
			id, "owner", "master").
		Scan(&userData)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	if userData.Name != "" && userData.Email != "" {
		a.OwnerName = userData.Name
		a.OwnerEmail = userData.Email
	}

	return nil
}

/*
 * account.subscription()
 * get the subscription status for this account
 */
func (a *Account) Subscription(id string) (string, error) {
	db := getDB()

	result := db.Where("id = ?", id).First(&a)
	if result.Error != nil {
		return "", result.Error
	}

	if a.Plan == "free" {
		return "active", nil
	}

	if a.StripeSubID == "" {
		return "inactive", nil
	}

	subscription, err := stripe.Subscription(a.StripeSubID)
	if err != nil {
		return "", err
	}

	status := subscription.Status
	if status != "active" && status != "trialing" {
		// update account to inactive
		db.Model(&a).Update("active", false)
	}

	return status, nil
}

/*
 * account.update()
 * update the account profile
 */
func (a *Account) Update(data map[string]interface{}) error {
	db := getDB()

	result := db.Model(&a).Updates(data)
	return result.Error
}

/*
 * account.delete()
 * delete the account and all its users
 */
func (a *Account) Delete(id string) error {
	db := getDB()

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// delete all users
	result := tx.Where("account_id = ?", id).Delete(&User{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// delete account
	result = tx.Where("id = ?", id).Delete(&a)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return tx
}