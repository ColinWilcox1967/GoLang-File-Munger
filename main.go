package main

import (
	"fmt"
	"flag"
	"os"
	"strconv"
	"strings"
	"io/ioutil"

	fileutils "./fileutils"
)

const (
	VERSION = "1.0"
	DEFAULT_BLOCK_SIZE  = 1024
)

var (
	filePath string
	maxBlockSize int
	keyPhrase string
)

func showBanner() {
	fmt.Printf ("File Munger, Version %s.\n\n", VERSION)
}

func doFileMunge(data []byte) []byte {
	words := strings.Fields(keyPhrase)

	fmt.Printf ("Applying %d interations over %d bytes ...\n", len(words), maxBlockSize)
	for iteration := 0; iteration < len(words); iteration++ {
		var offset int
		for blockIndex := 0; blockIndex < maxBlockSize; blockIndex++ {
			data[blockIndex] = data[blockIndex] ^ words[iteration][offset]
			offset++
			if offset == len(words) {
				offset = 0
			}
		}
	}

	return data
}



func getArguments () {
	var blockStr string

	flag.StringVar (&filePath, "file", "", "Path of file to be munged.")
	flag.StringVar (&blockStr, "n", "1024", "Maximum size of data block to munge (in bytes).")
	flag.StringVar (&keyPhrase, "key", "", "Keyphrase used as part of munging.")

	flag.Parse ()

	var err error
	maxBlockSize, err = strconv.Atoi(blockStr)
	if err != nil || maxBlockSize < 0 {
		maxBlockSize = DEFAULT_BLOCK_SIZE
	}

	if len(filePath) == 0 {
		fmt.Printf ("No data file specified.\n")
		os.Exit(-2)
	}

	if !fileutils.FileExists(filePath) {
		fmt.Printf ("Specified file '%s' does not exist.\n", strings.ToUpper(filePath))
		os.Exit(-2)
	}

	if keyPhrase == "" {
		fmt.Printf ("No key phrase specified.\n")
		os.Exit(-3)
	}

}

func main () {

	showBanner()
	getArguments()

	ok, data := fileutils.ReadFileAsBytes(filePath)
	if !ok {
		fmt.Printf ("Unable to read file '%s'.\n", filePath)
		os.Exit(-1)
	}

	if maxBlockSize > len(data) {
		maxBlockSize = len(data)
	}

	fmt.Printf ("Munging '%s' ...\n", strings.ToUpper(filePath))

	fmt.Println (data)

	// the good stuff
	data = doFileMunge(data)

	outputFile := fileutils.GetFileNameWithoutExtension(filePath)
	outputFile += ".OUT"

	err := ioutil.WriteFile(outputFile, data, 0644)
	if err != nil { 
		fmt.Printf ("Unable to write to file '%s'.\n", outputFile)
		os.Exit(-5)
	}
	
}