package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yxzhm/hydra-login-consent-poc/handlers"
)

func main() {
	handlers.HydraAdmin = handlers.GetEnvStrValue("HYDRA_ADMIN", "http://127.0.0.1:4445")
	router := gin.Default()

	router.GET("/login", handlers.LoginHandler)
	router.GET("/consent", handlers.ConsentHandler)
	router.GET("/callback", handlers.CallbackHandler)

	router.Run(":9020")
}
