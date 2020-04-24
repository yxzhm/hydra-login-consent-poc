package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yxzhm/hydra-login-consent-poc/handlers"
)

func main() {
	handlers.HydraAdmin = handlers.GetEnvStrValue("HYDRA_ADMIN", "http://127.0.0.1:4445")
	handlers.OemClientID = handlers.GetEnvStrValue("GITHUB_CLIENT_ID", "")
	router := gin.Default()
	n := 1
	a := &n
	fmt.Println(a)
	router.GET("/login", handlers.LoginHandler)
	router.GET("/consent", handlers.ConsentHandler)
	router.GET("/callback", handlers.CallbackHandler)
	router.GET("/dex/callback", handlers.AcceptLogin)

	router.Run(":9020")
}
