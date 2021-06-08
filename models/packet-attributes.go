package models

type PacketAttributes struct {
	Id                  int                   `json:"id,omitempty"`
	Number              string                `json:"number"`
	Name                string                `json:"name"`
	Surname             string                `json:"surname"`
	Company             string                `json:"company,omitempty"`
	Email               string                `json:"email"`
	Phone               string                `json:"phone"`
	AddressId           int                   `json:"address_id"`
	Currency            string                `json:"currency,omitempty"`
	Cod                 float32               `json:"cod,omitempty"`
	Value               float32               `json:"value"`
	Weight              float32               `json:"weight"`
	DeliverOn           ZasilkovnaDate        `json:"deliver_on,omitempty"`
	Eshop               string                `json:"eshop"`
	AdultContent        bool                  `json:"adult_content,omitempty"`
	Note                string                `json:"note,omitempty"`
	Street              string                `json:"street"`
	HouseNumber         string                `json:"house_number"`
	City                string                `json:"city"`
	Province            string                `json:"province,omitempty"`
	Zip                 string                `json:"zip"`
	CarrierService      string                `json:"carrier_service,omitempty"`
	CustomerBarcode     string                `json:"customer_barcode,omitempty"`
	CarrierPickupPoint  string                `json:"carrier_pickup_point"`
	CustomsDeclaration  Size                  `json:"customs_declaration"`
	Size                Size                  `json:"size"`
	AttributeCollection []AttributeCollection `json:"attribute_collection"`
	Items               []ItemCollection      `json:"items"`
}
