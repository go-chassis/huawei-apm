namespace java com.huawei.apm.thrift
namespace go apm
struct TSpan{ 
	1:i64 traceId
	3:string name
	4:i64 id
	5:optional i64 parentId
	6:list<TAnnotation> annotations 
	8:list<TBinaryAnnotation> binaryAnnotations 
	9:optional bool debug 
	10:optional i64 timestamp 
	11:optional i64 duration
}

struct TAnnotation{ 
	1:i64 timestamp 
  	2:optional string value 
    3:optional TEndpoint endpoint  
} 

struct TEndpoint{ 
	1:i32 ipv4 
	2:i16 port
	3:string serviceName 
}

struct TBinaryAnnotation{ 
 1:string key 
 2:string value 
 3:i32 type 
 4:optional TEndpoint endpoint 
}