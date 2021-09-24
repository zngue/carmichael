package main

import (
	"context"
	"github.com/zngue/go_helper/pkg/response"
	"log"

	"github.com/zngue/carmichael/app/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zngue/carmichael/app/router"
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/go_helper/pkg/sign_chan"
)

func main() {

	if conErr := pkg.NewConfig(); conErr != nil {
		log.Fatal(conErr)
		return
	}
	mysql, err := pkg.NewMysql()
	if err != nil {
		log.Fatal(err)
		return
	}
	if mysql != nil {
		//auto.Auto(mysql)
		mysql.AutoMigrate(new(model.ZngUser), new(model.ZngKm), new(model.ZngOrder))
	}
	port := viper.GetString("AppPort")
	run, errs := pkg.GinRun(port, func(engine *gin.Engine) {
		engine.NoRoute(func(c *gin.Context) {
			response.HttpFailWithCodeAndMessage(404,"路由不存在",c)
		})
		group := engine.Group("carmichael")
		router.Router(group)
	})
	if errs != nil {
		sign_chan.SignLog(errs)
	}
	go func() {

		err := run.ListenAndServe()
		if err != nil {
			sign_chan.SignLog(err)
		}
	}()
	sign_chan.SignChalNotify()
	sign_chan.ListClose(func(ctx context.Context) error {
		return run.Shutdown(ctx)
	})

}
