package main

import (
	"testing"
)

func Test_GetSum(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{
			name: "Test Case 1",
			data: "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\nasefe1\nasef11pogh",
			want: 164,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSum(tt.data)
			if got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}