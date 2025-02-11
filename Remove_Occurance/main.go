package main

func removeOccurrences(s string, part string) string {
	targetLen := len(part)
	endChar := part[len(part)-1]
	var resultStack []rune

	for _, currentChar := range s {
		resultStack = append(resultStack, currentChar)

		if currentChar == rune(endChar) && len(resultStack) >= targetLen {
			if string(resultStack[len(resultStack)-targetLen:]) == part {
				resultStack = resultStack[:len(resultStack)-targetLen]
			}
		}
	}
	return string(resultStack)
}
