package events

import (
	"reflect"
	"testing"
)

func TestSub(t *testing.T) {
	tests := []struct {
		Sub
		exname   string
		exparams map[string][]string
	}{
		{
			AllSub{},
			"all",
			map[string][]string{
				"characters": []string{"all"},
				"worlds":     []string{"all"},
			},
		},
		{
			CharSub{
				Event: "test",
				Chars: []string{"not", "a", "real", "event"},
			},
			"test",
			map[string][]string{
				"characters": []string{"not", "a", "real", "event"},
			},
		},
		{
			WorldSub{
				Event:  "another test",
				Worlds: []string{"also", "not", "a", "real", "event"},
			},
			"another test",
			map[string][]string{
				"worlds": []string{"also", "not", "a", "real", "event"},
			},
		},
	}

	for i, test := range tests {
		name := test.name()
		if name != test.exname {
			t.Errorf("%v: Names are not equal. Ex: %q Got: %q", i, test.exname, name)
		}

		pmap := test.params()
		if !reflect.DeepEqual(pmap, test.exparams) {
			t.Errorf("%v: Params are not equal.", i)
			t.Errorf("\tEx: %v", test.exparams)
			t.Errorf("\tGot: %v", pmap)
		}
	}
}
