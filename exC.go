package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func error() {
	fmt.Println("You must enter an argument corresponding to the name of a file.")
}

func getUrl(file *os.File) (Url []string) {
	reader := bufio.NewReader(file)
	replacer := strings.NewReplacer("\n", "")

	line := []byte("something")
	for len(line) != 0 {
		line, _ = reader.ReadBytes('\n')
		formattedUrl := replacer.Replace(string(line))
		Url = append(Url, formattedUrl)
	}
	return Url
}

func getContentFromUrl(url string) (JSonByte []byte) {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("We could not access to the url", url, "please make sure your file is correctly written.")
		return JSonByte
	}

	defer r.Body.Close()

	JSonByte, _ = ioutil.ReadAll(r.Body)
	return JSonByte
}

func isAlreadyInArray(needle string, haystack []string) (answer bool) {
	answer = false
	for _, element := range haystack {
		if element == needle {
			answer = true
			break
		}
	}
	return answer
}

func countAtInUrl(url string, waitGroupAdress *sync.WaitGroup, nbAt *int, sliceEmail *[]string, removeDuplicate bool) {
	defer waitGroupAdress.Done() //equivalent to defer *(waitGroupAdress).Done()
	if url != "" {
		haystack := getContentFromUrl(url)
		regExpEmail := regexp.MustCompile("([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+)")

		result := regExpEmail.FindAll(haystack, -1)
		for _, char := range haystack {
			if char == 64 {
				*nbAt++
			}
		}

		for _, email := range result {
			if !removeDuplicate || !isAlreadyInArray(string(email), *sliceEmail) {
				*sliceEmail = append(*sliceEmail, string(email))
			}
		}
	}
}

func displayWholeSlice(arrayString []string) (stringToDisplay string) {
	for key, littleString := range arrayString {
		var endChar string
		if key == (len(arrayString) - 1) {
			endChar = "."
		} else {
			if key == (len(arrayString) - 2) {
				endChar = " and\n"
			} else {
				endChar = ", \n"
			}
		}
		stringToDisplay += littleString + endChar
	}
	return stringToDisplay
}

func main() {
	if len(os.Args) >= 2 {
		removeDuplicate := false
		if len(os.Args) >= 3 {
			if os.Args[2] == "y" {
				removeDuplicate = true
			}
		}
		var nameFile string
		var endOfSentence string
		var waitGroup sync.WaitGroup
		var sliceEmail []string
		var secondSentence string

		nameFile = os.Args[1]
		file, err := os.Open(nameFile)
		defer file.Close()
		if err != nil {
			fmt.Print("something went wrong... Remember :")
			error()
			panic(err)
		}

		Urls := getUrl(file)

		nbAt := 0

		for _, url := range Urls {
			waitGroup.Add(1)
			go countAtInUrl(url, &waitGroup, &nbAt, &sliceEmail, removeDuplicate)
		}

		waitGroup.Wait()

		sort.Strings(sliceEmail)

		if nbAt > 1 {
			endOfSentence = "ats."
		} else {
			endOfSentence = "at."
		}

		fmt.Println("Going through the url linked to your file, we've found", nbAt, endOfSentence)

		if len(sliceEmail) == 0 {
			secondSentence = "Sorry, but we didn't found any email adress."
		} else {

			var emailAdress string
			if len(sliceEmail) == 1 {
				emailAdress = "email adress : "
			} else {
				emailAdress = "email adresses : "
			}

			secondSentence = "We also found " + strconv.Itoa(len(sliceEmail)) + " " + emailAdress + displayWholeSlice(sliceEmail)
			fmt.Println(secondSentence)
		}

		if len(sliceEmail) == 1 {
			secondSentence = "We found "
		}

	} else {
		error()
	}
}
