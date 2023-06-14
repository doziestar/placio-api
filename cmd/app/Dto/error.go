package Dto

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// ErrorDTO represents the structure of an error response.
type ErrorDTO struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorDto struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type PermissionDTO struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessDTO struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// NewErrorDTO creates a new ErrorDTO instance with the provided status and message.
func NewErrorDTO(status int, message string) *ErrorDTO {
	return &ErrorDTO{Status: status, Message: message}
}

// BadRequest creates a new ErrorDTO instance for a 400 Bad Request error.
func BadRequest(message string) *ErrorDTO {
	return NewErrorDTO(fiber.StatusBadRequest, message)
}

// NotFound creates a new ErrorDTO instance for a 404 Not Found error.
func NotFound(message string) *ErrorDTO {
	return NewErrorDTO(fiber.StatusNotFound, message)
}

// InternalServerError creates a new ErrorDTO instance for a 500 Internal Server Error.
func InternalServerError(message string) *ErrorDTO {
	return NewErrorDTO(fiber.StatusInternalServerError, message)
}

// JSON returns the JSON representation of the ErrorDTO.
func (e *ErrorDTO) JSON() ([]byte, error) {
	return json.Marshal(e)
}
