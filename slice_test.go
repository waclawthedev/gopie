package gopie_test

import (
	"reflect"
	"testing"

	"github.com/waclawthedev/gopie"
)

func TestMap(t *testing.T) {
	result := gopie.Map([]int{1, 2, 3, 4}, func(k int) int {
		return k - 1
	})
	if result[0] != 0 || result[1] != 1 || result[2] != 2 || result[3] != 3 {
		t.Fail()
	}

	result = gopie.Map([]int{}, func(k int) int {
		return k - 1
	})

	shouldEqual(t, len(result), 0)
}

func TestFilter(t *testing.T) {
	result := gopie.Filter([]int{1, 2, 3, 4}, func(k int) bool {
		return k <= 2
	})

	if result[0] != 1 || result[1] != 2 || len(result) != 2 {
		t.Fail()
	}

	result = gopie.Filter([]int{}, func(k int) bool {
		return k <= 2
	})

	shouldEqual(t, len(result), 0)
}

func TestReduce(t *testing.T) {
	result := gopie.Reduce([]int{1, 2, 3, 4}, func(a int, b int) int {
		return a + b
	})

	shouldEqual(t, result, 10)

	result = gopie.Reduce([]int{}, func(a int, b int) int {
		return a + b
	})

	shouldEqual(t, result, 0)

	result = gopie.Reduce([]int{1}, func(a int, b int) int {
		return a + b
	})

	shouldEqual(t, result, 1)
}

func TestFlatten(t *testing.T) {
	result := gopie.Flatten([][]int{{1, 2}, {3, 4}})

	if result[0] != 1 || result[1] != 2 || result[2] != 3 || result[3] != 4 {
		t.Fail()
	}

	result = gopie.Flatten([][]int{})
	if result != nil {
		t.Fail()
	}
}

//nolint:funlen
func Test_SplitByN(t *testing.T) {
	type args struct {
		s []int
		n int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Basic case",
			args: args{
				s: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				n: 3,
			},
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}},
		},
		{
			name: "n is bigger than slice length",
			args: args{
				s: []int{1, 2, 3},
				n: 5,
			},
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "n is same as slice length",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				n: 5,
			},
			want: [][]int{{1, 2, 3, 4, 5}},
		},
		{
			name: "empty slice",
			args: args{
				s: []int{},
				n: 3,
			},
			want: nil,
		},
		{
			name: "n is zero",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				n: 0,
			},
			want: nil,
		},
		{
			name: "n is negative",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				n: -1,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gopie.SplitByN(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByN() = %v, want %v", got, tt.want)
			}
		})
	}
}
