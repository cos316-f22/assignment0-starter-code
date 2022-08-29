package main

import (
	"crypto/sha256"
	"regexp"
	"runtime"
	"fmt"
	"io/ioutil"
	"os"

)

func main() {
	sha := sha256.New()
	previous, _ := ioutil.ReadAll(os.Stdin)
	sha.Write(previous)
	r := regexp.MustCompile(`go[0-9]+.[0-9]+`)
	version := r.FindStringSubmatch(runtime.Version())[0]
	sha.Write([]byte(version))
	sum := sha.Sum(nil)
	fmt.Printf("%x", sha256.Sum256([]byte(sum)))
}
