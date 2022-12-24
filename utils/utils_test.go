package utils

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignum(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{
		{2, 1},
		{0, 0},
		{-3, -1},
		{math.MaxInt32, 1},
		{math.MinInt32, -1},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Signum(tt.arg))
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{
		{2, 2},
		{0, 0},
		{-3, 3},
		{math.MaxInt32, math.MaxInt32},
		{math.MinInt32, math.MaxInt32 + 1},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Abs(tt.arg))
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		arg1 int
		arg2 int
		want int
	}{
		{10, 3, 1},
		{2, 4, 2},
		{24, 30, 6},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, GCD(tt.arg1, tt.arg2))
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		arg1 int
		arg2 int
		want int
	}{
		{10, 3, 30},
		{2, 4, 4},
		{24, 30, 120},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, LCM(tt.arg1, tt.arg2))
	}
}

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		arg  []int
		want []int
	}{
		{
			[]int{1, 5, 2, -8, 0},
			[]int{0, -8, 2, 5, 1},
		},
		{
			[]int{1, 5, 2, -8},
			[]int{-8, 2, 5, 1},
		},
		{
			nil,
			nil,
		},
	}
	for _, tt := range tests {
		ReverseSlice(tt.arg)
		assert.Equal(t, tt.want, tt.arg)
	}
}
