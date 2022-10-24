package lib

import "testing"

func TestReplaceAllCommasAndPeriods(t *testing.T) {
	isMatch := func(input, expect string) {
		output := ReplaceAll(input)
		if !(output == expect) {
			t.Error("input:", input, "|", "output:", output, "|", "expect:", expect)
		}
	}

	isMatch(",   s", "，s")
	isMatch("あ.", "あ．")
	isMatch("あ。", "あ．")
	isMatch("G.711", "G.711")
	isMatch("あさ, ひる。    よる、  ", "あさ，ひる．よる，")
}
