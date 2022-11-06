package aes3k

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAll(t *testing.T) {
	//Create DummyStruct struct
	type DummyStruct struct {
		Name    string `json:"name"`
		Numbers []int  `json:"numbers"`
	}

	//Initialize a new DummyStruct
	d := DummyStruct{
		Name:    "Bob",
		Numbers: []int{1, 2, 3, 4, 5},
	}

	log.Printf("As Struct -> %v", d)

	//Marshal DummyStruct into []byte
	marshalledStruct, err := json.Marshal(&d)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	log.Printf("As marshalled []byte -> %v", marshalledStruct)

	//Encrypt the marshalledStruct with a key
	encryptedData := Encrypt(marshalledStruct, "1234567890123456")
	log.Printf("As encrypted []byte -> %s", encryptedData)

	//Decrypt the encrypted []byte
	decryptedData := Decrypt(encryptedData, "1234567890123456")
	log.Printf("As decrypted []byte -> %s", decryptedData)

	//Unmarshal the []byte to a type
	unmarshalledStruct := &DummyStruct{}
	json.Unmarshal(decryptedData, &unmarshalledStruct)
	log.Printf("As unmarshalled type -> %v", unmarshalledStruct)

}
