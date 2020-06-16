package sql

import (
	"errors"
	"fmt"
	"strings"
)

func ParseQuery(query []*Condition) (string,error) {
	if len(query) == 0 {
		return "", errors.New(fmt.Sprintf("parse query failed"))
	}

	conditionStrarr := make([]string, 0)
	for _, condition := range query {
		switch {
		case condition.opt == "=" && condition.value.t == 1:
			conditionStrarr = append(conditionStrarr, fmt.Sprintf("%s = %d", condition.field, int(condition.value.v_number)))
		case condition.opt == "=" && condition.value.t == 2:
			// TODO 存在sql注入
			conditionStrarr = append(conditionStrarr, fmt.Sprintf("%s = '%s'", condition.field, condition.value.v_string))

		}
	}

	return strings.Join(conditionStrarr, " and "), nil
}