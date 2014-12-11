// String array utility functions
package strarr

import (
	"regexp"
	"strings"
)

func TrimAll(items []string) {
	for i, item := range items {
		items[i] = strings.Trim(item, " \n\r\t")
	}
}

func IndexOf(items []string, query string) int {
	for i, val := range items {
		if val == query {
			return i
		}
	}
	return -1
}

func Contains(items []string, query string) bool {
	return IndexOf(items, query) >= 0
}

func FindMatchWithRegex(items []string, regex string) (string, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	for _, item := range items {
		found := re.FindString(item)
		if found != "" {
			return found, nil
		}
	}
	return "", nil
}
