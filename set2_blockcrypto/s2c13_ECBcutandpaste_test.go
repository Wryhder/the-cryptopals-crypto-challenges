package set2_blockcrypto

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	toParse := "foo=bar&baz=qux&zap=zazzle"
	actual := parse(toParse)
    expected := map[string]string{
		"foo": "bar",
		"baz": "qux",
		"zap": "zazzle",
	}

	eq := reflect.DeepEqual(actual, expected)
	if !eq {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}

func TestEncodeUserProfile(t *testing.T) {
	// Test 1
	userInput_1 := "foo@bar.com"
	actual_1 := encodeUserProfile(userInput_1)
	expected_1 := "email=foo@bar.com&uid=1&role=user"

	if actual_1 != expected_1 {
        t.Errorf("actual %q, expected %q", actual_1, expected_1)
    }

	// Test 2
	userInput_2 := "foo@bar.com&role=admin"
	actual_2 := encodeUserProfile(userInput_2)
	expected_2 := "email=foo@bar.com&uid=2&role=user"

	if actual_2 != expected_2 {
        t.Errorf("actual %q, expected %q", actual_2, expected_2)
    }
}

func TestDecryptAndParseUserProfile(t *testing.T) {
	userInput := "foo@bar.com"
	encodedProfile := encodeUserProfile(userInput)
	encryptedProfile := encryptUserProfile(encodedProfile)
	decryptedProfile, _ := decryptAndParseUserProfile(encryptedProfile)

	actual := decryptedProfile
    expected := map[string]string{
		"email": "foo@bar.com",
		"uid": "3", // 3rd function call to encodeUserProfile()
		"role": "user",
	}

	eq := reflect.DeepEqual(actual, expected)
	if !eq {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}

func TestECBCutAndPaste(t *testing.T) {
	isAdmin := ECBCutAndPaste_IsAdmin()

	if !isAdmin {
        t.Errorf("Role is 'user' but expected to be 'admin'")
    }
}