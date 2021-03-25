package main

import (
	"context"
	Spec "oapi-codegen-testbed/pkg/generated"

	log "github.com/sirupsen/logrus"
)

func String(v string) *string {
	return &v
}

func main() {

	client, err := Spec.NewClient("http://localhost:8080", []Spec.ClientOption{}...)
	if err != nil {
		log.Fatal(err)
	}

	requestBody := Spec.RequestObject{
		FooObject: Spec.FooObject{
			Foo: String("fooValue"),
			// ObjType: "FooObject",
		},
		BarObject: Spec.BarObject{
			Bar: String("barValue"),
			// ObjType: "BarOject",
		},
	}

	_, err = client.Get(context.Background(), Spec.GetJSONRequestBody(requestBody))
	if err != nil {
		log.Fatal(err)
	}
}
