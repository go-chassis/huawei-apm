package huaweiapm

import (
	"github.com/go-chassis/huawei-apm/pkg/fifo"
	"github.com/go-mesh/openlogging"
)

const (
	kpi       = "profiler.rpckpis.enabled"
	tracing   = "profiler.spans.zipkin.enabled"
	inventory = "profiler.discovery.enabled"
)

func watchConfigs(app, service string) {
	r, err := fifo.NewReader(app, service)
	if err != nil {
		openlogging.Fatal(err.Error())
	}
	openlogging.Debug("reading config")
	for {

		b, err := r.ReadBytes('\n')
		if err != nil {
			openlogging.Error(err.Error())
		}
		openlogging.Info("read config:" + string(b))
	}
}
