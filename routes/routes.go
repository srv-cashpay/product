package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/middlewares/middlewares"
	"github.com/srv-cashpay/product/configs"
	h_product "github.com/srv-cashpay/product/handlers/product"
	r_product "github.com/srv-cashpay/product/repositories/product"
	s_product "github.com/srv-cashpay/product/services/product"
)

var (
	DB       = configs.InitDB()
	JWT      = middlewares.NewJWTService()
	productR = r_product.NewProductRepository(DB)
	productS = s_product.NewProductService(productR, JWT)
	productH = h_product.NewProductHandler(productS)
)

func New() *echo.Echo {
	e := echo.New()

	product := e.Group("/api/product", middlewares.AuthorizeJWT(JWT))
	{
		product.GET("/:id", productH.GetById)
		product.GET("/pagination", productH.Get)
	}

	return e
}
