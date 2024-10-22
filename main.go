package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	outputflag := flag.String("align", "", "align file name")
	flag.Parse()

	if *outputflag == "" {
		fmt.Println("Hata: --align=<fileName.txt> bayraği ile bir dosya adi belirtmelisiniz.")
		os.Exit(1)
	}

	word := os.Args[2]
	words := os.Args[3]
	word = strings.ReplaceAll(word, "\\n", "\n")

	var fileContent string
	if words == "standard" {
		file, err := os.ReadFile("standard.txt")
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file)
	} else if words == "shadow" {
		file, err := os.ReadFile("shadow.txt")
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file)
	} else if words == "thinkertoy" {
		file, err := os.ReadFile("thinkertoy.txt")
		if err != nil {
			fmt.Println("Dosya okunurken bir hata oluştu")
			panic(err)
		}
		fileContent = string(file)
	} else {
		fmt.Println("Geçersiz kelime grubu:", words)
		return
	}

	lines := strings.Split(fileContent, "\n")

	for i, line := range strings.Split(word, "\n") {
		if line == "" {
			if i != 0 {
				fmt.Println()
			}
			continue
		}

		for h := 1; h < 9; h++ {
			for _, char := range line {
				printAsciiArtForCharacter(char, h, lines, *outputflag)
			}
			fmt.Println()
		}
	}
}

func printAsciiArtForCharacter(char rune, lineIndex int, lines []string, alignFlag string) {
	index := (int(char) - 32) * 9

	if index >= 0 && index+8 <= len(lines) {
		switch {
		case strings.Contains(alignFlag, "center"):
			fmt.Print(center(lines[index+lineIndex], 255))
		case strings.Contains(alignFlag, "left"):
			fmt.Print(left(lines[index+lineIndex], 255))
		case strings.Contains(alignFlag, "right"):
			fmt.Print(right(lines[index+lineIndex], 255))
		default:
			fmt.Print(lines[index+lineIndex])
		}
	}
}

func center(word string, termWidth int) string {
	lines := strings.Split(word, "\n")
	var centeredText string

	for _, line := range lines {
		padding := (termWidth - len(line)) / 2
		centeredLine := strings.Repeat(" ", padding) + line + strings.Repeat(" ", padding)
		centeredText += centeredLine + "\n"
	}
	return centeredText
}

func left(word string, termWidth int) string {
	lines := strings.Split(word, "\n")
	var leftAlignedText string

	for _, line := range lines {
		leftAlignedText += line + "\n"
	}
	return leftAlignedText
}

func right(word string, termWidth int) string {
	lines := strings.Split(word, "\n")
	var rightAlignedText string

	for _, line := range lines {
		padding := termWidth - len(line)
		rightAlignedLine := strings.Repeat(" ", padding) + line
		rightAlignedText += rightAlignedLine + " \n"
	}
	return rightAlignedText
}
