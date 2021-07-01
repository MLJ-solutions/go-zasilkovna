package models

type ItemCollection struct {
	Item []Item `xml:"item"`
}

func NewItemCollection(Item []Item) *ItemCollection {
	return &ItemCollection{
		Item: Item,
	}
}

type Item struct { // Item represent one kind of thing in the packet. Is an array of type Attribute.
	Attribute Attribute `xml:"attribute"`
}

func NewItem(Attribute Attribute) *Item {
	return &Item{
		Attribute: Attribute,
	}
}
