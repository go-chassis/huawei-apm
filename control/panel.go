package control

import (
	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/huawei-apm/pkg/fifo"
)



func WatchConfigs(app,service string){
	r, err := fifo.NewReader(app,service)
}
func GetKPIEnabled() bool {
	return archaius.GetBool("profiler.rpckpis.enabled", true)

}

func GetTracingEnabled() bool {
	return archaius.GetBool("profiler.spans.zipkin.enabled", true)
}

func GetDiscoveryEnabled() bool {
	return archaius.GetBool("profiler.discovery.enabled", true)
}
