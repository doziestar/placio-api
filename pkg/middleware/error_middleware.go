package middleware

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/ent"
	appErrors "placio-pkg/errors"
)

func ErrorMiddleware(fn func(*gin.Context) error) gin.HandlerFunc {
	errorMappings := map[error]func(appError *appErrors.AppError, c *gin.Context){
		appErrors.ErrInvalid:           respondBadRequest,
		appErrors.ErrUnauthorized:      respondUnauthorized,
		appErrors.ErrForbidden:         respondForbidden,
		appErrors.ErrNotFound:          respondNotFound,
		appErrors.ErrConflict:          respondConflict,
		appErrors.ErrUnprocessable:     respondUnprocessableEntity,
		appErrors.ErrAlreadyExists:     respondConflict,
		appErrors.ErrInternal:          respondInternalServerError,
		appErrors.ErrTimeout:           respondRequestTimeout,
		appErrors.InvalidItemType:      respondBadRequest,
		appErrors.ErrInvalidInput:      respondBadRequest,
		appErrors.IDMissing:            respondBadRequest,
		appErrors.IDMismatch:           respondBadRequest,
		appErrors.ErrTemporaryDisabled: respondServiceUnavailable,
	}

	return func(c *gin.Context) {
		defer sentry.Recover()
		err := fn(c)
		if err != nil {
			sentry.CaptureException(err)

			// Check if the error is of type AppError
			if appErr, ok := err.(*appErrors.AppError); ok {
				for knownError, responseFn := range errorMappings {
					if errors.Is(appErr.Unwrap(), knownError) {
						responseFn(appErr, c)
						return
					}
				}
				if ent.IsNotFound(appErr.Unwrap()) {
					respondNotFound(appErr, c)
					return
				}
				if ent.IsConstraintError(appErr.Unwrap()) {
					respondConflict(appErr, c)
					return
				}
				respondInternalServerError(appErr, c)
				return
			}

			// If it's not an AppError, check if it's an ent error
			if ent.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Resource Not Found"})
				return
			}
			if ent.IsConstraintError(err) {
				c.JSON(http.StatusConflict, gin.H{"error": "Constraint Error"})
				return
			}
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: %s", err.Error()))
			return
		}
		c.Next()
	}
}

func respondBadRequest(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func respondUnauthorized(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
}

func respondForbidden(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
}

func respondNotFound(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
}

func respondConflict(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
}

func respondUnprocessableEntity(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
}

func respondInternalServerError(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func respondRequestTimeout(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, gin.H{"error": err.Error()})
}

func respondServiceUnavailable(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
}

func respondNotImplemented(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
}

func respondBadGateway(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
}

func respondGatewayTimeout(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, gin.H{"error": err.Error()})
}

func respondHTTPVersionNotSupported(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusHTTPVersionNotSupported, gin.H{"error": err.Error()})
}

func respondVariantAlsoNegotiates(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusVariantAlsoNegotiates, gin.H{"error": err.Error()})
}

func respondInsufficientStorage(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusInsufficientStorage, gin.H{"error": err.Error()})
}

func respondLoopDetected(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusLoopDetected, gin.H{"error": err.Error()})
}

func respondNotExtended(err *appErrors.AppError, c *gin.Context) {
	c.JSON(http.StatusNotExtended, gin.H{"error": err.Error()})
}
