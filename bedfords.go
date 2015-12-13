package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Root struct {
	XMLName xml.Name `xml:"Root"`
	Data
}

type Data struct {
	XMLName xml.Name `xml:"data"`
	Records []Record
}

type Record struct {
	XMLName xml.Name `xml:"record"`
	Fields  []Field
}

type Field struct {
	XMLName xml.Name `xml:"field"`
	Value   int      `xml:"value"`
	Name    string   `xml:"name,attr"`
}

func main() {
	xmlFile, err := os.Open("data.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	bytes, _ := ioutil.ReadAll(xmlFile)

	var q Root
	xml.Unmarshal(bytes, &q)

	fmt.Println(q)

	field := &Field{Value: 2, Name: "hello"}
	fields := []Field{*field}
	record := &Record{Fields: fields}
	records := []Record{*record}
	data := &Data{Records: records}
	root := &Root{Data: *data}

	output, err := xml.MarshalIndent(root, "", "    ")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(string(output))
}
