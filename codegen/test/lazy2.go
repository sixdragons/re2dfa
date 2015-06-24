// Code generated by re2dfa (https://github.com/opennota/re2dfa).

package test

import "unicode/utf8"

func matchLazy2(s string) (end int) {
	end = -1
	var r rune
	var rlen int
	i := 0
	lazy := false
	type jmp struct{ s, i int }
	var lazyArr [1]jmp
	lazyStack := lazyArr[:0]
	_, _, _ = r, rlen, i
s1:
	if lazy {
		lazy = false
		goto s2
	}
	lazyStack = append(lazyStack, jmp{s: 1, i: i})
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 98:
		end = i
	}
	goto bt
s2:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 97:
		goto s3
	}
	goto bt
s3:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 98:
		end = i
	}
bt:
	if end >= 0 || len(lazyStack) == 0 {
		return
	}
	var to jmp
	to, lazyStack = lazyStack[len(lazyStack)-1], lazyStack[:len(lazyStack)-1]
	lazy = true
	i = to.i
	switch to.s {
	case 1:
		goto s1
	}
	return
}