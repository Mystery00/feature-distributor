package enum

type ValueType int8

var (
	BoolValueType   ValueType = 1
	StringValueType ValueType = 2
	FloatValueType  ValueType = 3
	IntValueType    ValueType = 4
	JsonValueType   ValueType = 5
)

func (ValueType) Values() []ValueType {
	return []ValueType{
		BoolValueType,
		StringValueType,
		FloatValueType,
		IntValueType,
		JsonValueType,
	}
}

func ValueTypeEnum(v int8) ValueType {
	switch v {
	case 1:
		return BoolValueType
	case 2:
		return StringValueType
	case 3:
		return FloatValueType
	case 4:
		return IntValueType
	case 5:
		return JsonValueType
	default:
		panic("unknown value type")
	}
}
