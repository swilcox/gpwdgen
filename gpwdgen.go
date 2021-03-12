package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GeneratePassPhrase(fileName string, numWords int, separator string, capitalize int, number int) string {
	// commandline parameters

	words := make([]string, 0)
	// words - list of available words
	if fileName == "" {
		words = internalWordList
	} else {
		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			words = append(words, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	// build list of words for the phrase
	pWords := make([]string, 0)
	for i := 0; i < numWords; i += 1 {
		pWords = append(pWords, words[rand.Intn(len(words))])
	}

	// capitalize a particular word
	if capitalize > 0 && capitalize <= len(words) {
		pWords[capitalize-1] = strings.Title(pWords[capitalize-1])
	}

	// make a random number to append
	var numString string = ""
	if number > 0 {
		numString = strconv.Itoa(rand.Intn(number))
	}
	return fmt.Sprintf("%s%s", strings.Join(pWords, separator), numString)
}

func main() {
	defaultFile := ""
	fileNamePtr := flag.String("filename", defaultFile, "filename for source words")
	numWordsPtr := flag.Int("words", 4, "number of words")
	separatorPtr := flag.String("separator", "-", "word separator")
	capitalizePtr := flag.Int("capitalize", 1, "word to capitalize (0 = none)")
	numberPtr := flag.Int("number", 9, "maximum random number to append (0 = none)")
	flag.Parse()

	// random seed
	rand.Seed(time.Now().Unix())

	// output the final passphrase
	passPhrase := GeneratePassPhrase(*fileNamePtr, *numWordsPtr, *separatorPtr, *capitalizePtr, *numberPtr)
	fmt.Println(passPhrase)
}
