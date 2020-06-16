package sql

type Value struct {
	// 类型：1数值，2字符串，3数值数组，4字符串数组
	t int

	// 数值
	v_number float64
	// 字符串
	v_string string
	// 数值数组
	v_number_arr []float64
	// 字符串数组
	v_string_arr []string
}

func NewValue(t int, v_number float64, v_string string, v_number_arr []float64, v_string_arr []string) *Value {
	return &Value{t: t, v_number: v_number, v_string: v_string, v_number_arr: v_number_arr, v_string_arr: v_string_arr}
}

type Condition struct {
	// 条件字段
	field string
	// 操作
	opt string
	// 条件值
	value *Value
}

func NewCondition(field string, opt string, value *Value) *Condition {
	return &Condition{field: field, opt: opt, value: value}
}