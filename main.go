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
		// 用户
		r1.POST("/users/create", handler.CreateUser)
		r1.GET("/users/:userid", handler.UserInfo)
		r1.POST("/users/updatepasswd/:userid/:old/:new", handler.UpdateUserPasswd)
		r1.GET("/user/list/:search/:page", handler.ListUser)
		r1.GET("/user/count", handler.CountUser)

		// 账户
		r1.POST("/account/create/:userid/:name/:password", handler.CreateAccount)
		r1.GET("/account/list/:userid", handler.ListAccount)
		r1.GET("/account/info/:address", handler.GetKeystore)
		r1.POST("/account/updatepwd/:addr/:password/:newpassword", handler.UpdateAccountPwd)
		r1.POST("/account/recover/:addr/:newpassword", handler.RecoverAccountPwd)

		// 用户-账户-token视图
		r1.POST("/view/create/:userid/:address/:tokenaddress", handler.CreateUserAddressTokens)
		r1.POST("/view/delete/:userid/:address/:tokenaddress", handler.DeleteUserAddressTokens)
		r1.GET("/view/info/:userid/:address/:tokenaddress", handler.GetUserAddressTokens)
		r1.GET("/view/list/:userid/:address", handler.ListUserAddressTokens)
		r1.GET("/view/balance/:userid/:address", handler.ListUserAddressTokenBanlance)

		// ether
		r1.POST("/token/deploy", handler.DeployToken)
		r1.POST("/token/selfdeploy", handler.SelfDeployToken)
		r1.POST("/token/balance/:addr", handler.BalanceOfToken)
		r1.POST("/token/transfer", handler.TransferToken)
		// db - token
		r1.POST("/token/create", handler.CreateToken)
		r1.GET("/token/info/:address", handler.TokenInfo)
		r1.POST("/token/weight/:address/:weight", handler.UpdateTokenWeight)
		r1.GET("/token/list/:search/:page", handler.ListToken)
		r1.GET("/token/count", handler.CountToken)

		// db - transfer
		r1.POST("/token/transfer/create", handler.CreateTokenTransfer)
		r1.GET("/token/transfer/list/:tokenaddress/:address/:page", handler.ListTokenTransfer)
		r1.GET("/token/transfer/all/:tokenaddress/:page", handler.AllTokenTransfer)
		r1.GET("/token/transfer/count/:tokenaddress", handler.CountTokenTransfer)
		r1.GET("/token/transfer/allcount", handler.CountAllTokenTransfer)
		r1.GET("/token/transfer/allsum", handler.SumAllTokenTransfer)

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

	// admin api
	r3 := router.Group("/api/admin/v1")
	{
		r3.POST("/users/create", handler.CreateAdminUser)
		r3.GET("/users/:userid", handler.AdminUserInfo)
		r3.POST("/users/updatepasswd/:userid/:old/:new", handler.UpdateAdminUserPasswd)
		r3.POST("login", AdminLogin)
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
