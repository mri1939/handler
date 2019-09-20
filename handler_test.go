package handler_test

import (
	"testing"

	"github.com/mri1939/handler"
)

func TestGetURIParam(t *testing.T) {
	table := []struct {
		name   string
		prefix string
		uri    string
		result []string
	}{
		{"Single Param", "/user", "/user/1", []string{"1"}},
		{"Double Params", "/user", "/user/1/2", []string{"1", "2"}},
		{"More Params", "/something", "/something/1/2/3/4/5", []string{"1", "2", "3", "4", "5"}},
		{"No Param", "/user", "/user/", nil},
		{"Trailing slash", "/user", "/user/1/", []string{"1"}},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			res := handler.GetURIParam(test.prefix, test.uri)
			if test.result == nil && res != nil {
				t.Fatalf("Expect nil result, got %v", res)
			}
			expect := len(test.result)
			got := len(res)
			if len(res) != len(test.result) {
				t.Fatalf("Mismatch the length of parsed param(s) (expect: %d,got: %d)", expect, got)
			}
			for i, r := range res {
				if r != test.result[i] {
					t.Errorf("Failed expected : %s , got :%s", test.result[i], r)
				}
			}
		})
	}
}
