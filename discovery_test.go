package huaweiapm_test

import (
	"fmt"

	"github.com/stretchr/testify/assert"

	"github.com/go-chassis/huawei-apm"
	"os"
	"testing"
)

func TestNewResourceID(t *testing.T) {
	s := huaweiapm.NewResourceID(nil, "123")
	t.Log(s)

	s = huaweiapm.NewRawResourceID(nil, "123")
	assert.Equal(t, fmt.Sprintf("1230%d", os.Getpid()), s)

	s = huaweiapm.NewRawResourceID([]string{"8080"}, "123")
	assert.Equal(t, fmt.Sprintf("1238080"), s)
}

func TestGetAppID(t *testing.T) {
	s := huaweiapm.NewAppID("region", "group")
	t.Log(s)
	s = huaweiapm.NewRawAppID("region", "group")
	assert.Equal(t, "region|default|group", s)
}

func TestNewTDiscoveryInfo(t *testing.T) {
	os.Setenv("PAAS_POD_ID", "1")
	opts := huaweiapm.Options{
		MonitoringGroup: "engine-app",
		ServiceName:     "cart",
		ServiceType:     "go-chassis"}
	err := huaweiapm.Start(opts)
	assert.Error(t, err)
	t.Log(err)

	opts = huaweiapm.Options{
		MonitoringGroup: "engine-app",
		ServiceName:     "cart",
		ServiceType:     "go-chassis",
	}
	d, err := huaweiapm.BuildTDiscoveryInfo(opts)
	assert.NoError(t, err)
	t.Log(d)
}
