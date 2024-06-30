package enum

import (
	"feature-distributor/common/operation"
)

type OperationType int8

const (
	StringInType OperationType = 1 + iota
	StringOutType
	StringStartType
	StringNotStartType
	StringEndType
	StringNotEndType
	StringContainType
	StringNotContainType
	StringRegexType
	StringNotRegexType

	NumberEqType
	NumberNotEqType
	NumberGtType
	NumberGteType
	NumberLtType
	NumberLteType
	NumberBetweenType

	TimeBeforeType
	TimeAfterType
)

var (
	operationValues = []OperationType{
		StringInType,
		StringOutType,
		StringStartType,
		StringNotStartType,
		StringEndType,
		StringNotEndType,
		StringContainType,
		StringNotContainType,
		StringRegexType,
		StringNotRegexType,
		NumberEqType,
		NumberNotEqType,
		NumberGtType,
		NumberGteType,
		NumberLtType,
		NumberLteType,
		NumberBetweenType,
		TimeBeforeType,
		TimeAfterType,
	}
	operationTypeMap = map[OperationType]operation.Type{
		StringInType:         operation.StringInType,
		StringOutType:        operation.StringOutType,
		StringStartType:      operation.StringStartType,
		StringNotStartType:   operation.StringNotStartType,
		StringEndType:        operation.StringEndType,
		StringNotEndType:     operation.StringNotEndType,
		StringContainType:    operation.StringContainType,
		StringNotContainType: operation.StringNotContainType,
		StringRegexType:      operation.StringRegexType,
		StringNotRegexType:   operation.StringNotRegexType,
		NumberEqType:         operation.NumberEqType,
		NumberNotEqType:      operation.NumberNotEqType,
		NumberGtType:         operation.NumberGtType,
		NumberGteType:        operation.NumberGteType,
		NumberLtType:         operation.NumberLtType,
		NumberLteType:        operation.NumberLteType,
		NumberBetweenType:    operation.NumberBetweenType,
		TimeBeforeType:       operation.TimeBeforeType,
		TimeAfterType:        operation.TimeAfterType,
	}
)

func ParseOperationType(s string) OperationType {
	vType := operation.ParseType(s)
	if vType == nil {
		panic("unknown operation type")
	}
	for k, v := range operationTypeMap {
		if v == *vType {
			return k
		}
	}
	panic("unknown operation type")
}

func (v OperationType) String() string {
	return operationTypeMap[v].String()
}

func OperationTypeEnum(v int8) OperationType {
	for _, t := range operationValues {
		if int8(t) == v {
			return t
		}
	}
	panic("unknown operation type")
}
