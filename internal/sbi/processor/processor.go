package processor

import (
	"github.com/nycu-ucr/nrf/internal/sbi/consumer"
	"github.com/nycu-ucr/nrf/pkg/app"
)

type ProcessorNrf interface {
	app.App
	Consumer() *consumer.Consumer
}

type Processor struct {
	ProcessorNrf
}

func NewProcessor(nrf ProcessorNrf) (*Processor, error) {
	p := &Processor{
		ProcessorNrf: nrf,
	}
	return p, nil
}
