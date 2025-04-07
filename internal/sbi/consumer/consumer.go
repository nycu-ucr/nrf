package consumer

import (
	"github.com/nycu-ucr/nrf/pkg/app"
	"github.com/nycu-ucr/openapi/nrf/NFManagement"
)

type ConsumerNrf interface {
	app.App
}

type Consumer struct {
	ConsumerNrf

	*nnrfService
}

func NewConsumer(nrf ConsumerNrf) (*Consumer, error) {
	c := &Consumer{
		ConsumerNrf: nrf,
	}

	c.nnrfService = &nnrfService{
		consumer:        c,
		nfMngmntClients: make(map[string]*NFManagement.APIClient),
	}
	return c, nil
}
