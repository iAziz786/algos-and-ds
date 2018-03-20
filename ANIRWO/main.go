// Link to the problem: https://www.codechef.com/problems/ANIRWO

package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)
	var d = map[string]string{
		"A": "Z",
		"B": "Y",
		"C": "X",
		"D": "W",
		"E": "V",
		"F": "U",
		"G": "T",
		"H": "S",
		"I": "R",
		"J": "Q",
		"K": "P",
		"L": "O",
		"M": "N",
		"N": "M",
		"O": "L",
		"P": "K",
		"Q": "J",
		"R": "I",
		"S": "H",
		"T": "G",
		"U": "F",
		"V": "E",
		"W": "D",
		"X": "C",
		"Y": "B",
		"Z": "A",
		"a": "z",
		"b": "y",
		"c": "x",
		"d": "w",
		"e": "v",
		"f": "u",
		"g": "t",
		"h": "s",
		"i": "r",
		"j": "q",
		"k": "p",
		"l": "o",
		"m": "n",
		"n": "m",
		"o": "l",
		"p": "k",
		"q": "j",
		"r": "i",
		"s": "h",
		"t": "g",
		"u": "f",
		"v": "e",
		"w": "d",
		"x": "c",
		"y": "b",
		"z": "a",
		"0": "9",
		"1": "8",
		"2": "7",
		"3": "6",
		"4": "5",
		"5": "4",
		"6": "3",
		"7": "2",
		"8": "1",
		"9": "0",
	}
	for ; T > 0; T-- {
		var s string
		fmt.Scan(&s)
		for _, v := range s {
			fmt.Print(d[string(v)])
		}
		fmt.Println()
	}
}
