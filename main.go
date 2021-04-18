package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"xml_parsing_with_dto/dto"
)

func main() {
	file, err := os.Open("./books.xml")
	defer file.Close()

	if err !=nil{
		panic(err)
	}
	data, _ := ioutil.ReadAll(file)
	var catalog dto.Catalog
	xml.Unmarshal(data, &catalog)

	for i:=0; i < len(catalog.Books); i++{
		book := catalog.Books[i]
		fmt.Println(book.GetTitle())
	}
}
