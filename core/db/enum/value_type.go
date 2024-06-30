package enum

import "feature-distributor/common/value"

type ValueType int8

const (
	BoolValueType ValueType = 1 + iota
	StringValueType
	FloatValueType
	IntValueType
	JsonValueType
)

var (
	valueTypes = []ValueType{
		BoolValueType,
		StringValueType,
		FloatValueType,
		IntValueType,
		JsonValueType,
	}
	valueTypeMap = map[ValueType]value.Type{
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
	for k, v := range valueTypeMap {
		if v == *vType {
			return k
		}
	}
	panic("unknown value type")
}

func (v ValueType) String() string {
	return valueTypeMap[v].String()
}

func ValueTypeEnum(v int8) ValueType {
	for _, t := range valueTypes {
		if int8(t) == v {
			return t
		}
	}
	panic("unknown value type")
}
