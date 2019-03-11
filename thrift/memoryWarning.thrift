namespace java com.huawei.apm.thrift
namespace go apm
struct TMemoryUsage {
	1: i64 used;
    2: i64 max;
    3: i64 free;
    4: double percentage;
}

struct TMemoryWarning {
	1: TMemoryUsage memoryUsage,
	2: i64 created,
	3: string memoryMeasurement,
	4: TMemoryEventType memoryEventType
}

enum TMemoryEventType {
    RESET,
    WARNING,
    NO_CHANGE
}