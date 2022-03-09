package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindEmail(t *testing.T) {
	text := "FindString возвращает строку, содержащую текст самого левого совпадения в s обычного выражение. Если совпадения нет, возвращаемое значение представляет собой пустую строку, но он также будет пустым, если регулярное выражение успешно соответствует umekovazamat@gmail.com пустая строка. Используйте FindStringIndex или FindStringSubmatch, если это необходимо различать эти случаи."
	email := "umekovazamat@gmail.com"
	e := FindEmailFromText(text)
	require.Equal(t, email, e)
}

func TestFindIIN(t *testing.T) {
	text := "FindString returns a string holding the text of the leftmost match in s of the regular 910130233220 expression. If there is no match, the return value is an empty string, but it will also be empty if the regular expression successfully matches umekovazamat@gmail.com an empty string. Use FindStringIndex or FindStringSubmatch if it is necessary to distinguish these cases."
	email := "910130233220"
	e := FindIINFromText(text)
	require.Equal(t, email, e)
}
