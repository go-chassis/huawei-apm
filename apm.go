package huaweiapm

import (
	"github.com/go-chassis/huawei-apm/api"
	"github.com/go-chassis/huawei-apm/thrift/gen-go/apm"
	"github.com/go-mesh/openlogging"
	"time"
)

const (
	DefaultDiscoveryInterval = "290s"
)

var client api.APM

func Start(opts Options) error {
	disco, err := BuildTDiscoveryInfo(opts)
	if err != nil {
		openlogging.Error("can not build discovery info: " + err.Error())
		return err
	}
	client, err = api.GetAPMClient(opts.App, opts.ServiceName)
	if err != nil {
		return err
	}
	startDiscovery(disco)
	return nil
}

func startDiscovery(disco *apm.TDiscoveryInfo) {
	t, _ := time.ParseDuration(DefaultDiscoveryInterval)
	ticker := time.Tick(t)
	for range ticker {
		client.ReportDiscoveryInfo(disco)
	}

}
