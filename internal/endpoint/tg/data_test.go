package tg

import (
	"regexp"
	"strings"
	"testing"
)

func BenchmarkSplit(t *testing.B) {
	s := "next 2005-12-12"
	for i := 0; i < t.N; i++ {
		ss := strings.Split(s, " ")
		if ss[0] == "next" {
			_ = ss[1]
		}
	}
}

func BenchmarkRegEx(t *testing.B) {
	nextReg := regexp.MustCompile(`next_(.*)`)
	s := "next_2005-12-12"
	for i := 0; i < t.N; i++ {
		if nextReg.MatchString(s) {
			data := nextReg.FindAllStringSubmatch(s, 1)
			_ = data[0][1]
		}
	}
}
