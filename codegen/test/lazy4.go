// Code generated by re2dfa (https://github.com/opennota/re2dfa).

package test

import "unicode/utf8"

func matchLazy4(s string) (end int) {
	end = -1
	var r rune
	var rlen int
	i := 0
	lazyOn := false
	type jmp struct{ s, i int }
	var lazyArr [2]jmp
	lazyStack := lazyArr[:0]
	_, _, _ = r, rlen, i
s1:
	if lazyOn {
		lazyOn = false
		goto s2
	}
	lazyStack = append(lazyStack, jmp{s: 1, i: i})
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto lazy
	}
	i += rlen
	switch {
	case r == 98:
		end = i
	}
	goto lazy
s2:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto lazy
	}
	i += rlen
	switch {
	case r == 97:
		goto s3
	}
	goto lazy
s3:
	if lazyOn {
		lazyOn = false
		goto s2
	}
	lazyStack = append(lazyStack, jmp{s: 3, i: i})
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto lazy
	}
	i += rlen
	switch {
	case r == 98:
		end = i
	}
lazy:
	if end >= 0 || len(lazyStack) == 0 {
		return
	}
	var to jmp
	to, lazyStack = lazyStack[len(lazyStack)-1], lazyStack[:len(lazyStack)-1]
	lazyOn = true
	i = to.i
	switch to.s {
	case 1:
		goto s1
	case 3:
		goto s3
	}
	return
}
