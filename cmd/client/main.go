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

	requestBody := Spec.GetJSONRequestBody{
		BaseObject: Spec.BaseObject{
			Foo:     String("fooValue"),
			ObjType: "BarOject",
		},
		ChildObject: Spec.ChildObject{
			Bar: String("barValue"),
		},
	}

	_, err = client.Get(context.Background(), requestBody)
	if err != nil {
		log.Fatal(err)
	}
}
