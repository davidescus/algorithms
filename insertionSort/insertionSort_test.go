package main

import "testing"

type opt struct {
	val      []int
	expect   []int
	sortType string
}

func TestEmptySlice(t *testing.T) {
	value := sort([]int{}, asc)
	if cap(value) > 0 {
		t.Error("Must return an empty slice, but return this: ", value)
	}
}

func TestMultipleValueAscending(t *testing.T) {
	opt := []opt{
		{
			val:      []int{1, 5, 2},
			expect:   []int{1, 2, 5},
			sortType: "asc",
		},
		{
			val:      []int{5, 3},
			expect:   []int{3, 5},
			sortType: "asc",
		},
		{
			val:      []int{4},
			expect:   []int{4},
			sortType: "asc",
		},
		{
			val:      []int{},
			expect:   []int{},
			sortType: "asc",
		},
		{
			val:      []int{1, 5, 2},
			expect:   []int{5, 2, 1},
			sortType: "desc",
		},
		{
			val:      []int{5, 3},
			expect:   []int{5, 3},
			sortType: "desc",
		},
		{
			val:      []int{4},
			expect:   []int{4},
			sortType: "desc",
		},
		{
			val:      []int{},
			expect:   []int{},
			sortType: "desc",
		},
	}

	for _, cases := range opt {

		var sorted []int
		if cases.sortType == "asc" {
			sorted = sort(cases.val, asc)
		}

		if cases.sortType == "desc" {
			sorted = sort(cases.val, desc)
		}

		if len(sorted) != len(cases.expect) {
			t.Error("Expecting: ", cases.expect, " got: ", sorted)
		}

		for i, v := range sorted {
			if cases.expect[i] != v {
				t.Error("Expecting: ", cases.expect, " got: ", sorted)
			}
		}
	}
}
