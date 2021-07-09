package main

import (
	"fmt"
	"go-zasilkovna/models"
)

func main() {
	client, _ := New(&Options{
		Creds: NewCredentials("c4cb0ed91697f5c8927f926df041c75c"),
	})
	//a, err := client.PacketInfo(2811053071)
	//a, err := client.PacketStatus(2811053071)
	//a, err := client.PacketTracking(2811053071)
	//a, err := client.PacketGetStoredUntil(2811053071)
	//a, err := client.BarcodePng("Z2811053071")
	//a, err := client.PacketLabelPdf(models.Ids{PacketIds: []int{2811053071}}, "A7 on A7", 0)
	a, err := client.PacketAttributesValid(models.PacketAttributes{
		Number:              "123abc",
		Name:                "John",
		Surname:             "Doe",
		Email:               "john.doe@test.te",
		Phone:               "123321123",
		AddressId:           95,
		Value:               100.00,
		Eshop:               "my.eshop",
	})
	/*a, err := client.ClaimAttributesValid(models.ClaimAttributes{
		Number:   "123abc",
		Email:    "john.doe@test.te",
		Phone:    "123321123",
		Value:    100.00,
		Eshop:    "my.eshop",
		Currency: "CZK",
	})*/
	fmt.Println(a)
	fmt.Println(err)

	//PacketAttributes test
	/*isValidated, err := models.ValidatePacketAttributes()
	fmt.Println(isValidated)
	fmt.Println(err)*/
}
