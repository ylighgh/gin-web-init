package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStrings_WriteOnCondition(t *testing.T) {
	type Field struct {
		Name string

		Fuzzy bool

		Domain string

		CicdUrl string

		AppType int

		Release string
	}
	var ff = &Field{
		Name:    "example_test_field_name",
		Fuzzy:   true,
		Domain:  "example_test_field_domain",
		CicdUrl: "",
		AppType: 1,
		Release: "example_test_field_release",
	}
	tests := []struct {
		name    string
		fields  *Strings
		args    *Field
		want    int
		wantErr bool
	}{
		{
			name:   "example_test_write_on_condition",
			fields: StringBuilder(),
			args:   ff,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Strings{
				build: tt.fields.build,
			}
			if tt.args.Fuzzy {
				s.WriteOnCondition(func() bool {
					return !empty(tt.args.Name)
				}, fmt.Sprintf("name LIKE %%%s%%", tt.args.Name)).
					WriteOnCondition(func() bool {
						return !s.Empty() && !empty(tt.args.Release)
					}, " OR ").
					WriteOnCondition(func() bool {
						return !empty(tt.args.Release)
					}, fmt.Sprintf("release LIKE %%%s%%", tt.args.Release)).
					WriteOnCondition(func() bool {
						return !s.Empty() && !empty(tt.args.CicdUrl)
					}, " OR ").
					WriteOnCondition(func() bool {
						return !empty(tt.args.CicdUrl)
					}, fmt.Sprintf("cicdurl LIKE %%%s%%", tt.args.CicdUrl)).
					WriteOnCondition(func() bool {
						return !s.Empty() && !empty(tt.args.Domain)
					}, " OR ").
					WriteOnCondition(func() bool {
						return !empty(tt.args.Domain)
					}, fmt.Sprintf("domain LIKE %%%s%%", tt.args.Domain))
			}
			if !reflect.DeepEqual(s.String(), "name LIKE %example_test_field_name% OR release LIKE %example_test_field_release% OR domain LIKE %example_test_field_domain%") {
				t.Errorf("Should be equal, but got = %s", s.String())
			}
		})
	}
}

func empty(s string) bool {
	return len(s) == 0
}
