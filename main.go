package main

import (
	"bytes"
	"fmt"
	"vandal/pdf"
	"vandal/toukibo"
)

func main() {
	content, err := readPdf("sample/houjin/sample1.pdf")
	if err != nil {
		panic(err)
	}
	houjin, err := toukibo.Extract(content)
	if err != nil {
		panic(err)
	}
	fmt.Println(houjin.CreatedAt)
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
