package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

// Test for function returnTime.
func Test_parseTime(t *testing.T) {
	tests := []struct {
		name    string
		t       string
		want    time.Time
		wantErr bool
	}{
		{"unix basic test", "1450137600", time.Unix(1450137600, 0).UTC(), false},
		{"natural language test", "December 15, 2015", time.Unix(1450137600, 0).UTC(), false},
		{"nonsense", "ooglyboo", time.Time{}, true},
		{"Check that UNIX time is correct", "January 29, 2031", time.Unix(1927411200, 0).UTC(), false},
		{"natural language test with short month", "Dec 15, 2015", time.Unix(1450137600, 0).UTC(), false},
		{"short garbage", "1 1 1", time.Time{}, true},
	}
	for _, tt := range tests {
		got, err := parseTime(tt.t)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. returnTime() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. returnTime() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_timestamp(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"1450137600", "{\"unix\":1450137600,\"natural\":\"December 15, 2015\"}\n"},
		{"December 15, 2015", "{\"unix\":1450137600,\"natural\":\"December 15, 2015\"}\n"},
		{"garbage Not Date", "{\"unix\":null,\"natural\":null}\n"},
	}
	for _, c := range cases {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/"+c.in,
			nil,
		)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		timestamp(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200; got %d", rec.Code)
		}
		if rec.Body.String() != c.out {
			t.Errorf("unexpected body in response: %q\nExpected: %q", rec.Body.String(), c.out)
		}
	}
}
