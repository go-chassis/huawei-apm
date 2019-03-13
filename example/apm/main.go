package main

import (
	"github.com/go-chassis/huawei-apm"
	"github.com/stretchr/testify/assert"
	"github.com/openzipkin-contrib/zipkin-go-opentracing/thrift/gen-go/zipkincore"
)

func main(){
	err := huaweiapm.Start(huaweiapm.Options{})

	span := &zipkincore.Span{}
	huaweiapm.WriteSpan(span)
}
