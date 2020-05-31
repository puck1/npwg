/* ASN.1
 */
package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)


type IDCard struct {
	Name 		string
	Age			int
	IDNumber	string
}

func (this IDCard) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, IDNumber: %s", this.Name, this.Age, this.IDNumber)
}

func main() {
	myIDCard := IDCard{
		Name:     "Puck",
		Age:      18,
		IDNumber: "123456",
	}

	// The fields of a structure must all be exportable, that is,
	// field names must begin with an uppercase letter.
	// Go uses the reflect package to marshall/unmarshall structures,
	// so it must be able to examine all fields.
	mdata, err := asn1.Marshal(myIDCard)
	checkError(err)

	// Unmarshal should receive a pointer to object.
	var otherIDCard IDCard
	_, err1 := asn1.Unmarshal(mdata, &otherIDCard)
	checkError(err1)
	fmt.Printf("After marshal/unmarshal: %s", myIDCard)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
