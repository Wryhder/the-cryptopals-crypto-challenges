package set1_basics

import "testing"

func TestDetectSingleCharXOR(t *testing.T) {
	actual := DetectSingleCharXOR("../data/s1c4_60charstrings.txt")
    expected := "Now that the party is jumping\n"

    if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}