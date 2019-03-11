namespace java com.huawei.apm.thrift
namespace go apm
struct TDiscoveryInfo {
	1:string hostname
	2:string ip
	3:string agentId
	4:string appName
	5:string clusterKey
	6:string serviceType
	7:string displayName
	8:string instanceName
	9:string containerId
	10:i32 pid
	12:string projectId
	13:string podId
	14:string collectorId
	15:string appId
	20: map<string,string> props
	30:list<i32> ports
	40:list<string> ips
	50:string tier
	60:string namespaceName
	61:i64 created
	62:i64 updated
	63:i64 deleted
	70:optional list<TProfilingInfo> profilingInfo
}

struct TProfilingInfo {
  1:string pointcut
  2:i32 requests
  3:i64 elapsedMs
  4:i32 status
  5:string exclusion
}

struct TInventoryInfo {
	1:string kind
	2:map<string, string> metadata
	3:list<TItem> items
	70:optional list<TProfilingInfo> profilingInfo
}

struct TItem {
	1:TMetadata metadata
	2:map<string, string> labels
	3:TSpec spec
}

struct TMetadata {
	1:string name
	2:string type
	3:string id
	4:string schemaVersion
	5:i64 created
	6:i64 updated
	7:list<TReference> references
}

struct TReference {
	1:string kind
	2:string name
	3:list<map<string, string>> references
}

struct TSpec {
	1:map<string, string> properties
	2:list<string> ips
	3:list<i32> ports
}