package dandistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EditDistance(t *testing.T) {
	cases := []struct {
		name                string
		word1               string
		word2               string
		expectedDistance    int
		expectedErrorString string
	}{
		{
			name:             "horse to ros",
			word1:            "horse",
			word2:            "ros",
			expectedDistance: 3,
		},
		{
			name:             "intention to execution",
			word1:            "intention",
			word2:            "execution",
			expectedDistance: 5,
		},
		{
			name:             "abcd to acdb",
			word1:            "abcd",
			word2:            "acdb",
			expectedDistance: 2,
		},
		{
			name:             "word1 is empty string",
			word1:            "",
			word2:            "converted",
			expectedDistance: 9,
		},
		{
			name:             "word2 is empty string",
			word1:            "string",
			word2:            "",
			expectedDistance: 6,
		},
		{
			name:                "string has capitals",
			word1:               "CAPITALS",
			word2:               "string",
			expectedDistance:    -1,
			expectedErrorString: "string a contains a character that is not lower case English",
		},
		{
			name:                "string has numeric characters",
			word1:               "string",
			word2:               "1234",
			expectedDistance:    -1,
			expectedErrorString: "string b contains a character that is not lower case English",
		},
		{
			name:                "string has special characters",
			word1:               "!@#$%",
			word2:               "string",
			expectedDistance:    -1,
			expectedErrorString: "string a contains a character that is not lower case English",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := EditDistance(c.word1, c.word2)

			assert.Equal(t, c.expectedDistance, actual)
			if err != nil {
				assert.EqualError(t, err, c.expectedErrorString)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
