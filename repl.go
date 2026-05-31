package main

import "strings"

func cleanInput(text string) []string {
	clean_string := []string{}

	trimmed := strings.TrimSpace(text)

	for _, s := range strings.Split(trimmed, " ") {
		if s != "" {
			clean_string = append(clean_string, strings.ToLower(s))
		}
	}

	return clean_string
}
