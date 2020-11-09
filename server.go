package main

import (
	"net/http"
	"os"
	"xcoin_rate_tracker/services"
	"xcoin_rate_tracker/strategies/arbitrage"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	os.Setenv("MARGIN", "-21")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/:exchange/:pair", getTickerExchangeRate)
	e.GET("/:exchange", getAllExchangeRates)
	e.GET("/rates/:pair", getAllRates)
	e.GET("/strategies/arbitrage", rateGaps)
	e.GET("/strategies/compare", compare)
	e.Logger.Fatal(e.Start(":1323"))
}

func getTickerExchangeRate(c echo.Context) error {
	exchange := c.Param("exchange")
	pair := c.Param("pair")
	ticker := services.ExchangeTickerRate(pair, exchange)

	return c.JSON(http.StatusOK, ticker)
}

func getAllExchangeRates(c echo.Context) error {
	exchange := c.Param("exchange")
	services.AllRates(exchange)
	return c.String(http.StatusOK, "All rates for "+exchange)
}

func getAllRates(c echo.Context) error {
	pair := c.Param("pair")
	rates := services.GetRates(pair, false)
	return c.JSON(http.StatusOK, rates)
}

func rateGaps(c echo.Context) error {
	gaps := arbitrage.GetRateGaps(false)
	return c.JSON(http.StatusOK, gaps)
}

func compare(c echo.Context) error {
	gaps := arbitrage.GetRateGaps(true)
	return c.JSON(http.StatusOK, gaps)
}
