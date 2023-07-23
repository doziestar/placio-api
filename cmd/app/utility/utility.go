package utility

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math/big"
	"net/http"
	"placio-app/ent"
	appErrors "placio-app/errors"
	"placio-app/models"
	"reflect"
	"strings"

	sentry "github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

type IUtility interface {
	Validate(form map[string]interface{}, fields []string) error
	//ValidateEmail(email string) error
	//ValidatePassword(password string) error
	//ValidatePhone(phone string) error
	//ValidateName(name string) error
	//ValidateAddress(address string) error
	//ValidateDate(date string) error
	//ValidateTime(time string) error
	//ValidatePasswordMatch(password string, confirmPassword string) error
	//ValidatePasswordStrength(password string) error
}

type Utility struct {
}

func NewUtility() *Utility {
	return &Utility{}
}
func (u *Utility) Validate(form map[string]interface{}, fields []string) error {
	// sanitise the input
	for f, v := range form {
		// sanitise
		if s, ok := v.(string); ok && strings.Contains(s, "<script>") {
			form[f] = strings.ReplaceAll(strings.ReplaceAll(s, "<script>", ""), "</script>", "")
		}
	}

	// check for required fields
	if len(fields) > 0 {
		for _, f := range fields {
			if _, ok := form[f]; !ok || form[f] == nil {
				// field is required
				return fmt.Errorf("%s field is required", f)
			}
		}
	}

	return nil

}

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return fmt.Errorf("invalid email address")
	}
	return nil
}
func ValidatePassword(password string) error {
	// check password length
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	// check password complexity
	if !strings.ContainsAny(password, "0123456789") {
		return fmt.Errorf("password must contain at least one number")
	}

	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	return nil
}

func ValidatePasswordMatch(password, confirmPassword string) error {
	if password != confirmPassword {
		return fmt.Errorf("passwords do not match")
	}
	return nil
}

func Assert(data interface{}, err string, input map[string]interface{}) bool {
	if data == nil {
		m := map[string]interface{}{"message": err}
		if input != nil {
			for k, v := range input {
				m[k] = v
			}
		}
		panic(m)
	}
	return true
}

//
//func Use(fn func(*fiber.Ctx) error) fiber.Handler {
//	//fmt.Println("Entering utility.Use function")
//	//logger.Info(context.Background(), "middleware.Use")
//	defer sentry.Recover()
//	//defer sentry.Flush(2 * time.Second)
//	return func(c *fiber.Ctx) error {
//		err := fn(c)
//		if err != nil {
//			sentry.CaptureException(err)
//			return c.Next()
//		}
//		return nil
//	}
//}

type MyError struct {
	StatusCode int
	Msg        string
}

func (e *MyError) Error() string {
	return e.Msg
}

func Use(fn func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer sentry.Recover()
		err := fn(c)
		if err != nil {
			// Enhanced logging
			sentry.CaptureException(err)

			// Check if the error is of type AppError
			if appErr, ok := err.(*appErrors.AppError); ok {
				// It's an AppError, now match on its underlying error
				switch {
				case errors.Is(appErr.Unwrap(), appErrors.ErrInvalid):
					c.JSON(http.StatusBadRequest, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.ErrUnauthorized):
					c.JSON(http.StatusUnauthorized, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.ErrForbidden):
					c.JSON(http.StatusForbidden, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.ErrNotFound), ent.IsNotFound(appErr.Unwrap()):
					c.JSON(http.StatusNotFound, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.ErrConflict), ent.IsConstraintError(appErr.Unwrap()):
					c.JSON(http.StatusConflict, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.ErrUnprocessable):
					c.JSON(http.StatusUnprocessableEntity, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.ErrAlreadyExists):
					c.JSON(http.StatusConflict, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.ErrInternal):
					c.JSON(http.StatusInternalServerError, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.ErrTimeout):
					c.JSON(http.StatusRequestTimeout, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.InvalidItemType):
					c.JSON(http.StatusBadRequest, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.ErrInvalidInput):
					c.JSON(http.StatusBadRequest, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.IDMissing):
					c.JSON(http.StatusBadRequest, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.IDMismatch):
					c.JSON(http.StatusBadRequest, gin.H{"error": appErr.Error()})
				case errors.Is(appErr.Unwrap(), appErrors.ErrTemporaryDisabled):
					c.JSON(http.StatusServiceUnavailable, gin.H{"error": appErr.Error()})
				default:
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				}
				return
			}

			// If it's not an AppError, check if it's an ent error
			switch {
			case ent.IsNotFound(err):
				c.JSON(http.StatusNotFound, gin.H{"error": "Resource Not Found"})
			case ent.IsConstraintError(err):
				c.JSON(http.StatusConflict, gin.H{"error": "Constraint Error"})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
			return
		}
		c.Next()
	}
}

// StructToMap converts a struct to a map, omitting fields that are nil.
func StructToMap(v interface{}) (map[string]interface{}, error) {
	var m map[string]interface{}
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// MergeMaps mergeMaps merges two map[string]interface{} objects.
func MergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for k, v := range map1 {
		merged[k] = v
	}
	for k, v := range map2 {
		merged[k] = v
	}
	return merged
}

// ProcessResponse processes the response from a service.
func ProcessResponse(data interface{}, status string, message, nextPageToken string) gin.H {
	return gin.H{
		"data":          data,
		"status":        status,
		"message":       message,
		"nextPageToken": nextPageToken,
	}
}

// ProcessError processes an error response from a service.
func ProcessError(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

// SplitString splits a string into a slice of strings.
func SplitString(s string, sep string) []string {
	return strings.Split(s, sep)
}

// ConvertToInterfaceSlice converts a slice of any type to a slice of interface{}.
func ConvertToInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

// RemoveSensitiveInfo removes sensitive information from a user object.
func RemoveSensitiveInfo(user *models.User) *models.User {
	if user.Auth0Data == nil {
		return user
	}

	// Create a deep copy of the user
	newUser := *user
	newAuth0Data := *user.Auth0Data

	// Remove sensitive info from the top level of Auth0Data
	newAuth0Data.Email = nil
	newAuth0Data.LastIP = nil
	newAuth0Data.LoginsCount = nil

	for i := range newAuth0Data.Identities {
		newAuth0Data.Identities[i].AccessToken = nil
	}

	// Assign the new Auth0Data to the new user
	newUser.Auth0Data = &newAuth0Data

	return &newUser

	//auth0DataVal := reflect.ValueOf(user.Auth0Data).Elem()
	//
	//sensitiveFields := []string{"Email", "LastIP", "LoginsCount"}
	//for _, field := range sensitiveFields {
	//	fieldVal := auth0DataVal.FieldByName(field)
	//	if fieldVal.IsValid() && fieldVal.CanSet() {
	//		switch fieldVal.Kind() {
	//		case reflect.String:
	//			fieldVal.SetString("")
	//		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	//			fieldVal.SetInt(0)
	//			// Add other cases as necessary for other field types
	//		}
	//	}
	//}
	//
	//// We also need to handle the nested "identities" slice
	//identitiesVal := auth0DataVal.FieldByName("Identities")
	//if identitiesVal.IsValid() && identitiesVal.CanSet() {
	//	for i := 0; i < identitiesVal.Len(); i++ {
	//		identityVal := identitiesVal.Index(i)
	//		if identityVal.Kind() == reflect.Ptr {
	//			identityVal = identityVal.Elem() // Dereference the pointer
	//		}
	//		fieldVal := identityVal.FieldByName("AccessToken")
	//		if fieldVal.IsValid() && fieldVal.CanSet() {
	//			if fieldVal.Kind() == reflect.Ptr {
	//				fieldVal = fieldVal.Elem() // Dereference the pointer
	//			}
	//			if fieldVal.Kind() == reflect.String {
	//				fieldVal.SetString("")
	//			}
	//		}
	//	}
	//}

	//return
}

func GenerateID() string {
	uid, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return uid.String()
}

func GenerateRandomUsername() string {
	// Generate a random number
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1e14))
	if err != nil {
		panic(err) // handle error
	}

	// Combine the random number with "user" prefix
	return fmt.Sprintf("user%s", randomNumber.String())
}
