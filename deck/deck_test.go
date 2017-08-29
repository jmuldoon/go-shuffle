package deck

import "testing"

const targetTestVersion = 0

func TestVersionValidation(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestGenerate(t *testing.T) {
	for _, test := range testPokerDeckCases {
		observed := &Deck{}
		observed.Generate(test.Tested.Value.(string))
		t.Logf("Running test for `%s`\n", test.Description)
		if observed != test.Expected.Value {
			t.Fatalf("%s(%v):\n"+
				"Brief (%s)\n"+
				"Observed: %v\n"+
				"Expected: %v\n",
				"TestGenerate", test.Tested.Value,
				test.Description, observed, test.Expected.Value)
		}
	}
}
