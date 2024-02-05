package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// buildBalloon takes a slice of strings of max width maxWidth
// prepends/appends margins on first and last line, and at start/end of each line
// and returns a string with the contents of the balloon
// Basically just a format function
func buildBalloon(lines []string, maxWidth int) string {
	var borders []string
	count := len(lines)
	var res []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxWidth+2)
	bottom := " " + strings.Repeat("-", maxWidth+2)

	res = append(res, top)

	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		res = append(res, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], lines[0], borders[1])
		res = append(res, s)
		for i := 1; i < count-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[4], lines[i], borders[4])
			res = append(res, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[2], lines[count-1], borders[3])
		res = append(res, s)
	}

	res = append(res, bottom)
	return strings.Join(res, "\n")
}

// tabsToSpaces converts all tabs found in the strings founds
// in the `lines` slices to 4 spaces, to prevent misalignments in counting the runes
func tabsToSpaces(lines []string) []string {
	var res []string
	for _, line := range lines {
		res = append(res, strings.Replace(line, "\t", "    ", -1))
	}
	return res
}

// calculateMaxWidth returns the length of the longest line in the `lines` slice
func calculateMaxWidth(lines []string) int {
	max := 0
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}
	return max
}

func normalizeStringsLength(lines []string, maxWidth int) []string {
	var res []string
	for _, line := range lines {
		if len(line) < maxWidth {
			line += strings.Repeat(" ", maxWidth-len(line))
		}
		res = append(res, line)
	}
	return res
}

func main() {
	info, _ := os.Stdin.Stat()

	// Check if the input is from a pipe
	if info.Mode()&os.ModeCharDevice != 0 {
		println("The command is intended to work with pipes.")
		println("Usage: fortune | gocowsay")
		return
	}

	// Read the input from the pipe
	reader := bufio.NewReader(os.Stdin)
	var output []string

	for {
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, string(input))
	}

	// Cow image
	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	// Format the output
	output = tabsToSpaces(output)
	maxWidth := calculateMaxWidth(output)
	messages := normalizeStringsLength(output, maxWidth)
	balloon := buildBalloon(messages, maxWidth)
	fmt.Println(balloon)
	fmt.Println(cow)
	fmt.Println()
}
