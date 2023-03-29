package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	c := make(chan Profile)
	go func() {
		c <- Profile{33, "Ray"}
		c <- Profile{34, "Sumire"}
		close(c)
	}()

	tests := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
		check         func(t *testing.T, got []string, want []string)
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Ray"},
			[]string{"Ray"},
			func(t *testing.T, got, want []string) {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Ray", "Tokyo"},
			[]string{"Ray", "Tokyo"},
			func(t *testing.T, got, want []string) {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			},
		},
		{
			"Struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Ray", 10},
			[]string{"Ray"},
			func(t *testing.T, got, want []string) {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			},
		},
		{
			"Nested fields",
			Person{
				"Ray",
				Profile{33, "Tokyo"},
			},
			[]string{"Ray", "Tokyo"},
			func(t *testing.T, got, want []string) {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			},
		},
		{
			"Pointer to things",
			&Person{
				"Ray",
				Profile{33, "Tokyo"},
			},
			[]string{"Ray", "Tokyo"},
			func(t *testing.T, got, want []string) {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			},
		},
		{
			"Slices",
			[]Profile{
				{33, "Tokyo"},
				{33, "Osaka"},
			},
			[]string{"Tokyo", "Osaka"},
			func(t *testing.T, got, want []string) {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			},
		},
		{
			"Slices",
			[2]Profile{
				{33, "Tokyo"},
				{33, "Osaka"},
			},
			[]string{"Tokyo", "Osaka"},
			func(t *testing.T, got, want []string) {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			},
		},
		{
			"Map",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
			func(t *testing.T, got, want []string) {
				assertContains(t, got, "Bar")
				assertContains(t, got, "Boz")
			},
		},
		{
			"with channels",
			c,
			[]string{"Ray", "Sumire"},
			func(t *testing.T, got, want []string) {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			},
		},
		{
			"with function",
			func() (Profile, Profile) {
				return Profile{33, "Ray"}, Profile{34, "Sumire"}
			},
			[]string{"Ray", "Sumire"},
			func(t *testing.T, got, want []string) {
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			test.check(t, got, test.ExpectedCalls)
		})
	}
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
