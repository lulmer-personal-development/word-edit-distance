package dandistance

import (
	"fmt"
)

const MAX_STRING_LENGTH = 500

func EditDistance(a, b string) (int, error) {
	aLength := len(a)
	bLength := len(b)

	err := validateStrings(a, b, aLength, bLength)
	if err != nil {
		return -1, err
	}

	if aLength == 0 {
		return bLength, nil
	}
	if bLength == 0 {
		return aLength, nil
	}

	convertMatrix := make([][]int, aLength+1)
	for i := range convertMatrix {
		convertMatrix[i] = make([]int, bLength+1)
	}

	for i := 0; i <= aLength; i++ {
		convertMatrix[i][0] = i
	}
	for j := 0; j <= bLength; j++ {
		convertMatrix[0][j] = j
	}

	for i := 1; i <= aLength; i++ {
		for j := 1; j <= bLength; j++ {
			if a[i-1] == b[j-1] {
				convertMatrix[i][j] = convertMatrix[i-1][j-1]
			} else {
				convertMatrix[i][j] = min(convertMatrix[i-1][j-1], convertMatrix[i-1][j], convertMatrix[i][j-1]) + 1
			}
		}
	}

	return convertMatrix[aLength][bLength], nil
}

// Constraints: strings must be 0 to 500 characters and lower case English
func validateStrings(a, b string, aLength, bLength int) error {
	if aLength > MAX_STRING_LENGTH || bLength > MAX_STRING_LENGTH {
		return fmt.Errorf("input string exceeds the maximum length of %d", MAX_STRING_LENGTH)
	}

	for _, aChar := range a {
		if !isLowerCaseEnglishChar(aChar) {
			return fmt.Errorf("string a contains a character that is not lower case English")
		}
	}
	for _, bChar := range b {
		if !isLowerCaseEnglishChar(bChar) {
			return fmt.Errorf("string b contains a character that is not lower case English")
		}

	}
	return nil
}

func isLowerCaseEnglishChar(char rune) bool {
	return (char >= 'a' && char <= 'z')
}

// determine minimum
func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
