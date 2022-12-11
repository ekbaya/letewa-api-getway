package main

import (
	"log"

	"github.com/ekbaya/letewa-api-getway/pkg/auth"
	"github.com/ekbaya/letewa-api-getway/pkg/config"
	"github.com/ekbaya/letewa-api-getway/pkg/order"
	"github.com/ekbaya/letewa-api-getway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
