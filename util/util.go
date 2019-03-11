package util

import (
	"net"
	"bytes"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/go-chassis/huawei-apm/thrift/gen-go/apm"
)

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
func Serialize(spans []*apm.TDiscoveryInfo) *bytes.Buffer {
	t := thrift.NewTMemoryBuffer()
	p := thrift.NewTBinaryProtocolTransport(t)
	if err := p.WriteListBegin(thrift.STRUCT, len(spans)); err != nil {
		panic(err)
	}
	for _, s := range spans {
		if err := s.Write(p); err != nil {
			panic(err)
		}
	}
	if err := p.WriteListEnd(); err != nil {
		panic(err)
	}
	return t.Buffer
}