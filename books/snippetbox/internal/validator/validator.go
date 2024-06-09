package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Define a new Validator type which contains a map of validation errors for our form fields
type Validator struct {
	FieldErrors map[string]string
	NonFieldErrors []string

}

// Valid() returns true if the FieldErrors map doesn't contain any entries.
// Update the Valid() method to also check that the NonFieldErrors slice is
// empty
func (v *Validator) Validator() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0

}

// Create an AddNonFieldError() helper for adding error messages to the new
// NonFieldErrors slice.
func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
   }
   

// AddFieldError() adds an error message to the FieldErrors map (so long as no
// entry already exists for the given key).
func (v *Validator) AddFieldError(key, message string) {
	// Note: We need to initialize the map first, if it isn't already
	// initialized.
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// CheckField() adds an error message to the FieldErrors map only if a
// validation check is not 'ok'.
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// NotBlank() returns true if a value is not an empty string.
func (v *Validator) NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars() returns true if a value contains no more than n characters.
func (v *Validator) MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}


// Replace PermittedInt() with a generic PermittedValue() function. This returns
// true if the value of type T equals one of the variadic permittedValues
// parameters.
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
	if value == permittedValues[i] {
	return true
	}
	}
	return false
   }
// PermittedInt() returns true if a value is in a list of permitted integers.
// func PermittedInt(value int, permittedValues ...int) bool {
// 	for i := range permittedValues {
// 		if value == permittedValues[i] {
// 			return true
// 		}
// 	}
// 	return false
// }

// Use the regexp.MustCompile() function to parse a regular expression pattern
// for sanity checking the format of an email address. This returns a pointer to
// a 'compiled' regexp.Regexp type, or panics in the event of an error. Parsing
// this pattern once at startup and storing the compiled *regexp.Regexp in a
// variable is more performant than re-parsing the pattern each time we need it.

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z]{2,})+$")

// MinChars() returns true if a value contains at least n characters.
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

// Matches() returns true if a value matches a provided compiled regular
// expression pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
