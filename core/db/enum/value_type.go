package enum

import "feature-distributor/common/value"

type ValueType int8

const (
	BoolValueType   ValueType = 1
	StringValueType ValueType = 2
	FloatValueType  ValueType = 3
	IntValueType    ValueType = 4
	JsonValueType   ValueType = 5
)

var (
	values = []ValueType{
		BoolValueType,
		StringValueType,
		FloatValueType,
		IntValueType,
		JsonValueType,
	}
	typeMap = map[ValueType]value.Type{
		BoolValueType:   value.BoolValueType,
		StringValueType: value.StringValueType,
		FloatValueType:  value.FloatValueType,
		IntValueType:    value.IntValueType,
		JsonValueType:   value.JsonValueType,
	}
)

func ParseValueType(s string) ValueType {
	vType := value.ParseType(s)
	if vType == nil {
		panic("unknown value type")
	}
	for k, v := range typeMap {
		if v == *vType {
			return k
		}
	}
	panic("unknown value type")
}

func (v ValueType) String() string {
	return typeMap[v].String()
}

func ValueTypeEnum(v int8) ValueType {
	for _, valueType := range values {
		if int8(valueType) == v {
			return valueType
		}
	}
	panic("unknown value type")
}
