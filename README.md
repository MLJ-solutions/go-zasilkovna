# go-zasilkovna

## Connecting to API

```go
package main

import (
	"log"
	"go-zasilkovna/models"
)

func main() {
	client, _ := New(&Options{
		Creds: NewCredentials("API KEY"),
	})

	log.Println(client)
}
```

## Full example of getting info about packet

```go
package main

import (
	"log"
	"go-zasilkovna/models"
)

func main() {
	client, _ := New(&Options{
		Creds: NewCredentials("API KEY"),
	})

	log.Println(client)

	a, err := client.PacketInfo(1234567890) //packetId
	log.Println(a)
	log.Println(err)
}
```

## Full example of validating packet attributes

```go
package main

import (
	"log"
	"go-zasilkovna/models"
)

func main() {
	client, _ := New(&Options{
		Creds: NewCredentials("API KEY"),
	})

	log.Println(client)

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
	log.Println(a)
	log.Println(err)
}
```

## Full example of creating new packet

```go
package main

import (
	"log"
	"go-zasilkovna/models"
)

func main() {
	client, _ := New(&Options{
		Creds: NewCredentials("API KEY"),
	})

	log.Println(client)

	a, err := client.CreatePacket(models.PacketAttributes{
		Number:    "123abc",
		Name:      "John",
		Surname:   "Doe",
		Email:     "john.doe@test.te",
		Phone:     "123321123",
		AddressId: 95,
		Value:     100.00,
		Eshop:     "my.eshop",
	})
	log.Println(a)
	log.Println(err)
}
```