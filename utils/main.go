package utils

import "strings"

func Format422Error(error string) map[string]interface{} {
	errorSplited := strings.Split(error, "\n")

	result := make(map[string]interface{})

	for _, s := range errorSplited {

		splited := strings.Split(s, "'")

		key := splited[3]
		tag := splited[5]

		println(splited, key, tag)
		result[key] = tag
	}

	return result
}
