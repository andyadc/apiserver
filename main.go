package main

import (
	"apiserver/config"
	"apiserver/model"
	"apiserver/router"
	"errors"
	"github.com/gin-gonic/gin"
	logger "github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {

	// log
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// db
	model.DB.Init()
	defer model.DB.Close()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	var middlewares []gin.HandlerFunc

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			logger.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		logger.Info("The router has been deployed successfully.")
	}()

	go func() {
		for {
			r := rand.Intn(9999999)
			logger.Info(strconv.Itoa(r))
			time.Sleep(time.Minute)
		}
	}()

	logger.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("port"))
	logger.Info(http.ListenAndServe(viper.GetString("port"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("serverurl") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		logger.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
