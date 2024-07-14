package strings

import "strings"

func CenterString(str string, width int) string {
	spaces := int(float64(width-len(str)) / 2)
	return strings.Repeat(" ", spaces) + str + strings.Repeat(" ", width-(spaces+len(str)))
}

func CenterStringBetween(leftStr, centerStr, rightStr string, width int) string {
	totalLen := len(leftStr) + len(centerStr) + len(rightStr)
	spaces := int(float64(width-totalLen) / 2)
	if spaces < 0 {
		return leftStr + centerStr + rightStr
	}
	return leftStr +
		strings.Repeat(" ", spaces) + centerStr + strings.Repeat(" ", width-(spaces+totalLen)) +
		rightStr
}
