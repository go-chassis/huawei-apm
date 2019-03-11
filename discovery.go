package huaweiapm

import (
	"crypto/md5"
	"fmt"
	"github.com/go-chassis/go-chassis/pkg/util/iputil"
	"github.com/go-chassis/huawei-apm/runtime"
	"github.com/go-chassis/huawei-apm/thrift/gen-go/apm"
	"github.com/go-mesh/openlogging"
	"gopkg.in/validator.v2"
	"io"
	"os"
	"time"
)

const (
	AgentVersion = "2.1.0"
)

var instance *apm.TDiscoveryInfo

func GetDiscoveryInfo() *apm.TDiscoveryInfo {
	return instance
}

//BuildTDiscoveryInfo create a APM instance info
//if you report this info to cloud, APM can discovery your process
//should report it every 5 mins
func BuildTDiscoveryInfo(opts Options) (*apm.TDiscoveryInfo, error) {
	if err := validator.Validate(opts); err != nil {
		return nil, err
	}
	hostname := opts.Hostname
	var err error
	if hostname == "" {
		hostname, err = os.Hostname()
		if err != nil {
			return nil, err
		}
	}

	ip := opts.IP
	if ip == "" {
		ip = iputil.GetLocalIP()
	}
	app := opts.App
	if app == "" {
		app = opts.MonitoringGroup
	}
	if opts.Project == "" {
		opts.Project = "default"
		openlogging.Warn("project is empty")
	}
	if opts.MonitoringGroup == "" {
		opts.MonitoringGroup = "default"
	}
	pod, err := runtime.PodID()
	if err != nil {
		return nil, err
	}
	instance = &apm.TDiscoveryInfo{
		Hostname:    hostname,
		IP:          ip,
		AgentId:     NewResourceID(opts.Ports, pod),
		Pid:         int32(os.Getpid()),
		ProjectId:   opts.Project,
		AppName:     opts.MonitoringGroup,
		ClusterKey:  runtime.Cluster(),
		ServiceType: opts.ServiceType,
		DisplayName: opts.ServiceName,
		PodId:       pod,
		Props: map[string]string{
			"config.status": "true",
			"agent.version": AgentVersion,
		},
		NamespaceName: runtime.Namespace(),
		Created:       time.Now().Unix(),
		Updated:       time.Now().Unix(),
		Deleted:       0,
	}
	instance.CollectorId = instance.AgentId
	instance.AppId = NewAppID(instance.ProjectId, instance.GetAppName())
	instance.Tier = instance.DisplayName

	return instance, nil
}

//NewResourceID generate unique md5 ID for a process
func NewResourceID(ports []string, podID string) string {
	h := md5.New()
	io.WriteString(h, NewRawResourceID(ports, podID))
	return fmt.Sprintf("%x", h.Sum(nil))

}
func NewAppID(project, appName string) string {
	h := md5.New()
	io.WriteString(h, NewRawAppID(project, appName))
	return fmt.Sprintf("%x", h.Sum(nil))
}
func NewRawAppID(project, appName string) string {
	return project + "|" + "default" + "|" + appName
}
func NewRawResourceID(ports []string, podID string) string {
	if len(ports) == 0 {
		return fmt.Sprintf("%s0%d", podID, os.Getpid())
	}
	return podID + ports[0]
}
