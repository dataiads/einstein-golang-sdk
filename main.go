package main

import (
	"fmt"
	"log"

	"github.com/Dataiads/einstein-golang-sdk/client/operations"

	apiclient "github.com/myproject/client"
)

func main() {

	// make the request to get all items
	resp, err := apiclient.Default.Operations.All(operations.AllParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}
