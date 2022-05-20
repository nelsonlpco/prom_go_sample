package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/nelsonlpco/classic_cc_problens/internal/api/models"
	"github.com/nelsonlpco/classic_cc_problens/internal/infra/metrics"
	"github.com/nelsonlpco/classic_cc_problens/internal/shared"
)

const FibonacciPath = "/fibonacci/:number"

func (h Handler) FibonacciCalc(c echo.Context) error {
	span := shared.Tracer.Start(c.Request().Context(), "fibonacci")
	defer span.End()

	numberParam := c.Param("number")
	number, err := strconv.ParseUint(numberParam, 10, 64)
	if err != nil {
		span.Error(err)
		log.Println(fmt.Printf("invalid param %s error %v", number, err))
		return err
	}

	span.WithAttribute("input", fmt.Sprintf("%d", number))
	result := h.fibonacci.Calc(number)
	span.WithAttribute("value", fmt.Sprintf("%d", result))

	metrics.AddCounter(c.Request().Context(), "fibonnaci.request", "requests to fibonacci endpoint")

	c.JSON(http.StatusOK, &models.DefaultResponseModel{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("Fibonnaci %d = %d", number, result),
	})

	return nil
}
