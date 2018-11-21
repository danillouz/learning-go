package walk

import (
	"reflect"
	"testing"
)

type WalkSpy struct {
	Calls []string
}

func (s *WalkSpy) fn(field string) {
	s.Calls = append(s.Calls, field)
}

func TestWalk(t *testing.T) {
	tests := []struct {
		Name      string
		Input     interface{}
		WantCalls []string
	}{
		{
			Name:      "struct with one string field",
			Input:     struct{ Name string }{"daniel"},
			WantCalls: []string{"daniel"},
		},
		{
			Name:      "struct with two string fields",
			Input:     struct{ Name, Email string }{"daniel", "daniel.illouz@me.com"},
			WantCalls: []string{"daniel", "daniel.illouz@me.com"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"daniel", 31},
			WantCalls: []string{"daniel"},
		},
		{
			Name: "nested structs",
			Input: User{
				"daniel",
				Info{"Groningen", 31},
			},
			WantCalls: []string{"daniel", "Groningen"},
		},
		{
			Name: "struct with pointer",
			Input: &User{
				"daniel",
				Info{"Groningen", 31},
			},
			WantCalls: []string{"daniel", "Groningen"},
		},
		{
			Name: "slice of structs",
			Input: []Info{
				{"Groningen", 31},
				{"Groningen", 31},
			},
			WantCalls: []string{"Groningen", "Groningen"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			spy := &WalkSpy{}

			Walk(tt.Input, spy.fn)

			got := spy.Calls
			want := tt.WantCalls

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

type User struct {
	Name string
	Info Info
}

type Info struct {
	City string
	Age  int
}
