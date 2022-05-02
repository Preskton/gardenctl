/*===========================================================================*\
*  MIT License Copyright (c) 2022 Preston Doster <preston.doster@genoq.com>   *
*                                                                             *
*                                888                            888    888    *
*                                888                            888    888    *
*                                888                            888    888    *
*   .d88b.   8888b.  888d888 .d88888  .d88b.  88888b.   .d8888b 888888 888    *
*  d88P"88b     "88b 888P"  d88" 888 d8P  Y8b 888 "88b d88P"    888    888    *
*  888  888 .d888888 888    888  888 88888888 888  888 888      888    888    *
*  Y88b 888 888  888 888    Y88b 888 Y8b.     888  888 Y88b.    Y88b.  888    *
*   "Y88888 "Y888888 888     "Y88888  "Y8888  888  888  "Y8888P  "Y888 888    *
*       888                                                                   *
*  Y8b d88P                                                                   *
*   "Y88P"                                                                    *
*                                                                             *
\*===========================================================================*/

package plant

import (
	"time"

	"github.com/preskton/gardenctl/pkg/api"
	"github.com/sirupsen/logrus"
)

// Compile check *Plant implements Runner interface
var _ api.Runner = &Plant{}

type Plant struct {
	// Fields
}

func NewPlant() *Plant {
	return &Plant{}
}

var (
	runtimePlant bool = true
)

func (n *Plant) Run() error {
	client := api.Client{}
	server := api.Server{}
	logrus.Infof("Client: %x", client)
	logrus.Infof("Server: %x", server)
	for runtimePlant {
		time.Sleep(1 * time.Second)
		logrus.Infof("Sleeping...\n")
	}
	return nil
}
