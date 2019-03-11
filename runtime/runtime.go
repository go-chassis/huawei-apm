package runtime

import (
	"bufio"
	"github.com/go-mesh/openlogging"
	"os"
)

const (
	PAASPodID   = "PAAS_POD_ID"
	PAASCluster = "PAAS_CLUSTER_ID"
	PAASNS      = "PAAS_NAMESPACE"
	EnvProject  = "PAAS_PROJECT_ID"
)
const (
	ICAgentIDPath = "/tmp/apm/caf/caf-agent-id"
)

//PodID return huawei cloud CCE pod id
func PodID() (string, error) {
	pod := os.Getenv(PAASPodID)
	if pod == "" {
		var err error
		openlogging.Info("Not running in CCE, try to read ICAgentID")
		f, err := os.Open(ICAgentIDPath)
		if err != nil {
			openlogging.Error("can not open ICAgentID file, err: " + err.Error())
			return "", err
		}
		defer f.Close()
		reader := bufio.NewReader(f)

		pod, err = reader.ReadString('\n')
		if err != nil {
			openlogging.Error("can not read ICAgentID, err: " + err.Error())
			return "", err
		}
	}
	return pod, nil
}

func Cluster() string {
	return os.Getenv(PAASCluster)
}

func Namespace() string {
	return os.Getenv(PAASNS)
}
