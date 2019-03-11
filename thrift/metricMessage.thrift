namespace java com.huawei.apm.thrift
namespace go apm
struct TMetricValues {
	1: list<double> values
}

struct TMetricMessage {
	1: string name,
	2: TMetricValues values,
	3: map<string,string> labels
	4: string resouceId,
	5: string appId  
}

