package util

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"

	nrf_context "github.com/nycu-ucr/nrf/internal/context"
	"github.com/nycu-ucr/nrf/internal/logger"
	"github.com/nycu-ucr/openapi/models"
)

type (
	NFContextGetter          func() *nrf_context.NRFContext
	RouterAuthorizationCheck struct {
		serviceName models.ServiceName
	}
)

func NewRouterAuthorizationCheck(serviceName models.ServiceName) *RouterAuthorizationCheck {
	return &RouterAuthorizationCheck{
		serviceName: serviceName,
	}
}

func (rac *RouterAuthorizationCheck) Check(c *gin.Context, nrfContext nrf_context.NFContext) {
	token := c.Request.Header.Get("Authorization")
	err := nrfContext.AuthorizationCheck(token, rac.serviceName)
	if err != nil {
		logger.UtilLog.Debugf("RouterAuthorizationCheck::Check Unauthorized: %s", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	logger.UtilLog.Debugf("RouterAuthorizationCheck::Check Authorized")
}
