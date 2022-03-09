package util

import (
	"regexp"
)

// [a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$
// [a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}
const EMAIL_REGEX = "[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}"
const IIN_REGEX = `(\d{2}(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01]))[0-9]{6}`

func Increment(num uint64) uint64 {
	return num + 1
}

func Decrement(num uint64) uint64 {
	if num == 0 {
		return 0
	}

	return num - 1
}

func RandomStringList() []string {
	return []string{}
}

func FindEmailFromText(text string) string {
	emailRegexp := regexp.MustCompile(EMAIL_REGEX)
	first := emailRegexp.FindString(text)
	return first
}

func FindAllEmailFromText(text string) []string {
	emailRegexp := regexp.MustCompile(EMAIL_REGEX)
	all := emailRegexp.FindAllString(text, -1)
	return all
}

func FindIINFromText(text string) string {
	iinRegexp := regexp.MustCompile(IIN_REGEX)
	first := iinRegexp.FindString(text)
	return first
}
