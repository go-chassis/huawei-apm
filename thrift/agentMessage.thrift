namespace java com.huawei.apm.thrift
namespace go apm
struct TAgentMessage {
	1: string agentContext,
	2: string tenantName,
	20: map<string,map<i64,list<binary>>> messages
}


