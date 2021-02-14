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

func main() {
	// commandline parameters
	defaultFile := ""
	fileNamePtr := flag.String("filename", defaultFile, "filename for source words")
	numWordsPtr := flag.Int("words", 4, "number of words")
	separatorPtr := flag.String("separator", "-", "word separator")
	capitalizePtr := flag.Int("capitalize", 1, "word to capitalize (0 = none)")
	numberPtr := flag.Int("number", 9, "maximum random number to append (0 = none)")
	flag.Parse()

	words := make([]string, 0)
	// words - list of available words
	if *fileNamePtr == "" {
		words = internalWordList
	} else {
		f, err := os.Open(*fileNamePtr)
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

	// random seed
	rand.Seed(time.Now().Unix())

	// build list of words for the phrase
	pWords := make([]string, 0)
	for i := 0; i < *numWordsPtr; i += 1 {
		pWords = append(pWords, words[rand.Intn(len(words))])
	}

	// capitalize a particular word
	if *capitalizePtr > 0 && *capitalizePtr <= len(words) {
		pWords[*capitalizePtr-1] = strings.Title(pWords[*capitalizePtr-1])
	}

	// make a random number to append
	var numString string = ""
	if *numberPtr > 0 {
		numString = strconv.Itoa(rand.Intn(*numberPtr))
	}

	// output the final passphrase
	fmt.Printf("%s%s\n", strings.Join(pWords, *separatorPtr), numString)
}
