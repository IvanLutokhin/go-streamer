package log

type FieldType uint8

const (
	FieldUnknownType FieldType = iota
	FieldBoolType
	FieldIntType
	FieldFloat64Type
	FieldStringType
	FieldErrorType
)

type Field struct {
	Type  FieldType
	Key   string
	Value interface{}
}

func FieldBool(key string, value bool) Field {
	return Field{Type: FieldBoolType, Key: key, Value: value}
}

func FieldInt(key string, value int) Field {
	return Field{Type: FieldIntType, Key: key, Value: value}
}

func FieldFloat64(key string, value float64) Field {
	return Field{Type: FieldFloat64Type, Key: key, Value: value}
}

func FieldString(key string, value string) Field {
	return Field{Type: FieldStringType, Key: key, Value: value}
}

func FieldError(key string, value error) Field {
	return Field{Type: FieldErrorType, Key: key, Value: value}
}

func FieldAny(key string, value interface{}) Field {
	switch v := value.(type) {
	case bool:
		return FieldBool(key, v)
	case int:
		return FieldInt(key, v)
	case float64:
		return FieldFloat64(key, v)
	case string:
		return FieldString(key, v)
	case error:
		return FieldError(key, v)
	}

	return Field{Type: FieldUnknownType, Key: key, Value: value}
}

func Fields(context map[string]interface{}) (fields []Field) {
	for key, value := range context {
		fields = append(fields, FieldAny(key, value))
	}

	return
}
