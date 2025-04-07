package sbi

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"
)

func (s *Server) getBootstrappingRoutes() []Route {
	return []Route{
		{
			"Index",
			http.MethodGet,
			"/",
			func(c *gin.Context) {
				c.JSON(http.StatusOK, "free5gc")
			},
		},
		{
			"BootstrappingInfoRequest",
			http.MethodGet,
			"/bootstrapping",
			s.HTTPBootstrappingInfoRequest,
		},
	}
}

func (s *Server) HTTPBootstrappingInfoRequest(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
