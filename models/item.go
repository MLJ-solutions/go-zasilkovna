package models

type ItemCollection struct {
	Item Item `xml:"item"`
}

type Item struct { // Item represent one kind of thing in the packet. Is an array of type Attribute.
	Attribute Attribute `xml:"attribute"`
}
