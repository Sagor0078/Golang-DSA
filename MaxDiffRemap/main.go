import (
	"fmt"
	"strconv" 
	"strings" 
)


func minMaxDifference(num int) int {
    sNum := strconv.Itoa(num) 

    var firstNon9 rune = 0
    var firstNon0 rune = 0

    for _, r := range sNum {
        if firstNon9 == 0 && r >= '0' && r <= '8' {
            firstNon9 = r
        }
        if firstNon0 == 0 && r >= '1' && r <= '9' {
            firstNon0 = r
        }
    }

    maxStr := sNum 
    if firstNon9 != 0 { 
        maxStr = strings.ReplaceAll(maxStr, string(firstNon9), "9")
    }

    minStr := sNum 
    if firstNon0 != 0 { 
        minStr = strings.ReplaceAll(minStr, string(firstNon0), "0")
    }

    maxVal, _ := strconv.Atoi(maxStr)
    minVal, _ := strconv.Atoi(minStr)

    return maxVal - minVal
}