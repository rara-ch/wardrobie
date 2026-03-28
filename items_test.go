package main

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestParseItem(t *testing.T) {
	type test struct {
		input     []string
		want      item
		expectErr bool
	}

	tests := []test{
		{
			input: []string{"hoodie", "-color", "red", "-brand", "Nike"},
			want: item{
				name:  "hoodie",
				color: sql.NullString{String: "red", Valid: true},
				brand: sql.NullString{String: "Nike", Valid: true},
			},
			expectErr: false,
		},
		{
			input:     []string{},
			expectErr: true,
		},
		{
			input: []string{
				"socks",
			},
			want: item{
				name: "socks",
			},
			expectErr: false,
		},
		{
			input:     []string{"-color", "red"},
			expectErr: true,
		},
		{
			input:     []string{"-brand"},
			expectErr: true,
		},
		{
			input: []string{"underwear", "-brand", "Nike", "-color", "blue", "-material", "cotton", "-category", "sports"},
			want: item{
				name:     "underwear",
				brand:    sql.NullString{String: "Nike", Valid: true},
				color:    sql.NullString{String: "blue", Valid: true},
				material: sql.NullString{String: "cotton", Valid: true},
				category: sql.NullString{String: "sports", Valid: true},
			},
			expectErr: false,
		},
	}

	for _, tc := range tests {
		got, err := parseItem(tc.input)
		if tc.expectErr && err == nil {
			t.Fatal("expected to get error but did not")
		} else if !tc.expectErr && err != nil {
			t.Fatalf("did not expect an error but got an error: %s", err)
		} else if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
