package main

import (
	"fmt"
	"strings"
)

const (
	ZERO WordSize = 6 * iota
	SMALL
	MEDIUM
	LARGE
	XLARGE
	XXLARGE  WordSize = 50
	SEPARATOR = ", "
)

type WordSize int

type ChainLink struct {
	Data []string
}

func (v *ChainLink) Value() []string {
	return v.Data
}

// stringFunc is a first-class method, used as a parameter to Map
type stringFunc func(s string) (result string)

// Map uses stringFunc to modify (up-case) each string in the slice
func (v *ChainLink) Map(fn stringFunc) *ChainLink {
	var mapped []string
	orig := *v
	for _, s := range orig.Data {
		mapped = append(mapped, fn(s))  // first-class function
	}
	v.Data = mapped
	return v
}

// Filter uses embedded logic to filter the slice of strings
// Note: We could have chosen to use a first-class function
func (v *ChainLink) Filter(max WordSize) *ChainLink {
	filtered := []string{}
	orig := *v
	for _, s := range orig.Data {
		if len(s) <= int(max) {             // embedded logic
			filtered = append(filtered, s)
		}
	}
	v.Data = filtered
	return v
}


func main() {
	constants := `
** Constants ***
ZERO: %v
SMALL: %d
MEDIUM: %d
LARGE: %d
XLARGE: %d
XXLARGE: %d
`
	fmt.Printf(constants, ZERO, SMALL, MEDIUM, LARGE, XLARGE, XXLARGE)

	words := []string{
		"tiny",
		"marathon",
		"philanthropinist",
		"supercalifragilisticexpialidocious"}

	data := ChainLink{words};
	fmt.Printf("unfiltered: %#v\n", data.Value())

	filtered := data.Filter(SMALL)
	fmt.Printf("filtered: %#v\n", filtered)

	fmt.Printf("filtered and mapped (<= SMALL sized words): %#v\n",
		filtered.Map(strings.ToUpper).Value())

	data = ChainLink{words}
	fmt.Printf("filtered and mapped (<= Up to MEDIUM sized words): %#v\n",
		data.Filter(MEDIUM).Map(strings.ToUpper).Value())

	data = ChainLink{words}
	fmt.Printf("filtered twice and mapped (<= Up to LARGE sized words): %#v\n",
		data.Filter(XLARGE).Map(strings.ToUpper).Filter(LARGE).Value())

	data = ChainLink{words}
	val := data.Map(strings.ToUpper).Filter(XXLARGE).Value()
	fmt.Printf("mapped and filtered (<= Up to XXLARGE sized words): %#v\n", val)
}
