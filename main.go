package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"jaytaylor.com/html2text"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//convert UTF-8 to BIG5
func EncodeBig5(f *os.File) (output []byte, err error) {
	output = []byte{}
	r := transform.NewReader(f, traditionalchinese.Big5.NewDecoder())
	// Read converted UTF-8 from `r` as needed.
	// As an example we'll read line-by-line showing what was read:
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		output = append(output, sc.Bytes()...)
		//fmt.Printf("Read line: %s\n", sc.Bytes())
	}
	if err = sc.Err(); err != nil {
		log.Fatal(err)
	}

	if err = f.Close(); err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	check(err)

	dat := []byte{}
	if len(os.Args) > 2 {
		encoding := os.Args[2]
		if encoding == "big5" {
			dat, err = EncodeBig5(f)
		} else if encoding == "utf-8" {
			dat, err = ioutil.ReadFile(os.Args[1])
			check(err)
		} else {
			check(fmt.Errorf("encoding: %s no supported", encoding))
		}
	}
	inputHTML := string(dat)
	//ioutil.WriteFile("/Users/masato/Dev/go/src/github.com/masato25/html2text/oo", dat, 0644)
	text, err := html2text.FromString(inputHTML, html2text.Options{PrettyTables: true})
	check(err)
	fmt.Println(text)
}
