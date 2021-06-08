package models

type AttributeCollection struct {
	Attribute Attribute `json:"attribute"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
