package huaweiapm

import (
	"crypto/md5"
	"fmt"
	"github.com/go-chassis/huawei-apm/runtime"
	"github.com/go-chassis/huawei-apm/thrift/gen-go/apm"
	"github.com/go-mesh/openlogging"
	"io"
	"net"
	"os"
	"time"
)

const (
	AgentVersion = "2.1.0"
)

var instance *apm.TDiscoveryInfo

func startDiscovery(disco *apm.TDiscoveryInfo) {
	if err := client.ReportDiscoveryInfo(disco); err != nil {
		openlogging.Error("can not report inventory: " + err.Error())
	}
	openlogging.Debug("report inventory success")
	t, _ := time.ParseDuration(DefaultDiscoveryInterval)
	ticker := time.Tick(t)

	for {
		select {
		case <-ticker:
			if err := client.ReportDiscoveryInfo(disco); err != nil {
				openlogging.Error("can not report inventory: " + err.Error())
			}
			openlogging.Debug("report inventory success")
		case stop := <-StopInventory:
			if stop {
				openlogging.Info("inventory stopped")
				break
			}
		}
	}
}
func GetDiscoveryInfo() *apm.TDiscoveryInfo {
	return instance
}

//GetLocalIP 获得本机IP
func GetLocalIP() string {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addresses {
		// Parse IP
		var ip net.IP
		if ip, _, err = net.ParseCIDR(address.String()); err != nil {
			return ""
		}
		// Check if valid global unicast IPv4 address
		if ip != nil && (ip.To4() != nil) && ip.IsGlobalUnicast() {
			return ip.String()
		}
	}
	return ""
}

//BuildTDiscoveryInfo create a APM instance info
//if you report this info to cloud, APM can discovery your process
//should report it every 5 mins
func BuildTDiscoveryInfo(opts Options) (*apm.TDiscoveryInfo, error) {
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
		ip = GetLocalIP()
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
	openlogging.Debug(fmt.Sprintf("build inventory %s", instance))
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
