package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

type statLetter struct {
	Letter    []string
	Occurence int
}

func error() {
	fmt.Println("You must enter an argument corresponding to the name of a file")
}

func replaceIfOccurenceDiffers(structStatLetter *statLetter, key string, numberOfOccurence int) {
	if structStatLetter.Occurence == numberOfOccurence {
		structStatLetter.Letter = append(structStatLetter.Letter[0:], key)
	} else {
		structStatLetter.Letter = append(structStatLetter.Letter[:0], key)
		structStatLetter.Occurence = numberOfOccurence
	}
}

func pickWords(structStat statLetter) (subjectVerb string, beginning string, end string) {
	if len(structStat.Letter) == 1 {
		subjectVerb = "letter is"
		beginning = "It appears"
	} else {
		subjectVerb = "letters are :"
		beginning = "They each appear"
	}

	if structStat.Occurence == 1 {
		end = "time."
	} else {
		end = "times."
	}

	return subjectVerb, beginning, end
}

func displayWholeSlice(arrayString []string) (stringToDisplay string) {
	for key, littleString := range arrayString {
		var endChar string
		if key == (len(arrayString) - 1) {
			endChar = "."
		} else {
			endChar = ", "
		}
		stringToDisplay += littleString + endChar
	}
	return stringToDisplay
}

func main() {
	var nameFile string
	if len(os.Args) >= 2 {
		nameFile = os.Args[1]
		contentFile, err := ioutil.ReadFile(nameFile)
		if err != nil {
			fmt.Print("something went wrong... Remember : ")
			error()
			panic(err)
		}

		var endSize string

		if len(contentFile) <= 1 {
			endSize = "byte."
		} else {
			endSize = "bytes."
		}

		fmt.Println("Your file has a size of", len(contentFile), endSize)

		mapNumberOfLetter := make(map[string]int)

		for _, unicodeChar := range []rune(string(contentFile)) {
			if unicode.IsLetter(unicodeChar) {
				mapNumberOfLetter[string(unicodeChar)]++
			}
		}

		var maxLetter statLetter
		var minLetter statLetter

		maxLetter.Occurence = -1
		minLetter.Occurence = -1

		for key, numberOfOccurence := range mapNumberOfLetter {
			if maxLetter.Occurence == -1 || maxLetter.Occurence <= numberOfOccurence {
				replaceIfOccurenceDiffers(&maxLetter, key, numberOfOccurence)
			}

			if minLetter.Occurence == -1 || minLetter.Occurence >= numberOfOccurence {
				replaceIfOccurenceDiffers(&minLetter, key, numberOfOccurence)
			}
		}

		if len(mapNumberOfLetter) == 0 {
			fmt.Println("Your file didn't contain any letter, sorry.")
		} else {
			var subjectVerb string
			var beginning string
			var end string

			subjectVerb, beginning, end = pickWords(maxLetter)

			fmt.Println("The most frequent", subjectVerb, displayWholeSlice(maxLetter.Letter))
			fmt.Println(beginning, maxLetter.Occurence, end)

			subjectVerb, beginning, end = pickWords(minLetter)

			fmt.Println("The less frequent", subjectVerb, displayWholeSlice(minLetter.Letter))
			fmt.Println(beginning, minLetter.Occurence, end)
		}
	} else {
		error()
	}
}
