package main

import (
"testing"
)


type numtest struct {
	value []int
	a [][]int
	b [][]int
}

var tests = []numtest{
	{ []int{8,29,34}, [][]int{{2,2},{5,2,2,5},{5,3,3,5}},[][]int{{2,2},{5,2},{5,3}}}

func TestTask243(t *testing.T) {
	for _, elem := range tests {
		for id := range elem.value {
			v := Task243(elem.value[id][id])
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

