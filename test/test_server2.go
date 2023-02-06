package main

import (
	"fmt"
	"log"

	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"

	"github.com/free5gc/path_util"
	"github.com/nycu-ucr/http2_util"
	"github.com/nycu-ucr/logger_util"
	"github.com/nycu-ucr/nrf/logger"
	. "github.com/nycu-ucr/openapi/models"
)

var (
	NrfLogPath = path_util.Free5gcPath("github.com/nycu-ucr/nrf/management/sslkeylog.log")
	NrfPemPath = path_util.Free5gcPath("free5gc/support/TLS/nrf.pem")
	NrfKeyPath = path_util.Free5gcPath("free5gc/support/TLS/nrf.key")
)

func main() {
	router := logger_util.NewGinWithLogrus(logger.GinLog)

	router.POST("", func(c *gin.Context) {
		/*buf, err := c.GetRawData()
		if err != nil {
			t.Errorf(err.Error())
		}
		// Remove NL line feed, new line character
		//requestBody = string(buf[:len(b uf)-1])*/
		var ND NotificationData

		if err := c.ShouldBindJSON(&ND); err != nil {
			log.Panic(err.Error())
		}
		fmt.Println(ND)
		c.JSON(http.StatusNoContent, gin.H{})
	})

	srv, err := http2_util.NewServer(":30678", NrfLogPath, router)
	if err != nil {
		log.Panic(err.Error())
	}

	err2 := srv.ListenAndServeTLS(NrfPemPath, NrfKeyPath)
	if err2 != nil && err2 != http.ErrServerClosed {
		log.Panic(err2.Error())
	}
}
