package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFileLine(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	strs := []string{}
	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		str := string(line)
		strs = append(strs, str)
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()
	return strs
}

func WriteFileLine(path string, str string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if f == nil {
		split := strings.Split(path, "/")
		filePath := strings.Join(split[:len(split)-1], "/")
		err = os.MkdirAll(filePath, os.ModePerm)
		f, err = os.Create(path)
	}
	if err != nil {
		panic(err)
	}
	r := bufio.NewWriter(f)
	str = str + "\r\n"
	//str = "mc cao"
	_, err = r.WriteString(str)
	//if err == io.EOF {
	//
	//}
	if err != nil {
		panic(err)
	}
	r.Flush()
	defer f.Close()
}
