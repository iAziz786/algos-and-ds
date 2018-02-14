// Link to the problem: https://www.codechef.com/problems/LAPIN

package main

import (
	"fmt"
	"strings"
)

func main() {
	var T uint8
	fmt.Scan(&T)
	for ; T > 0; T-- {
		var str, str1, str2 string
		fmt.Scan(&str)
		oddOrEven := len(str)
		if oddOrEven%2 != 0 {
			str1 = str[:oddOrEven/2]
			str2 = str[oddOrEven/2+1:]
		} else {
			str1 = str[:oddOrEven/2]
			str2 = str[oddOrEven/2:]
		}
		var isLapin = true
		for i := 0; i < (oddOrEven / 2); i++ {
			val1 := strings.Count(str1, string(str1[i]))
			val2 := strings.Count(str2, string(str1[i]))
			if val1 != val2 {
				isLapin = false
			}
		}
		if isLapin == false {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
