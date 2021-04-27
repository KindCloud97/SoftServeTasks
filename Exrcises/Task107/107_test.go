package main

import (
	"testing"
)

type numtest struct {
	value  []int
	result []int
}

var tests = []numtest{
	{[]int{123, 12, 45, 67, 17895}, []int{4, 2, 3, 4, 8}},
}

func TestTask107(t *testing.T) {
	for _, elem := range tests {
		for id := range elem.value {
			v := Task107(elem.value[id])
			if v != elem.result[id] {
				t.Error(
					"For", elem.value[id],
					"expected", elem.result[id],
					"got", v,
				)
			}
		}

	}
}
