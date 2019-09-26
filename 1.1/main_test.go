package main

import (
	"testing"
	"reflect"
	"math/rand"
)

func TestProduct(t *testing.T){
	tcs := []struct {
		input, want []int
	}{
		{
			input: []int{3,2,1},
			want: []int{2,3,6},
		},
		{
			input: []int{1,2,3,4,5},
			want: []int{120,60,40,30,24},
		},
		{
			input: []int{5,3,2,4},
			want: []int{24,40,60,30},

		},
	}
	for _, tc := range tcs {
		got := product(tc.input)
		if  !reflect.DeepEqual(got, tc.want){
			t.Errorf("product(%v); want %v, got %v", tc.input, tc.want, got)
		}
	}
}


func BenchmarkProduct(b *testing.B){
	var input []int
	for i := 1; i < b.N; i++{
		input = append(input, rand.Intn(99)+1)
	}
	b.ResetTimer()
	product(input)
}


func TestProductDiv(t *testing.T){
	tcs := []struct {
		input, want []int
	}{
		{
			input: []int{3,2,1},
			want: []int{2,3,6},
		},
		{
			input: []int{1,2,3,4,5},
			want: []int{120,60,40,30,24},
		},
		{
			input: []int{5,3,2,4},
			want: []int{24,40,60,30},

		},
	}
	for _, tc := range tcs {
		got := productDiv(tc.input)
		if  !reflect.DeepEqual(got, tc.want){
			t.Errorf("product(%v); want %v, got %v", tc.input, tc.want, got)
		}
	}
}

func BenchmarkProductDiv(b *testing.B){
	var input []int
	for i := 1; i < b.N; i++{
		input = append(input, rand.Intn(99)+1)
	}
	b.ResetTimer()
	productDiv(input)
}



func TestProductNoDiv(t *testing.T){
	tcs := []struct {
		input, want []int
	}{
		{
			input: []int{3,2,1},
			want: []int{2,3,6},
		},
		{
			input: []int{1,2,3,4,5},
			want: []int{120,60,40,30,24},
		},
		{
			input: []int{5,3,2,4},
			want: []int{24,40,60,30},

		},
	}
	for _, tc := range tcs {
		got := productNoDiv(tc.input)
		if  !reflect.DeepEqual(got, tc.want){
			t.Errorf("product(%v); want %v, got %v", tc.input, tc.want, got)
		}
	}
}

func BenchmarkProductNoDiv(b *testing.B){
	var input []int
	for i := 1; i < b.N; i++{
		input = append(input, rand.Intn(99)+1)
	}
	b.ResetTimer()
	productNoDiv(input)
}
