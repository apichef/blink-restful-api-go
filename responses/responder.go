package responses

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ValidationError struct {
	Errors map[string]interface{} `json:"errors"`
}

func formatErrors(err error) ValidationError {
	res := ValidationError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		res.Errors[e.Field()] = fmt.Sprintf("%s is %s", e.Field(), e.Tag())
	}

	return res
}

func BadRequest(context *gin.Context, err error) {
	context.JSON(http.StatusUnprocessableEntity, formatErrors(err))
}

func Ok(context *gin.Context, obj interface{}) {
	context.JSON(http.StatusOK, obj)
}

func Created(context *gin.Context, obj interface{}) {
	context.JSON(http.StatusCreated, obj)
}

func NoContent(context *gin.Context) {
	context.JSON(http.StatusNoContent, nil)
}
