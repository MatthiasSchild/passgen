package passgen_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/MatthiasSchild/passgen"
)

const (
	charactersLowerLetters = "abcdefghijklmnopqrstuvwxyz"
	charactersUpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charactersDigits       = "0123456789"
	charactersSpecialChars = "!\"#$%&’()*+,-./:;<=>?@[\\]^_`{|}~"
)

func countCharsFromSet(input string, set string) int {
	found := 0

	for i := 0; i < len(input); i++ {
		if strings.Contains(set, input[i:i+1]) {
			found++
		}
	}

	return found
}

func ExampleGenerate() {
	password, err := passgen.Generate(24, passgen.Options{
		MinimumLowerLetters:      4,
		MinimumUpperLetters:      4,
		MinimumDigits:            4,
		MinimumSpecialCharacters: 2,
	})
	if err != nil {
		fmt.Printf("Error while generating password: %v", err)
		return
	}

	fmt.Printf("Generated password: %s", password)
}

func TestGenerate_length(t *testing.T) {
	password, err := passgen.Generate(16, passgen.Options{})
	if err != nil {
		t.Errorf("Error while generating password: %v", err)
		return
	}

	t.Logf("Generated password: %s", password)
	if len(password) != 16 {
		t.Errorf("Password should have 16 characters, but has %d", len(password))
		return
	}
}

func TestGenerate_lowerCharacters(t *testing.T) {
	for loop := 0; loop < 100; loop++ {
		password, err := passgen.Generate(16, passgen.Options{
			MinimumLowerLetters: 8,
		})
		if err != nil {
			t.Errorf("Error while generating password: %v", err)
			return
		}

		t.Logf("Generated password: %s", password)
		found := countCharsFromSet(password, charactersLowerLetters)
		if found < 8 {
			t.Errorf("Password should have at least 8 lower chars, but has %d", found)
			return
		}
	}
}

func TestGenerate_lowerCharactersDisabled(t *testing.T) {
	for loop := 0; loop < 100; loop++ {
		password, err := passgen.Generate(16, passgen.Options{
			DisableLowerLetters: true,
		})
		if err != nil {
			t.Errorf("Error while generating password: %v", err)
			return
		}

		t.Logf("Generated password: %s", password)
		found := countCharsFromSet(password, charactersLowerLetters)
		if found > 0 {
			t.Errorf("Password should not have any lower chars, but has %d", found)
			return
		}
	}
}

func TestGenerate_upperCharacters(t *testing.T) {
	for loop := 0; loop < 100; loop++ {
		password, err := passgen.Generate(16, passgen.Options{
			MinimumUpperLetters: 8,
		})
		if err != nil {
			t.Errorf("Error while generating password: %v", err)
			return
		}

		t.Logf("Generated password: %s", password)
		found := countCharsFromSet(password, charactersUpperLetters)
		if found < 8 {
			t.Errorf("Password should have at least 8 upper chars, but has %d", found)
			return
		}
	}
}

func TestGenerate_upperCharactersDisabled(t *testing.T) {
	for loop := 0; loop < 100; loop++ {
		password, err := passgen.Generate(16, passgen.Options{
			DisableUpperLetters: true,
		})
		if err != nil {
			t.Errorf("Error while generating password: %v", err)
			return
		}

		t.Logf("Generated password: %s", password)
		found := countCharsFromSet(password, charactersUpperLetters)
		if found > 0 {
			t.Errorf("Password should not have any upper chars, but has %d", found)
			return
		}
	}
}

func TestGenerate_digits(t *testing.T) {
	for loop := 0; loop < 100; loop++ {
		password, err := passgen.Generate(16, passgen.Options{
			MinimumDigits: 8,
		})
		if err != nil {
			t.Errorf("Error while generating password: %v", err)
			return
		}

		t.Logf("Generated password: %s", password)
		found := countCharsFromSet(password, charactersDigits)
		if found < 8 {
			t.Errorf("Password should have at least 8 digits, but has %d", found)
			return
		}
	}
}

func TestGenerate_digitsDisabled(t *testing.T) {
	for loop := 0; loop < 100; loop++ {
		password, err := passgen.Generate(16, passgen.Options{
			DisableDigits: true,
		})
		if err != nil {
			t.Errorf("Error while generating password: %v", err)
			return
		}

		t.Logf("Generated password: %s", password)
		found := countCharsFromSet(password, charactersDigits)
		if found > 0 {
			t.Errorf("Password should not have any digits, but has %d", found)
			return
		}
	}
}

func TestGenerate_specialCharacters(t *testing.T) {
	for loop := 0; loop < 100; loop++ {
		password, err := passgen.Generate(16, passgen.Options{
			MinimumSpecialCharacters: 8,
		})
		if err != nil {
			t.Errorf("Error while generating password: %v", err)
			return
		}

		t.Logf("Generated password: %s", password)
		found := countCharsFromSet(password, charactersSpecialChars)
		if found < 8 {
			t.Errorf("Password should have at least 8 special chars, but has %d", found)
			return
		}
	}
}

func TestGenerate_specialCharactersDisabled(t *testing.T) {
	for loop := 0; loop < 100; loop++ {
		password, err := passgen.Generate(16, passgen.Options{
			DisableSpecialChars: true,
		})
		if err != nil {
			t.Errorf("Error while generating password: %v", err)
			return
		}

		t.Logf("Generated password: %s", password)
		found := countCharsFromSet(password, charactersSpecialChars)
		if found > 0 {
			t.Errorf("Password should not have any special chars, but has %d", found)
			return
		}
	}
}

func TestGenerate_customSpecialCharacters(t *testing.T) {
	expected := "!!!!!!!!!!!!!!!!"

	for loop := 0; loop < 100; loop++ {
		password, err := passgen.Generate(16, passgen.Options{
			MinimumSpecialCharacters: 16,
			SpecialCharacters:        "!",
		})
		if err != nil {
			t.Errorf("Error while generating password: %v", err)
			return
		}

		t.Logf("Generated password: %s", password)
		if password != expected {
			t.Errorf("Password should be %s, but is %s", expected, password)
			return
		}
	}
}

func TestGenerate_excludeSpecialCharacters(t *testing.T) {
	expected := strings.Repeat(charactersSpecialChars[0:1], 16)

	for loop := 0; loop < 100; loop++ {
		password, err := passgen.Generate(16, passgen.Options{
			MinimumSpecialCharacters: 16,
			ExcludeCharacters:        charactersSpecialChars[1:],
		})
		if err != nil {
			t.Errorf("Error while generating password: %v", err)
			return
		}

		t.Logf("Generated password: %s", password)
		if password != expected {
			t.Errorf("Password should be %s, but is %s", expected, password)
			return
		}
	}
}

func TestGenerate_errors(t *testing.T) {
	_, err := passgen.Generate(-1, passgen.Options{})
	if !errors.Is(err, passgen.ErrorPasswordLength) {
		t.Errorf("Password should throw ErrorPasswordLength, but thrown instead: %v", err)
		return
	}

	_, err = passgen.Generate(5, passgen.Options{
		MinimumLowerLetters: 3,
		MinimumUpperLetters: 3,
	})
	if !errors.Is(err, passgen.ErrorMinimumsTooHigh) {
		t.Errorf("Password should throw ErrorMinimumsTooHigh, but thrown instead: %v", err)
		return
	}

	_, err = passgen.Generate(16, passgen.Options{
		DisableLowerLetters: true,
		DisableUpperLetters: true,
		DisableDigits:       true,
		DisableSpecialChars: true,
	})
	if !errors.Is(err, passgen.ErrorEverythingDisabled) {
		t.Errorf("Password should throw ErrorEverythingDisabled, but thrown instead: %v", err)
		return
	}

	_, err = passgen.Generate(16, passgen.Options{
		MinimumLowerLetters: -1,
	})
	if !errors.Is(err, passgen.ErrorMinimumLowerLetters) {
		t.Errorf("Password should throw ErrorMinimumLowerLetters, but thrown instead: %v", err)
		return
	}

	_, err = passgen.Generate(16, passgen.Options{
		MinimumUpperLetters: -1,
	})
	if !errors.Is(err, passgen.ErrorMinimumUpperLetters) {
		t.Errorf("Password should throw ErrorMinimumUpperLetters, but thrown instead: %v", err)
		return
	}

	_, err = passgen.Generate(16, passgen.Options{
		MinimumDigits: -1,
	})
	if !errors.Is(err, passgen.ErrorMinimumDigits) {
		t.Errorf("Password should throw ErrorMinimumDigits, but thrown instead: %v", err)
		return
	}

	_, err = passgen.Generate(16, passgen.Options{
		MinimumSpecialCharacters: -1,
	})
	if !errors.Is(err, passgen.ErrorMinimumSpecialCharacters) {
		t.Errorf("Password should throw ErrorMinimumSpecialCharacters, but thrown instead: %v", err)
		return
	}
}
