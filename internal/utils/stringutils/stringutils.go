package stringutils

import "strings"

// GetFirstLine

func GetFirstLine(str string) string {
	return strings.Trim(str[:strings.Index(str, "\n")], " ")
}

func GetFirstLineCut(str string) string {
	before, _, _ := strings.Cut(str, "\n")
	return strings.Trim(before, " ")
}

func GetFirstLineSplit(str string) string {
	return strings.Trim(strings.Split(str, "\n")[0], " ")
}

func GetFirstLineSplitN(str string) string {
	return strings.Trim(strings.SplitN(str, "\n", 1)[0], " ")
}

// GetStringContent

func GetStringContent(bytes []byte) string {
	return string(bytes)
}

func GetStringContentWithLengthParameter(bytes []byte, length int) string {
	return string(bytes[:length])
}

func GetStringContentWithLengthEvaluated(bytes []byte) string {
	length := len(bytes)
	return string(bytes[:length])
}
