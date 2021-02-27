package compute_test

import (
	"github.com/glassmonkey/flutter_with_golang/golang/pkg/compute"
	"testing"
)

func TestAddDecimal(t *testing.T) {
	tests := []struct {
		name       string
		v1        string
		v2        string
		want string
		wantErr bool
	}{
		{
			name: "1+2=3",
			v1: "1",
			v2: "2",
			want: "3",
			wantErr: false,
		},
		{
			name: "big and small",
			v1: ".00001",
			v2: "100000",
			want: "100000.00001",
			wantErr: false,
		},
		{
			name: "first value is invalied",
			v1: "hoge",
			v2: "2",
			want: "",
			wantErr: true,
		},
		{
			name: "second value is invalied",
			v1: "1",
			v2: "hoge",
			want: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compute.Add(tt.v1, tt.v2)
			if got != tt.want {
				t.Fatalf("want = %s, got = %s", tt.want, got)
			}
			if ! tt.wantErr && err != nil{
				t.Fatalf("want no err, but has error %#v", err)
			}
		})
	}
}

func TestSubDecimal(t *testing.T) {
	tests := []struct {
		name       string
		v1        string
		v2        string
		want string
		wantErr bool
	}{
		{
			name: "10 - 0.1 = 9.9",
			v1: "10",
			v2: "0.1",
			want: "9.9",
			wantErr: false,
		},
		{
			name: "1 - 2 = 3",
			v1: "1",
			v2: "2",
			want: "-1",
			wantErr: false,
		},
		{
			name: "big and small",
			v1: ".00001",
			v2: "100000",
			want: "-99999.99999",
			wantErr: false,
		},
		{
			name: "first value is invalied",
			v1: "hoge",
			v2: "2",
			want: "",
			wantErr: true,
		},
		{
			name: "second value is invalied",
			v1: "1",
			v2: "hoge",
			want: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compute.Sub(tt.v1, tt.v2)
			if got != tt.want {
				t.Fatalf("want = %s, got = %s", tt.want, got)
			}
			if ! tt.wantErr && err != nil{
				t.Fatalf("want no err, but has error %#v", err)
			}
		})
	}
}