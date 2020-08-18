package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)
const (
	exitOK = iota
	exitOpenFile
	exitReadInput
	exitJSON
	)

// Objective of this is to parse the Burp suite scope files and parse all lines in STDIN from it, returning the ones that match.

func init() {
	flag.Usage = func() {
		h := "Apply a Burp scope JSON file to input from STDIN and return the in-scope results.\n\n"
		h += "Usage:\n"
		h += "  hiccup [FILE]\n\n"
		h += "Exit Codes:\n"
		h += fmt.Sprintf("  %d\t%s\n", exitOK, "OK")
		h += fmt.Sprintf("  %d\t%s\n", exitOpenFile, "Failed to open scope file")
		h += fmt.Sprintf("  %d\t%s\n", exitReadInput, "Failed to read input")
		h += fmt.Sprintf("  %d\t%s\n", exitJSON, "Failed to decode JSON")
		h += "\n"

		h += "Examples:\n"
		h += "  assetfinder --subs-only google.com | hiccup scope.json\n"

		fmt.Fprintf(os.Stderr, h)
	}
}

func main() {
	invert := flag.Bool("v", false, "Invert the match type")
	matchProtocol := flag.Bool("p", false, "Match based on protocol in addition to hosts")
	flag.Parse()

	args := os.Args[1:]
	if len(args) != 1 {
		flag.Usage()
		return
	}
	scopeFileName := args[0]

	scopeFile, err := ParseFile(scopeFileName)
	if err != nil {
		panic(err)
	}

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	var lines []string

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(input))
	}

	scopeFile.CheckScope(*matchProtocol, *invert, lines)


}