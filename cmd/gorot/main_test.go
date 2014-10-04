package main

import "testing"

func init() {
	exit = func(_ int) {}
}

func Test_main(t *testing.T) {
	main()
}
