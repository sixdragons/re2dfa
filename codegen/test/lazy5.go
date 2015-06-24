// Code generated by re2dfa (https://github.com/opennota/re2dfa).

package test

import "unicode/utf8"

func matchLazy5(s string) (end int) {
	end = -1
	var r rune
	var rlen int
	i := 0
	lazy := false
	type jmp struct{ s, i int }
	var lazyArr [1]jmp
	lazyStack := lazyArr[:0]
	_, _, _ = r, rlen, i
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 97:
		end = i
		goto s2
	}
	goto bt
s2:
	if lazy {
		lazy = false
		goto s3
	}
	lazyStack = append(lazyStack, jmp{s: 2, i: i})
	goto bt
s3:
	r, rlen = utf8.DecodeRuneInString(s[i:])
	if rlen == 0 {
		goto bt
	}
	i += rlen
	switch {
	case r == 97:
		end = i
		goto s2
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
	case 2:
		goto s2
	}
	return
}