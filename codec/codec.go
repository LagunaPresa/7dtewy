package codec

import (
	"fmt"
	"math"
	"strings"
)

// DecodeCandidates returns candidates for the decoded string of `encoded`.
// If `encoded` contains characters outside the range [a-z], DecodeCandidates returns an error.
//
// This algorithm may be imperfect, and therefore, may not be able to determine a single "correct" word.
func DecodeCandidates(encoded string) ([]string, error) {
	if strings.ContainsFunc(encoded, func(e rune) bool { return e < 'a' || 'z' < e }) {
		return nil, fmt.Errorf(`[a-z] characters can be used: "%s"`, encoded)
	}
	var candidates []string
	cand := encoded
	// Not 26
	for i := 0; i < 25; i++ {
		if i == key(cand) {
			candidates = append(candidates, cand)
		}
		cand = shift(cand, -1)
	}
	return candidates, nil
}

// EncodeCandidates returns candidates for the encoded string of `plain`.
// If `plain` contains characters outside the range [a-y] (z excluded), EncodeCandidates returns an error.
//
// This algorithm may be imperfect, and therefore, may not be able to determine a single result.
//
// Details:
//
// When decoding a character, whether it is a or z does not seem to affect the result.
// Therefore, when encoding a character such that the result is z,
// I do not have an answer whether it is correct to set the result to z or a.
//
// Note that it is not possible to Encode words containing plain z.
// This codec algorithm cannot decode words containing plain z.
// Even in the game, words that should contain plain z do not appear when it is decoded.
func EncodeCandidates(plain string) ([]string, error) {
	if strings.ContainsFunc(plain, func(e rune) bool { return e < 'a' || 'y' < e }) {
		return nil, fmt.Errorf(`[a-y] characters can be used: "%s"`, plain)
	}
	basicEncoded := shift(plain, key(plain))
	replaced := strings.ReplaceAll(basicEncoded, "a", "z")
	if basicEncoded == replaced {
		return []string{basicEncoded}, nil
	}
	return []string{basicEncoded, replaced}, nil
}

func key(word string) int {
	i := score(word) - len(word)
	if i <= 0 {
		return 0
	}
	return int(math.Round(float64(i) / 12.0))
}

func score(word string) int {
	var s int32
	for _, char := range word {
		s += char + 1 - 'a'
	}
	return int(s)
}

// Similar to the Caesar cipher
func shift(word string, count int) string {
	var rs []rune
	for _, c := range word {
		cc := c + int32(count)%25
		if cc < 'a' {
			cc += 25
		} else if 'y' < cc {
			cc -= 25
		}
		rs = append(rs, cc)
	}
	return string(rs)
}
