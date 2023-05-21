package utility

//
//func removeSensitiveFields(user *models.User) (map[string]interface{}, error) {
//	userMap, err := lo.Marshal(user)
//	if err != nil {
//		return nil, err
//	}
//
//	delete(userMap, "Password")
//
//	return userMap, nil
//}

// func RemoveSensitiveFields(user *models.User) (map[string]interface{}, error) {
// 	userMap := make(map[string]interface{})
// 	err := mapstructure.Decode(user, &userMap)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// remove sensitive fields from accounts
// 	accounts := make([]interface{}, len(user.Accounts))
// 	for i, account := range user.Accounts {
// 		accountMap := make(map[string]interface{})
// 		err := mapstructure.Decode(account, &accountMap)
// 		if err != nil {
// 			return nil, err
// 		}

// 		delete(accountMap, "CreatedAt")
// 		delete(accountMap, "UpdatedAt")
// 		delete(accountMap, "DeletedAt")
// 		delete(accountMap, "StripeCustomerID")
// 		delete(accountMap, "StripeSubscriptionID")
// 		delete(accountMap, "StripeSubscriptionStatus")
// 		delete(accountMap, "PayStackSubscriptionID")
// 		delete(accountMap, "PayStackSubscriptionStatus")

// 		accounts[i] = accountMap
// 	}
// 	userMap["Accounts"] = accounts

// 	delete(userMap, "Password")
// 	delete(userMap, "CreatedAt")
// 	delete(userMap, "UpdatedAt")
// 	delete(userMap, "DeletedAt")
// 	delete(userMap, "IP")
// 	delete(userMap, "StripeCustomerID")
// 	delete(userMap, "StripeSubscriptionID")
// 	delete(userMap, "StripeSubscriptionStatus")
// 	delete(userMap, "PayStackSubscriptionID")
// 	delete(userMap, "PayStackSubscriptionStatus")

// 	return userMap, nil
// }
