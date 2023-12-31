package test

import (
	"fmt"
	"reflect"
	"testing"
	"github.com/philipelima/brasilapi-go/cep"
)


func TestV1(t *testing.T) {

	cepNumber := "02045-970"
	response, err := cep.V1().Get(cepNumber)

	if err != nil {
		t.Error(err)
	}
	
	if isAValidCepStruct(response) == false {
		t.Fatal("response is not a CepV1Response struct", response)
	}

	fmt.Println(response.Cep)

}


func TestV2(t *testing.T) {

	cepNumber := "02045-970"
	response, err := cep.V2().Get(cepNumber)

	if err != nil {
		t.Error(err)
	}
	
	if isAValidCepStruct(response) == false {
		t.Fatal("response is not a CepV1Response struct", response)
	}

	fmt.Println(response.Cep)

}

func isAValidCepStruct(i interface{}) bool {
	switch i.(type) {
		case *cep.CepV1Response:
			return true

		case *cep.CepV2Response:
			return true
	
		default:
			fmt.Println(reflect.TypeOf(i))
			return false
	}
}