package api

import (
	"reflect"
	"testing"
)

func TestSuccess(t *testing.T) {
	tests := []struct {
		name string
		data interface{}
		want Response
	}{
		{
			name: "1. Data is a string",
			data: "a string",
			want: Response{
				Success: true,
				Data:    "a string",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Success(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Success() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFailed(t *testing.T) {
	tests := []struct {
		name string
		msg  string
		want Response
	}{
		{
			name: "1. error",
			msg:  "error",
			want: Response{
				Error: "error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Failed(tt.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Failed() = %v, want %v", got, tt.want)
			}
		})
	}
}
