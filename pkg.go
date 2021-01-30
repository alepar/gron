package main

import (
	"fmt"
	"io"
	"regexp"
)

var linePattern = regexp.MustCompile("^json\\.([^=]+) = (.+);$")

func Gron(r io.Reader) (map[string]string, error) {
	resultMap := make(map[string]string)

	ss, err := statementsFromJSON(r, statement{{"json", typBare}})
	if err != nil {
		goto out
	}

	for i, s := range ss {
		if i == 0 {
			continue
		}

		line := statementToString(s)
		matches := linePattern.FindStringSubmatch(line)
		if matches == nil {
			err = fmt.Errorf("could not parse line: %s", line)
			goto out
		}

		key := matches[1]
		value := matches[2]
		if value != "{}" && value != "[]" {
			resultMap[key] = value
		}
	}

out:
	if err != nil {
		return nil, fmt.Errorf("failed to Gron: %s", err)
	}
	return resultMap, nil
}
