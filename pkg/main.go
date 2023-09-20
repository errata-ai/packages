package main

import "os"

func cdf() int {
	os.Chdir(os.Args[1])
	return 0
}
