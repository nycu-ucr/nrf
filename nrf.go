/*
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/free5gc/version"
	"github.com/nycu-ucr/nrf/logger"
	nrf_service "github.com/nycu-ucr/nrf/service"
)

var NRF = &nrf_service.NRF{}

var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
}

func main() {
	app := cli.NewApp()
	app.Name = "nrf"
	fmt.Print(app.Name, "\n")
	appLog.Infoln("NRF version: ", version.GetVersion())
	app.Usage = "-free5gccfg common configuration file -nrfcfg nrf configuration file"
	app.Action = action
	app.Flags = NRF.GetCliCmd()

	if err := app.Run(os.Args); err != nil {
		appLog.Errorf("NRF Run Error: %v", err)
	}
}

func action(c *cli.Context) error {
	if err := NRF.Initialize(c); err != nil {
		logger.CfgLog.Errorf("%+v", err)
		return fmt.Errorf("Failed to initialize !!")
	}

	NRF.Start()

	return nil
}
