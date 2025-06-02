package spacebeacon_test

import (
	"testing"

	"github.com/reiver/space-base/lib/beacon"
)

func TestHasMagic(t *testing.T) {

	tests := []struct{
		Data []byte
		Expected bool
	}{
		{

		},



		{
			Data: nil,
			Expected: false,
		},
		{
			Data: []byte(nil),
			Expected: false,
		},
		{
			Data: []byte(""),
			Expected: false,
		},



		{
			Data: []byte("S"),
			Expected: false,
		},
		{
			Data: []byte("SP"),
			Expected: false,
		},
		{
			Data: []byte("SPA"),
			Expected: false,
		},
		{
			Data: []byte("SPAC"),
			Expected: false,
		},
		{
			Data: []byte("SPACE"),
			Expected: false,
		},
		{
			Data: []byte("SPACE/"),
			Expected: false,
		},
		{
			Data: []byte("SPACE/0"),
			Expected: false,
		},
		{
			Data: []byte("SPACE/0."),
			Expected: false,
		},
		{
			Data: []byte("SPACE/0.1"),
			Expected: false,
		},
		{
			Data: []byte("SPACE/0.1\n"),
			Expected: true,
		},
		{
			Data: []byte("SPACE/0.1\nD"),
			Expected: true,
		},
		{
			Data: []byte("SPACE/0.1\nDO"),
			Expected: true,
		},
		{
			Data: []byte("SPACE/0.1\nDOR"),
			Expected: true,
		},
		{
			Data: []byte("SPACE/0.1\nDORO"),
			Expected: true,
		},
		{
			Data: []byte("SPACE/0.1\nDOROO"),
			Expected: true,
		},
		{
			Data: []byte("SPACE/0.1\nDOROOD"),
			Expected: true,
		},
		{
			Data: []byte("SPACE/0.1\nDOROOD\n"),
			Expected: true,
		},



		{
			Data: []byte("A"),
			Expected: false,
		},
		{
			Data: []byte("AB"),
			Expected: false,
		},
		{
			Data: []byte("ABC"),
			Expected: false,
		},
		{
			Data: []byte("ABCD"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/0"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/0."),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/0.1"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/0.1\n"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/0.1\nD"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/0.1\nDO"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/0.1\nDOR"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/0.1\nDORO"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/0.1\nDOROO"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/0.1\nDOROOD"),
			Expected: false,
		},
		{
			Data: []byte("ABCDE/0.1\nDOROOD\n"),
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := spacebeacon.HasMagic(test.Data)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result is not what was expectef.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("DATA: %q", test.Data)
			continue
		}
	}
}
