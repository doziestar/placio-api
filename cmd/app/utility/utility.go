package utility

import (
	"fmt"
	"strings"

	sentry "github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
)

func Validate(form map[string]interface{}, fields []string) error {
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

func Use(fn func(*fiber.Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := fn(c)
		if err != nil {
			sentry.CaptureException(err)
			return c.Next()
		}
		return nil
	}
}
