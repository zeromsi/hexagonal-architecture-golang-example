package utility

import (
	"fmt"
	"net/url"
	"strings"
)

func ParseParams(str string) (map[string]string, error) {
	str, err := url.QueryUnescape(str)
	if err != nil {
		return nil, err
	}

	if strings.Contains(str, "?") {
		str = strings.Split(str, "?")[1]
	}

	if !strings.Contains(str, "=") {
		return nil, fmt.Errorf("\"%s\" contains no key-value pairs", str)
	}

	pairs := make(map[string]string)
	for _, pair := range strings.Split(string(str), "&") {
		items := strings.Split(pair, "=")
		pairs[items[0]] = items[1]
	}

	return pairs, nil
}

