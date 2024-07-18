package utils

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "example_test_int_reverse",
			args: args{source: []int{1, 2, 3, 4, 5, 6, 7, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reverse(tt.args.source)
		})
		n := 8
		expect := []int{8, 7, 6, 5, 4, 3, 2, 1}
		for m := 0; m < n; m++ {
			if !reflect.DeepEqual(tt.args.source[m], expect[m]) {
				t.Error("Reverse() want true, but false")
			}
		}
	}

}

func TestFlatten(t *testing.T) {
	type args struct {
		source [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example_test_2_dimension_slice_flatten",
			args: args{source: [][]int{{1, 2, 3}, {4, 5, 6}}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flatten(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinct(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example_distinct_1",
			args: args{source: []int{1, 2, 1, 2, 3, 1, 5, 6, 2}},
			want: []int{1, 2, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distinct(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	type args struct {
		s1   []int
		rest [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_case_for_join",
			args: args{
				s1:   []int{1, 2, 3, 4, 5, 6},
				rest: [][]int{{7, 8, 9, 10}, {11, 12, 13, 14}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.s1, tt.args.rest...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMap(t *testing.T) {
	type Struct struct {
		Value []int
	}
	type args struct {
		source []Struct
		f      func(p Struct) []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example_test_flatten_map",
			args: args{source: []Struct{
				{Value: []int{1, 2, 3, 4, 5}},
				{Value: []int{6, 7, 8, 9, 10}},
			}, f: func(p Struct) []int {
				return p.Value
			}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlatMap(tt.args.source, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
