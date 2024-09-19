package passgen

import (
	"errors"
	"math/rand"
	"strings"
)

const (
	charactersLowerLetters = "abcdefghijklmnopqrstuvwxyz"
	charactersUpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charactersDigits       = "0123456789"
	charactersSpecialChars = "!\"#$%&’()*+,-./:;<=>?@[\\]^_`{|}~"
)

var (
	ErrorPasswordLength           = errors.New("the requested password length needs to be larger than 0")
	ErrorMinimumsTooHigh          = errors.New("the minimum values are in sum larger than the requested length")
	ErrorEverythingDisabled       = errors.New("you should not disable all character types")
	ErrorMinimumLowerLetters      = errors.New("the minimum lower letters should be a positive number")
	ErrorMinimumUpperLetters      = errors.New("the minimum upper letters should be a positive number")
	ErrorMinimumDigits            = errors.New("the minimum digits should be a positive number")
	ErrorMinimumSpecialCharacters = errors.New("the minimum special characters should be a positive number")
)

type Options struct {
	DisableLowerLetters      bool
	MinimumLowerLetters      int
	DisableUpperLetters      bool
	MinimumUpperLetters      int
	DisableDigits            bool
	MinimumDigits            int
	DisableSpecialChars      bool
	MinimumSpecialCharacters int
	SpecialCharacters        string
	ExcludeCharacters        string
}

func buildSpecialCharSet(options Options) string {
	set := charactersSpecialChars
	if len(options.SpecialCharacters) != 0 {
		set = options.SpecialCharacters
	}
	if len(options.ExcludeCharacters) != 0 {
		for i := 0; i < len(options.ExcludeCharacters); i++ {
			remove := options.ExcludeCharacters[i : i+1]
			set = strings.ReplaceAll(set, remove, "")
		}
	}
	return set
}

func Generate(passwordLength int, options Options) (string, error) {
	var buffer string
	var result string

	if passwordLength <= 0 {
		return "", ErrorPasswordLength
	}

	// build an own special char set depending on the options
	specialChars := buildSpecialCharSet(options)

	// Fill the minimum required lower letters
	if !options.DisableLowerLetters {
		if options.MinimumLowerLetters < 0 {
			return "", ErrorMinimumLowerLetters
		}
		for i := 0; i < options.MinimumLowerLetters; i++ {
			pos := rand.Intn(len(charactersLowerLetters))
			buffer += charactersLowerLetters[pos : pos+1]
		}
	}

	// Fill the minimum required upper letters
	if !options.DisableUpperLetters {
		if options.MinimumUpperLetters < 0 {
			return "", ErrorMinimumUpperLetters
		}
		for i := 0; i < options.MinimumUpperLetters; i++ {
			pos := rand.Intn(len(charactersUpperLetters))
			buffer += charactersUpperLetters[pos : pos+1]
		}
	}

	// Fill the minimum required digits
	if !options.DisableDigits {
		if options.MinimumDigits < 0 {
			return "", ErrorMinimumDigits
		}
		for i := 0; i < options.MinimumDigits; i++ {
			pos := rand.Intn(len(charactersDigits))
			buffer += charactersDigits[pos : pos+1]
		}
	}

	// Fill the minimum required special chars
	if !options.DisableSpecialChars {
		if options.MinimumSpecialCharacters < 0 {
			return "", ErrorMinimumSpecialCharacters
		}
		for i := 0; i < options.MinimumSpecialCharacters; i++ {
			pos := rand.Intn(len(specialChars))
			buffer += specialChars[pos : pos+1]
		}
	}

	// Fill with all characters until the required password length is met
	if len(buffer) > passwordLength {
		return "", ErrorMinimumsTooHigh
	}
	if len(buffer) < passwordLength {
		set := ""
		if !options.DisableLowerLetters {
			set += charactersLowerLetters
		}
		if !options.DisableUpperLetters {
			set += charactersUpperLetters
		}
		if !options.DisableDigits {
			set += charactersDigits
		}
		if !options.DisableSpecialChars {
			set += specialChars
		}

		if len(set) == 0 {
			return "", ErrorEverythingDisabled
		}

		for len(buffer) < passwordLength {
			pos := rand.Intn(len(set))
			buffer += set[pos : pos+1]
		}
	}

	// Mix up the characters
	for len(buffer) > 0 {
		pos := rand.Intn(len(buffer))
		result += buffer[pos : pos+1]
		buffer = buffer[:pos] + buffer[pos+1:]
	}

	return result, nil
}

func MustGenerate(passwordLength int, options Options) string {
	password, err := Generate(passwordLength, options)
	if err != nil {
		panic(err)
	}
	return password
}
