package main

import (
	"bytes"
	"flag"
	"fmt"
	"vandal/pdf"
	"vandal/toukibo"
)

func main() {
	f := flag.String("path", "sample1", "")
	flag.Parse()
	path := fmt.Sprintf("sample/houjin/%s.pdf", *f)
	content, err := readPdf(path)

	if err != nil {
		panic(err)
	}
	//_, err = toukibo.Extract(content)
	tc, err := toukibo.Parse(content)
	if err != nil {
		panic(err)
	}
	houjin := toukibo.NewHoujinFromToukibo(tc)
	houjin.Extract()

	fmt.Println(houjin.String())
	return
}

func readPdf(path string) (string, error) {
	r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
