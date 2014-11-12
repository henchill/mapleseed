package main

import (
	"testing"
	"fmt"
)

func Test_1(t *testing.T) {

	// can only run one in a process, and there's no way
	// to shut it down in normal go.   Could use 'graceful'
	go serve("http://example.com", ":8090") 



	// see sandro@waldron:~/go/src/github.com/sandhawke$ more podtester/jm_test.go

	fmt.Printf("running\n")
	// if g3.URL() != "http://pod1.example/auto/2" { t.Fail() }
	
	
}