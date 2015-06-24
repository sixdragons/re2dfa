// Code generated by re2dfa (https://github.com/opennota/re2dfa).

package test

import "unicode/utf8"

func matchLazy3(s string) (end int) {
	end = 0
	var r rune
	var rlen int
	i := 0
	lazy := false
	type jmp struct{ s, i int }
	var lazyArr [2]jmp
	lazyStack := lazyArr[:0]
	_, _, _ = r, rlen, i
s1:
	if lazy {
		lazy = false
		goto s2
	}
	lazyStack = append(lazyStack, jmp{s: 1, i: i})
	goto bt
s2:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 97:
		end = i
		goto s3
	}
	goto bt
s3:
	if lazy {
		lazy = false
		goto s2
	}
	lazyStack = append(lazyStack, jmp{s: 3, i: i})
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
	case 3:
		goto s3
	}
	return
}