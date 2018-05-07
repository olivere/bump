package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestBump(t *testing.T) {
	tests := []struct {
		OldVersion string
		Kind       string
		NewVersion string
		Error      error
	}{
		{
			OldVersion: "1.0.0",
			Kind:       "patch",
			NewVersion: "1.0.1",
		},
		{
			OldVersion: "v1.0.0",
			Kind:       "patch",
			NewVersion: "v1.0.1",
		},
		{
			OldVersion: "1.0.0",
			Kind:       "minor",
			NewVersion: "1.1.0",
		},
		{
			OldVersion: "1.0.0",
			Kind:       "major",
			NewVersion: "2.0.0",
		},
		{
			OldVersion: "1.0.0\n",
			Kind:       "patch",
			NewVersion: "",
			Error:      errors.New("Invalid Semantic Version"),
		},
		{
			OldVersion: "1.0.0",
			Kind:       "kaboom",
			NewVersion: "",
			Error:      errors.New("Invalid kind"),
		},
	}
	for _, tt := range tests {
		testcase := fmt.Sprintf("bump(%q, %q)", tt.OldVersion, tt.Kind)
		have, err := bump(tt.OldVersion, tt.Kind)
		if err != nil {
			if tt.Error == nil {
				t.Fatalf("%s: want no error, have %v", testcase, err)
			}
		} else {
			if tt.Error != nil {
				t.Fatalf("%s: want error %v, have %v", testcase, err, nil)
			}
			if want := tt.NewVersion; want != have {
				t.Fatalf("%s: want %q, have %q", testcase, want, have)
			}
		}
	}
}
