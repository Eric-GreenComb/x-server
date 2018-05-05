package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/sync/errgroup"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Eric-GreenComb/x-server/config"
	"github.com/Eric-GreenComb/x-server/ether"
	"github.com/Eric-GreenComb/x-server/handler"
	"github.com/Eric-GreenComb/x-server/persist"
)

var (
	g errgroup.Group
)

func main() {
	if config.ServerConfig.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	persist.InitDatabase()

	ether.LoadEthClient()

	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.MaxMultipartMemory = 64 << 20 // 64 MiB

	router.Use(Cors())

	/* api base */
	r0 := router.Group("/")
	{
		r0.GET("", handler.Index)
		r0.GET("health", handler.Health)

		r0.POST("login", Login)
	}

	// api
	r1 := router.Group("/api/v1")
	{
		r1.POST("/users/create", handler.CreateUser)
		r1.GET("/users/:userid", handler.UserInfo)

		r1.POST("/account/create/:userid/:password", handler.CreateAccount)
		// r1.POST("/account/load/:keystore", handler.LoadKeystore)
		r1.GET("/account/info/:address", handler.GetKeystore)

		r1.POST("/token/deploy", handler.DeployToken)
		r1.GET("/token/balance", handler.BalanceOfToken)
		r1.POST("/token/transfer", handler.TransferToken)

		r1.POST("/badger/set/:key/:value", handler.SetBadgerKey)
		r1.POST("/badger/setwithttl/:key/:value", handler.SetBadgerKeyTTL)
		r1.GET("/badger/get/:key", handler.GetBadgerKey)
	}

	// auth api
	r2 := router.Group("/api/auth/v1")
	r2.Use(JWTAuth())
	{
		r2.GET("/hello", handler.GetHello)
		r2.POST("/hello", handler.PostHello)
		r2.GET("/refresh_token", RefreshToken)
	}

	for _, _port := range config.ServerConfig.Port {
		server := &http.Server{
			Addr:         ":" + _port,
			Handler:      router,
			ReadTimeout:  300 * time.Second,
			WriteTimeout: 300 * time.Second,
		}

		g.Go(func() error {
			return server.ListenAndServe()
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
