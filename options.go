package huaweiapm

type Options struct {
	//tracing
	TracingBatchSize     int
	TracingBatchInterval string `validate:"regexp=^\d+(s|m)$"`

	//Inventory options
	Ports           []string
	Hostname        string
	IP              string
	Project         string
	App             string
	ServiceType     string `validate:"min=1"`
	ServiceName     string `validate:"min=1"`
	MonitoringGroup string `validate:"min=1"`
}
