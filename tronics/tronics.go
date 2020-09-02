package tronics

import (
	"fmt"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/labstack/echo"
)

var e = echo.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
}

//serverMessage: example middleware
func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside custom middleware")
		c.Request().URL.Path = "/kirin"

		fmt.Printf("%+v", c.Request())
		return next(c)
	}
}

//Start application
func Start() {
	e.Pre(serverMessage) //run before action
	e.GET("/products", getProducts)
	e.GET("/products1/:vendor", func(c echo.Context) error {
		return c.JSON(http.StatusOK, c.QueryParam("oldlerThan"))
	})
	e.GET("/products/:id", getProduct)
	e.POST("/products", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port: %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
