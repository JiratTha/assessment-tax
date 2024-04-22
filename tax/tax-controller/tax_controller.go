package tax_controller

import (
	struc "github.com/JiratTha/assessment-tax/Personnel/Personnel-model"
	"github.com/JiratTha/assessment-tax/tax/calculation"
	"github.com/labstack/echo/v4"
	"net/http"
)

func TaxCalculationsPost(c echo.Context) error {
	var personnelIncome struc.Personnel
	if err := c.Bind(&personnelIncome); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	totalTax := calculation.TaxCalculation(personnelIncome.TotalIncome)
	totalTax.Tax -= personnelIncome.Wht
	taxResponse := struc.TaxResponse(totalTax)
	return c.JSON(http.StatusOK, taxResponse)
}
