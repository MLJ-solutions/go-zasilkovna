package models

type ItemCollection struct {
	Item []Item `json:"item"`
}

type Item struct { // Item represent one kind of thing in the packet. Is an array of type Attribute.
	Attribute Attribute `json:"attribute"`
}
