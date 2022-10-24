package lib

import (
	"regexp"
	"strings"
)

// replace all commas and periods in the line
func ReplaceAll(str string) string {
	reComma := regexp.MustCompile(`(,|、)( |　)*`)
	rePeriod := regexp.MustCompile(`(\W)(\.|。)( |　)*`)

	buf := reComma.ReplaceAllString(str, "，")
	result := rePeriod.ReplaceAllString(buf, "$1．")

	return result
}

func findDocIdx(src []string) (int, int) {
	beginDocRe := regexp.MustCompile(`\\begin\{\s*document\s*\}`)
	endDocRe := regexp.MustCompile(`\\end\{\s*document\s*\}`)

	var beginIdx, endIdx int
	for i, v := range src {
		bufBeginIdx := beginDocRe.FindIndex([]byte(v))
		bufEndIdx := endDocRe.FindIndex([]byte(v))

		if len(bufBeginIdx) > 0 {
			beginIdx = i
		} else if len(bufEndIdx) > 0 {
			endIdx = i
			break
		}
	}

	return beginIdx, endIdx
}

func parseLines(src []string) []string {
	beginIdx, endIdx := findDocIdx(src)

	var dest []string
	dest = append(dest, src[:beginIdx+1]...)

	commentRe := regexp.MustCompile(`^\s*%`)
	citeRe := regexp.MustCompile(`^(\s)*\\cite\{.+\}`)
	itemizeBeginRe := regexp.MustCompile(`^\\begin\{itemize\}`)
	itemizeEndRe := regexp.MustCompile(`^\\end\{itemize\}`)
	enumerateBeginRe := regexp.MustCompile(`^\\begin\{enumerate\}`)
	enumerateEndRe := regexp.MustCompile(`^\\end\{enumerate\}`)
	figureBeginRe := regexp.MustCompile(`^\\begin\{figure.*\}`)
	figureEndRe := regexp.MustCompile(`^\\end\{figure.*\}`)
	tableBeginRe := regexp.MustCompile(`^\\begin\{table.*\}`)
	tableEndRe := regexp.MustCompile(`^\\end\{table.*\}`)
	captionRe := regexp.MustCompile(`^(\s)*\\caption\{.+\}`)
	commandBlockBeginRe := regexp.MustCompile(`^\\begin\{.+\}`)
	commandBlockEndRe := regexp.MustCompile(`^\\end\{.+\}`)
	commandRe := regexp.MustCompile(`^\\`)

	isInItemize := false
	isInEnumerate := false
	isInFigureOrTable := false
	isInCommandBlock := false
	for _, v := range src[beginIdx+1 : endIdx] {
		// rule 2 (comment line)
		if commentRe.Match([]byte(v)) {
			dest = append(dest, v)
			continue
		}

		// rule 3 (cite)
		if citeRe.Match([]byte(v)) {
			dest = append(dest, ReplaceAll(v))
			continue
		}

		// rule 4 (command block)
		// itemize
		if itemizeBeginRe.Match([]byte(v)) {
			isInItemize = true
			dest = append(dest, v)
			continue
		} else if itemizeEndRe.Match([]byte(v)) {
			isInItemize = false
			dest = append(dest, v)
			continue
		} else if isInItemize {
			dest = append(dest, ReplaceAll(v))
			continue
		}

		// enumerate
		if enumerateBeginRe.Match([]byte(v)) {
			isInEnumerate = true
			dest = append(dest, v)
			continue
		} else if enumerateEndRe.Match([]byte(v)) {
			isInEnumerate = false
			dest = append(dest, v)
			continue
		} else if isInEnumerate {
			dest = append(dest, ReplaceAll(v))
			continue
		}

		// caption (rule 3)
		if figureBeginRe.Match([]byte(v)) || tableBeginRe.Match([]byte(v)) {
			isInFigureOrTable = true
			dest = append(dest, v)
			continue
		} else if figureEndRe.Match([]byte(v)) || tableEndRe.Match([]byte(v)) {
			isInFigureOrTable = false
			dest = append(dest, v)
			continue
		} else if isInFigureOrTable {
			if captionRe.Match([]byte(v)) {
				dest = append(dest, ReplaceAll(v))
				continue
			}
			dest = append(dest, v)
			continue
		}

		// ignore
		if commandBlockBeginRe.Match([]byte(v)) {
			isInCommandBlock = true
			dest = append(dest, v)
			continue
		} else if commandBlockEndRe.Match([]byte(v)) {
			isInCommandBlock = false
			dest = append(dest, v)
			continue
		} else if isInCommandBlock {
			dest = append(dest, v)
			continue
		}

		// rule 3 (command line)
		if commandRe.Match([]byte(v)) {
			dest = append(dest, v)
			continue
		}

		dest = append(dest, ReplaceAll(v))
	}

	dest = append(dest, src[endIdx:]...)

	return dest
}

func Parse(srcText string) string {
	reLines := regexp.MustCompile("\r\n|\n")
	srcLines := reLines.Split(srcText, -1)

	destLines := parseLines(srcLines)
	destText := strings.Join(destLines, "\n")

	return destText
}
