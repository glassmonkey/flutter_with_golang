package api_test

import (
	"errors"
	"testing"

	"github.com/glassmonkey/flutter_with_golang/golang/pkg/api"
)

func TestApi_Call(t *testing.T) {
	tests := []struct {
		name     string
		wantBody string
		wantErr  string
	}{
		{
			name:     "success",
			wantBody: "hello, world by golang",
			wantErr:  "fuga",
		},
	}

	for _, tt := range tests {
		got, err := api.GoCall()
		if got != tt.wantBody {
			t.Fatalf("want = %s, got = %s", tt.wantBody, got)
		}
		if tt.wantErr == "" {
			if err != nil {
				t.Fatalf("want no error, but has error: %#v", err)
			}
			return
		}
		wantErr := errors.New(tt.wantErr)
		if errors.Is(err, wantErr) {
			t.Fatalf("want err: %#v, but has error: %#v", tt.wantErr, err)
		}
	}
}

func TestApi_Formula(t *testing.T) {
	tests := []struct {
		name         string
		inputFormula string
		wantResult   float64
		wantErr      string
	}{
		{
			name:         "empty",
			inputFormula: "",
			wantResult:   0,
			wantErr:      "Unexpected end of expression",
		},
		{
			name:         "simple",
			inputFormula: "1 + 1",
			wantResult:   2,
			wantErr:      "",
		},
		{
			name:         "variable with invalid parameter",
			inputFormula: "a * b",
			wantResult:   0,
			wantErr:      "any",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.Evaluate(tt.inputFormula)
			if got != tt.wantResult {
				t.Errorf("want = %#v, got = %#v", tt.wantResult, got)
			}
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("want no error, but has error: %#v", err)
				}
				return
			}
			wantErr := errors.New(tt.wantErr)
			if errors.Is(err, wantErr) {
				t.Fatalf("want err: %#v, but has error: %#v", tt.wantErr, err)
			}
		})
	}
}
