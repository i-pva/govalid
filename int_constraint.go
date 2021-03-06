package govalid

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type intConstraint struct {
	field    string
	req      bool
	isMinSet bool
	min      int
	isMaxSet bool
	max      int
	in       []int
}

func (ic *intConstraint) violation(val reflect.Value) string {
	i := val.Interface().(int)
	if !ic.req && i == 0 {
		return ""
	}
	if ic.req && i == 0 {
		return fmt.Sprintf("%s is required", ic.field)
	}
	if ic.isMaxSet && i > ic.max {
		return fmt.Sprintf("%s can not be greater than %d", ic.field, ic.max)
	}
	if ic.isMinSet && i < ic.min {
		return fmt.Sprintf("%s must be at least %d", ic.field, ic.min)
	}
	if len(ic.in) > 0 {
		for _, opt := range ic.in {
			if i == opt {
				return ""
			}
		}
	} else {
		return ""
	}
	iStrSlice := []string{}
	for _, a := range ic.in {
		iStrSlice = append(iStrSlice, strconv.Itoa(a))
	}
	return fmt.Sprintf("%s must be in [%s]", ic.field, strings.Join(iStrSlice, ", "))
}

func (ic *intConstraint) violations(val reflect.Value) []string {
	var vs []string
	i := val.Interface().(int)
	if !ic.req && i == 0 {
		return nil
	}
	if ic.req && i == 0 {
		vs = append(vs, fmt.Sprintf("%s is required", ic.field))
	}
	if ic.isMaxSet && i > ic.max {
		vs = append(vs, fmt.Sprintf("%s can not be greater than %d", ic.field, ic.max))
	}
	if ic.isMinSet && i < ic.min {
		vs = append(vs, fmt.Sprintf("%s must be at least %d", ic.field, ic.min))
	}
	if len(ic.in) > 0 {
		for _, opt := range ic.in {
			if i == opt {
				return vs
			}
		}
	} else {
		return vs
	}
	iStrSlice := []string{}
	for _, a := range ic.in {
		iStrSlice = append(iStrSlice, strconv.Itoa(a))
	}
	return append(vs, fmt.Sprintf("%s must be in [%s]", ic.field, strings.Join(iStrSlice, ", ")))
}
