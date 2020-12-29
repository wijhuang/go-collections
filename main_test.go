package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"pos", []int{1, 2, 3, 4, 5}, 15},
		{"neg", []int{-1, -2, -3, -4, -5}, -15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.nums...); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
