package utils

import "testing"

func TestReverseString(t *testing.T) {
	testCases := []TestCase[string, string]{
		{Case: "oaic", Expected: "ciao"},
		{Case: "massa", Expected: "assam"},
		{Case: "ácido", Expected: "odicá"},
		{Case: "o adoçante", Expected: "etnaçoda o"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Case, func(t *testing.T) {
			reversed := Reverse(testCase.Case)
			if reversed != testCase.Expected {
				t.Errorf("expected %s got %s", testCase.Expected, reversed)
			}
		})
	}
}
