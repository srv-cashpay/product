package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/middlewares/middlewares"
	"github.com/srv-cashpay/product/configs"
	h_product "github.com/srv-cashpay/product/handlers/product"
	r_product "github.com/srv-cashpay/product/repositories/product"
	s_product "github.com/srv-cashpay/product/services/product"

	h_topup "github.com/srv-cashpay/product/handlers/topup/mobilelegend"
	r_topup "github.com/srv-cashpay/product/repositories/topup/mobilelegend"
	s_topup "github.com/srv-cashpay/product/services/topup/mobilelegend"
)

var (
	DB       = configs.InitDB()
	JWT      = middlewares.NewJWTService()
	productR = r_product.NewProductRepository(DB)
	productS = s_product.NewProductService(productR, JWT)
	productH = h_product.NewProductHandler(productS)

	topupR = r_topup.NewTopUpRepository(DB)
	topupS = s_topup.NewTopUpService(topupR, JWT)
	topupH = h_topup.NewMobileLegendHandler(topupS)
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/api/product/menu", productH.Menu)

	product := e.Group("/api/product", middlewares.AuthorizeJWT(JWT))
	{
		product.GET("/:id", productH.GetById)
		product.GET("/pagination", productH.Get)
	}

	topup := e.Group("/api/product")
	{
		topup.POST("/topup/mobilelegend", topupH.TopUp)
	}

	webmenu := e.Group("api/product/web", middlewares.AuthorizeJWT(JWT))
	{
		webmenu.GET("/menu", productH.GetUrl)
	}
	return e
}
