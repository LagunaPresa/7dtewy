package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/LagunaPresa/7dtewy/codec"
)

// $ go run . lipps duxrj
// hello world
func main() {
	verbose := flag.Bool("v", false, "Display verbose messages")
	encode := flag.Bool("e", false, "Encode arguments")
	flag.Parse()

	words := flag.Args()
	if *encode {
		encoded, ok := process(words, codec.EncodeCandidates)
		printResult(*verbose, "plain", words, "encoded", encoded)
		if !ok {
			os.Exit(1)
		}
	} else {
		decoded, ok := process(words, codec.DecodeCandidates)
		printResult(*verbose, "encoded", words, "plain", decoded)
		if !ok {
			os.Exit(1)
		}
	}
}

func process(words []string, codec func(string) ([]string, error)) (ss []string, ok bool) {
	ok = true
	for _, w := range words {
		if cands, err := codec(w); err != nil {
			ss = append(ss, fmt.Sprintf("[%s]", err.Error()))
			ok = false
		} else {
			ss = append(ss, format(cands))
		}
	}
	return
}

func format(ss []string) string {
	if l := len(ss); l == 1 {
		return ss[0]
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, ","))
}

func printResult(verbose bool, inputName string, input []string, outputName string, output []string) {
	if verbose {
		fmt.Printf("[input: %s]\n", inputName)
		fmt.Printf("%s\n\n", strings.Join(input, " "))
		fmt.Printf("[output: %s]\n", outputName)
		fmt.Printf("%s\n\n", strings.Join(output, " "))
		fmt.Println("[mapping]")
		for i := range input {
			fmt.Printf("%s -> %s\n", input[i], output[i])
		}
	} else {
		fmt.Println(strings.Join(output, " "))
	}
}
