package main

import (
	"reflect"
	"testing"
	"time"
)

// Test for function returnTime.
func Test_returnTime(t *testing.T) {
	tests := []struct {
		name    string
		t       string
		want    time.Time
		wantErr bool
	}{
		{"unix basic test", "1450137600", time.Unix(1450137600, 0).UTC(), false},
		{"natural language test", "December 15, 2015", time.Unix(1450137600, 0).UTC(), false},
		{"nonsense", "ooglyboo", time.Time{}, true},
	}
	for _, tt := range tests {
		got, err := returnTime(tt.t)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. returnTime() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. returnTime() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
