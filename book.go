package dto

import (
	"encoding/xml"
)

type Book struct {
	Name     xml.Name  `xml:"book"`
	Id          string    `xml:"id,attr"`
	Title       string    `xml:"title"`
	Genre       string    `xml:"genre"`
	Price       float32   `xml:"price"`
	Date        string `xml:"publish_date"`
	Description string    `xml:"description"`
}

func (b *Book) GetId() string {
	return b.Id
}

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) GetGenre() string {
	return b.Genre
}

func (b *Book) GetPrice() float32 {
	return b.Price
}

func (b *Book) GetDate() string {
	return b.Date
}

func (b *Book) GetDescription() string {
	return b.Description
}
