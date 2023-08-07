package util

import "strings"

// FilterSQLInjection 过滤sql注入字符
func FilterSQLInjection(input string) string {
	// Remove any characters that are not allowed in SQL queries.
	filteredString := strings.Replace(input, "'", "''", -1)
	filteredString = strings.Replace(filteredString, "\"", "\"\"", -1)
	filteredString = strings.Replace(filteredString, "\\", "\\\\", -1)
	filteredString = strings.Replace(filteredString, ";", " ", -1)
	filteredString = strings.Replace(filteredString, "<", " ", -1)
	filteredString = strings.Replace(filteredString, ">", " ", -1)
	filteredString = strings.Replace(filteredString, "&", " ", -1)
	filteredString = strings.Replace(filteredString, "|", " ", -1)
	filteredString = strings.Replace(filteredString, "#", " ", -1)
	filteredString = strings.Replace(filteredString, "*", " ", -1)
	filteredString = strings.Replace(filteredString, "-", " ", -1)
	return filteredString
}
