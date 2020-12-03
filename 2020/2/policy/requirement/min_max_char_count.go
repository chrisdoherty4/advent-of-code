package requirement

type MinMaxCharCount struct {
	char    rune
	minimum int
	maximum int
}

func NewMinMaxCharCount(char rune, min, max int) MinMaxCharCount {
	return MinMaxCharCount{
		char:    char,
		minimum: min,
		maximum: max,
	}
}

func (m MinMaxCharCount) Validate(stringToValidate string) bool {
	if !m.stringHasEnoughChars(stringToValidate) {
		return false
	}

	count := m.countExpectChar(stringToValidate)
	return m.isInsideBounds(count)
}

func (m MinMaxCharCount) isInsideBounds(count int) bool {
	return count >= m.minimum && count <= m.maximum
}

func (m MinMaxCharCount) countExpectChar(s string) int {
	count := 0
	for _, c := range s {
		if c == m.char {
			count++
		}
	}

	return count
}

func (m MinMaxCharCount) stringHasEnoughChars(stringToValidate string) bool {
	return len(stringToValidate) > m.minimum
}
