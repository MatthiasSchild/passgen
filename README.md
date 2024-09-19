# passgen for Go

A golang library to generate passwords with specified requirements.
With this library, you can specify your requirements for a password and then generate one.

Here is an example how to use it:

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

You pass your required password length to the Generate function with additional options.
The options are your personal safety requirements for the password, e.g. how many digits the password
should have at least.

There is also a function `MustGenerate`. It works the same, but does not return an error.
It will panic instead.

## Options

- `DisableLowerLetters`: Don't add any lower-case letters
- `MinimumLowerLetters`: Add at least x lower-case letter
- `DisableUpperLetters`:Don't add any upper letters
- `MinimumUpperLetters`: Add at least x upper-case letter
- `DisableDigits`: Don't add any digits
- `MinimumDigits`: Add at least x digits
- `DisableSpecialChars`: Don't add any special chars
- `MinimumSpecialCharacters`: Add at least x special chars
- `SpecialCharacters`: Defines, what special chars should be used
- `ExcludeCharacters`: Exclude specific special chars from the selection set

The default special chars are `!"#$%&’()*+,-./:;<=>?@[\]^_``{|}~`, but you can define
you own special chars using `SpecialCharacters`, or alternatively remove special chars
from the set with `ExcludeCharacters`, when you don't want it in your password
(e.g. because of some issues with specific platforms).
