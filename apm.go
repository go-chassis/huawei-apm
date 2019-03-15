package huaweiapm

import (
	"github.com/go-chassis/huawei-apm/api"
	"github.com/go-mesh/openlogging"
	"gopkg.in/validator.v2"
)

const (
	DefaultDiscoveryInterval    = "290s"
	DefaultTracingBatchInterval = "1m"
	DefaultTracingBatchSize     = 1000
)

var client api.APM

var tr *TracingReporter

//Switchers
var StopTracing = make(chan bool)
var StopInventory = make(chan bool)
var StopKPI = make(chan bool)

var StartTracing = make(chan bool)
var StartInventory = make(chan bool)
var StartKPI = make(chan bool)

func Start(opts Options) error {
	if err := validator.Validate(opts); err != nil {
		return err
	}
	opt := opts

	disco, err := BuildTDiscoveryInfo(opts)
	if err != nil {
		openlogging.Error("can not build discovery info: " + err.Error())
		return err
	}
	client, err = api.GetAPMClient(opts.App, opts.ServiceName)
	if err != nil {
		return err
	}
	openlogging.Debug("APM client init success")
	go startDiscovery(disco)

	go watchConfigs(opt.App, opt.ServiceName)

	//watch signal
	go func() {
		select {
		case s := <-StartTracing:
			if s && !tr.tracingRunning {
				openlogging.Info("restart tracing")
				go tr.StartReportSpans()
			}

		}
	}()
	return nil
}
