namespace java com.huawei.apm.thrift
namespace go apm
enum SpanType {
    INTERMEDIATE,
    FIRST_FOR_BACKEND,
    FIRST_FOR_CLIENT,
    FIRST_FOR_UNKNOWN,
    FIRST_FOR_ENDPOINT,
    ENDPOINT,
    EXTERNAL,
    FIRST_EXTERNAL,
    ENDPOINT_EXTERNAL
}

struct TKpiMessage {
10:	string sourceResouceId,
20:	string destResouceId,
30:     string transactionType,
40:     string appId,
50:	binary selfErrorLatency,
60:     i32 throughput,
70:	binary selfLatency,
71: binary selfActiveLatency,
80:	binary totalLatency,
81: binary totalActiveLatency,
90: SpanType spanType,
110:    list<i32> totalLatencyList,
120:    list<bool> totalErrorIndicatorList,
121:    binary totalErrorLatency,
130: string namespaceName
140: optional string srcTierName ="",
150: optional string destTierName =""
}


struct TInternalKpi {
10: TKpiMessage kpi
20: string partitionKey;
30: i64 utcTimeMin;
40: i64 insertTs;
}


struct TKpiOverall {
1:	string destResouceId,
2:      string appId,
3:	binary selfErrorLatency,
4:      i32 throughput,
5:	binary selfLatency,
6:	binary totalLatency,
7:  binary totalErrorLatency
}


struct TKpiMessageByTransactionType {
1:	string destResouceId,
2:      string transactionType,
3:      string appId,
4:       binary selfErrorLatency,
5:      i32 throughput,
6:	binary selfLatency,
7:	binary totalLatency,
8:  binary totalErrorLatency
}

struct TSQLKPIMessage {
10: string sqlId,
20: string txtype,
30: string resouceId,
40: string destResourceId,
50: string appId,
60: binary latency,
70: binary errorLatency,
80: set<string> traceIds4SlowSQL,
90: string traceId4SQLError
}
