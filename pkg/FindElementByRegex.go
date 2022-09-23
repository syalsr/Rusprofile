package pkg

import (
	"errors"
	"regexp"
)

func FindElementByRegex(doc string, pattern string) (string, error) {
	re := regexp.MustCompile(pattern)
	sub := re.FindAllStringSubmatch(doc, -1)
	if len(sub) == 0 {
		return "", errors.New("")
	}
	return sub[0][0], nil
}
