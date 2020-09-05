package subscriptions

import (
	"io"
	"log"
	"os"
)

type Collector struct {
	Params *Params
	Out    io.Writer
}

type Params struct {
}

func NewCollector() *Collector {
	return &Collector{
		Params: &Params{},
		Out:    os.Stdout,
	}
}

func (c *Collector) validateParams() error {
	return nil
}

func (c *Collector) Subscriptions() error {
	if err := c.validateParams(); err != nil {
		return err
	}

	_, err := c.Out.Write([]byte("sup\n"))
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
