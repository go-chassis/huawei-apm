package api_test

import (
	"github.com/go-chassis/huawei-apm"
	"github.com/go-chassis/huawei-apm/api"
	"github.com/go-chassis/huawei-apm/runtime"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetAPMClient(t *testing.T) {
	api, err := api.GetAPMClient("app", "service")
	assert.NoError(t, err)
	t.Log(api)
}
func TestDefaultAPM_ReportDiscoveryInfo(t *testing.T) {
	os.Setenv(runtime.PAASPodID, "1")
	api, _ := api.GetAPMClient("app", "service")
	disco, err := huaweiapm.BuildTDiscoveryInfo(huaweiapm.Options{
		ServiceName:     "A",
		ServiceType:     "go-chassis",
		MonitoringGroup: "app",
	})
	t.Log(disco)
	assert.NoError(t, err)
	err = api.ReportDiscoveryInfo(disco)
	assert.NoError(t, err)
}
