package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}*/

func TestLeerTlv(t *testing.T) {

	fmt.Println("----- INICIO TestLeerTlv")

	resultTlv := []ResultTLV{}
	tlvIngresarByte := []byte{49, 49, 65, 48, 53, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 51, 48, 48} //11A05AB398765UJ102N2300

	tlvIngresarMap1 := map[string]string{"largo": "11", "numeroCampo": "05", "tipoDato": "A", "valor": "AB398765UJ1"}
	resultTlv = append(resultTlv, ResultTLV{tlvIngresarMap1, "Sin error"})

	tlvIngresarMap2 := map[string]string{"largo": "02", "numeroCampo": "23", "tipoDato": "N", "valor": "00"}
	resultTlv = append(resultTlv, ResultTLV{tlvIngresarMap2, "Sin error"})

	for _, c := range []struct {
		ingresa []byte
		espera  []ResultTLV
	}{
		{tlvIngresarByte, resultTlv},
	} {
		result := LeerTlv(c.ingresa)

		for i, _ := range result {

			fmt.Println("TLV RESULT: ", result[i].TLV)
			fmt.Println("TLV ESPERA: ", c.espera[i].TLV)

			if !reflect.DeepEqual(result[i].TLV, c.espera[i].TLV) {
				t.Errorf("LeerTlv(%v) == %v, espera %v", c.ingresa, result, c.espera)
			}
		}
	}

	fmt.Println("----- FIN TestLeerTlv")
}

func TestLeerTlv2(t *testing.T) {

	fmt.Println("----- INICIO TestLeerTlv2")

	tlvEspera := []ResultTLV{}
	tlvIngresarByte := []byte{49, 49, 65, 48, 53, 65} //11A05A

	tlvIngresarMap := map[string]string{"largo": "", "numeroCampo": "", "tipoDato": "", "valor": ""}
	tlvEspera = append(tlvEspera, ResultTLV{tlvIngresarMap, "Estructura del TLV no válida."})

	for _, c := range []struct {
		ingresa []byte
		espera  []ResultTLV
	}{
		{tlvIngresarByte, tlvEspera},
	} {

		result := LeerTlv(c.ingresa)

		fmt.Println("TLV RESULT: ", result[0].TLV)
		fmt.Println("TLV ESPERA: ", c.espera[0].TLV)

		if !reflect.DeepEqual(result[0].TLV, c.espera[0].TLV) {
			t.Errorf("LeerTlv(%v) == %v, espera %v", c.ingresa, result, c.espera)
		}
	}

	fmt.Println("----- FIN TestLeerTlv2")
}

func TestLeerTlv3(t *testing.T) {

	fmt.Println("----- INICIO TestLeerTlv3")

	tlvEspera := []ResultTLV{}
	tlvIngresarByte := []byte{49, 49, 65, 48, 53} //11A05

	tlvIngresarMap := map[string]string{"largo": "", "numeroCampo": "", "tipoDato": "", "valor": ""}
	tlvEspera = append(tlvEspera, ResultTLV{tlvIngresarMap, "El tlv no cumple con la logitud mínima"})

	for _, c := range []struct {
		ingresa []byte
		espera  []ResultTLV
	}{
		{tlvIngresarByte, tlvEspera},
	} {

		result := LeerTlv(c.ingresa)

		fmt.Println("TLV RESULT: ", result[0].TLV)
		fmt.Println("TLV ESPERA: ", c.espera[0].TLV)

		if !reflect.DeepEqual(result[0].TLV, c.espera[0].TLV) {
			t.Errorf("LeerTlv(%v) == %v, espera %v", c.ingresa, result, c.espera)
		}
	}

	fmt.Println("----- FIN TestLeerTlv3")
}

func TestEsAlfanumerico(t *testing.T) {

	fmt.Println("----- INICIO TestEsAlfanumerico")

	for _, c := range []struct {
		ingresa string
		espera  bool
	}{
		{"123456789", false},
		{"1Q2W3E4R", true},
		{"1-", false},
	} {
		result := EsAlfanumerico(c.ingresa)

		fmt.Println("INGRESA: ", c.ingresa)
		fmt.Println("RESULT: ", result)
		fmt.Println("ESPERA: ", c.espera)

		if result != c.espera {
			t.Errorf("EsAlfanumerico(%q) == %t, espera %t", c.ingresa, result, c.espera)
		}
	}

	fmt.Println("----- FIN TestEsAlfanumerico")
}
