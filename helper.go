package main

import "strconv"

func parseInt(s string) int {
	v, _ := strconv.ParseInt(s, 10, 32)
	return int(v)
}
