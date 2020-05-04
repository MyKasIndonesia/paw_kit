package ctx

import (
	"context"
	"reflect"
	"testing"
)

func TestSetRequestID(t *testing.T) {
	type args struct {
		ctx       context.Context
		RequestID string
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "1. Set empty string",
			args: args{
				ctx:       context.Background(),
				RequestID: "",
			},
			want: context.WithValue(context.Background(), contextKeyRequestID, ""),
		},
		{
			name: "2. Set non-empty string",
			args: args{
				ctx:       context.Background(),
				RequestID: "RequestID",
			},
			want: context.WithValue(context.Background(), contextKeyRequestID, "RequestID"),
		},
		{
			name: "3. Set empty string to non empty context",
			args: args{
				ctx:       context.WithValue(context.Background(), contextKeyRequestID, "RequestID-1234"),
				RequestID: "",
			},
			want: context.WithValue(context.Background(), contextKeyRequestID, ""),
		},
		{
			name: "4. Set non-empty string to non-empty context",
			args: args{
				ctx:       context.WithValue(context.Background(), contextKeyRequestID, "RequestID-1234"),
				RequestID: "RequestID",
			},
			want: context.WithValue(context.Background(), contextKeyRequestID, "RequestID"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetRequestID(tt.args.ctx, tt.args.RequestID)
			if !reflect.DeepEqual(got.Value(contextKeyRequestID), tt.want.Value(contextKeyRequestID)) {
				t.Errorf(
					"SetRequestID() = %v, want %v",
					got.Value(contextKeyRequestID),
					tt.want.Value(contextKeyRequestID),
				)
			}
		})
	}
}

func TestRequestID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1. Get RequestID from default context background",
			args: args{
				ctx: context.Background(),
			},
			want: "",
		},
		{
			name: "2. Get RequestID empty string",
			args: args{
				ctx: context.WithValue(context.Background(), contextKeyRequestID, ""),
			},
			want: "",
		},
		{
			name: "3. Get RequestID non-empty string",
			args: args{
				ctx: context.WithValue(context.Background(), contextKeyRequestID, "RequestID"),
			},
			want: "RequestID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequestID(tt.args.ctx); got != tt.want {
				t.Errorf("RequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}
