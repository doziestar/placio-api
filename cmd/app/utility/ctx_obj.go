package utility

import (
	"github.com/gin-gonic/gin"
)

func GetContextData[TData string | interface{}](ctx *gin.Context, key string) (TData, error) {
	data, _ := ctx.Get(key)
	return data, nil
}
