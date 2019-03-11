namespace java com.huawei.apm.thrift
namespace go apm
struct TEndpointKpiMessage {
10:     string transactionType,
15:     string destAppId,
20:     string destResourceId,
30:     string destTierName,
40:     string originalDestination,
100:    i32 throughput,
110:    list<i32> totalLatencyList,
120:    list<bool> totalErrorIndicatorList,
130:    list<i32> totalActiveLatencyList,
200:    map<string,string> segments

}

enum TFailureEventType {
    CRASH,
    ANR
}

struct TEndpointFailureMessage {
10: TFailureEventType failureEventType,
20: map<string,string> environmentProperties
}

struct TEndpointMessage {
10:     string endpointType,
20:     string srcAppId,
100:    map<string,string> globalSegments,
200:    list<TEndpointKpiMessage> messages,
300:    list<TEndpointStatsMessage> stats,
400:    map<string,list<binary>> spans
500:    optional TEndpointFailureMessage failureMessage;
}

struct TEndpointStatsMessage {
100: map<string,double> numericStats,
200: map<string,string> segments
}


struct TEndpointWrapper {
10:    map<string,string> metadata,
100:   binary data
}