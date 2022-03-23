package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
	"unicode"
)

var re = regexp.MustCompile(`(,|^-$|!|\.|,|:|\")`)

func Top10(s string) []string {
	fs := make(map[string]int) // word by frequency index map

	type wi struct { // the word and it`s index
		Word string
		Inx  int
	}

	ss := strings.Fields(s) // split input string by whitespace
	for _, w := range ss {
		s := re.ReplaceAllString(w, "") // drop special letters
		if s != "" {                    // if string still valid
			us := []rune(s) // rune array for extract first rune only!
			for i, v := range us {
				// switch first rune to lower and concat in with string tail
				fs[string(unicode.ToLower(v))+string(us[i+1:])]++ // count word index
				break                                             // that enough!
			}
		}
	}

	wis := make([]wi, 0, 256) // slice of word index pair
	for w, i := range fs {    // convert map to slice of struct wi
		wis = append(wis, wi{w, i})
	}

	sort.Slice(wis, func(i, j int) bool {
		switch {
		case wis[i].Inx > wis[j].Inx: // sort by index
			return true
		case wis[i].Inx == wis[j].Inx: // if equal index sort by alphabet
			return wis[i].Word < wis[j].Word
		}
		return false
	})

	ret := make([]string, 0) // result strings by its rating
	i0 := 0
	for _, wi := range wis { // get first ten slots
		ret = append(ret, wi.Word)
		i0++
		if i0 >= 10 {
			break
		}
	}

	return ret
}
