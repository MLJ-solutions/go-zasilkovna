package main

import (
	"fmt"
	"go-zasilkovna/models"
	"os"
)

func main() {
	client, _ := New(&Options{
		Creds: NewCredentials(os.Getenv("ZASILKOVNA_KEY")),
	})
	a, err := client.PacketAttributesValid(models.PacketAttributes{
		Number:    "123abc",
		Name:      "John",
		Surname:   "Doe",
		Email:     "john.doe@test.te",
		Phone:     "123321123",
		AddressId: 95,
		Value:     100.00,
		Eshop:     "my.eshop",
	})
	fmt.Println(a)
	fmt.Println(err)
}
