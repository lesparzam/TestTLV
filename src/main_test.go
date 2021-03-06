package main

import (
	"fmt"
	"reflect"
	"testing"
)

var tlvIngresarMap = map[string]string{"largo": "", "numeroCampo": "", "tipoDato": "", "valor": ""}

func TestLeerTlvCorrecto(t *testing.T) {

	fmt.Println("----- TestLeerTlvCorrecto")

	resultTlv := []ResultTLV{}
	tlvIngresarByte := []byte{49, 49, 65, 48, 53, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 51, 48, 48} //11A05AB398765UJ102N2300

	tlvIngresarMap1 := map[string]string{"largo": "11", "numeroCampo": "05", "tipoDato": "A", "valor": "AB398765UJ1"}
	resultTlv = append(resultTlv, ResultTLV{tlvIngresarMap1, MensajeErrorOk})

	tlvIngresarMap2 := map[string]string{"largo": "02", "numeroCampo": "23", "tipoDato": "N", "valor": "00"}
	resultTlv = append(resultTlv, ResultTLV{tlvIngresarMap2, MensajeErrorOk})

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
}

func TestLeerTlvEstructuraNoEsValida(t *testing.T) {

	fmt.Println("----- TestLeerTlvEstructuraNoEsValida")

	tlvEspera := []ResultTLV{}
	tlvIngresarByte := []byte{49, 49, 65, 48, 53, 65} //11A05A
	tlvEspera = append(tlvEspera, ResultTLV{tlvIngresarMap, MensajeErrorEstructuraTLVNoValida})

	for _, c := range []struct {
		ingresa []byte
		espera  []ResultTLV
	}{
		{tlvIngresarByte, tlvEspera},
	} {
		result := LeerTlv(c.ingresa)
		if !reflect.DeepEqual(result[0].TLV, c.espera[0].TLV) {
			t.Errorf("LeerTlv(%v) == %v, espera %v", c.ingresa, result, c.espera)
		}
	}
}

func TestLeerTlvLongitudNoEsValida(t *testing.T) {

	fmt.Println("----- TestLeerTlvLongitudNoEsValida")

	tlvEspera := []ResultTLV{}
	tlvIngresarByte := []byte{49, 49, 65, 48, 53} //11A05
	tlvEspera = append(tlvEspera, ResultTLV{tlvIngresarMap, MensajeErrorTLVNoCumpleLongitudMinima})

	for _, c := range []struct {
		ingresa []byte
		espera  []ResultTLV
	}{
		{tlvIngresarByte, tlvEspera},
	} {
		result := LeerTlv(c.ingresa)
		if !reflect.DeepEqual(result[0].TLV, c.espera[0].TLV) {
			t.Errorf("LeerTlv(%v) == %v, espera %v", c.ingresa, result, c.espera)
		}
	}
}

func TestLargoTLVNoEsValido(t *testing.T) {

	fmt.Println("----- TestLargoTLVNoEsValido")

	tlvEspera := []ResultTLV{}
	tlvIngresarByte := []byte{49, 88, 65, 48, 53, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 51, 48, 48} //1XA05AB398765UJ102N2300
	tlvEspera = append(tlvEspera, ResultTLV{tlvIngresarMap, MensajeErrorLargoTLV})

	for _, c := range []struct {
		ingresa []byte
		espera  []ResultTLV
	}{
		{tlvIngresarByte, tlvEspera},
	} {
		result := LeerTlv(c.ingresa)
		if !reflect.DeepEqual(result[0].TLV, c.espera[0].TLV) {
			t.Errorf("LeerTlv(%v) == %v, espera %v", c.ingresa, result, c.espera)
		}
	}
}

func TestTipoDatoTLVNoEsValido(t *testing.T) {

	fmt.Println("----- TestTipoDatoTLVNoEsValido")

	tlvEspera := []ResultTLV{}
	tlvIngresarByte := []byte{49, 49, 88, 48, 53, 88, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 51, 48, 48} //11X05AB398765UJ102N2300
	tlvEspera = append(tlvEspera, ResultTLV{tlvIngresarMap, MensajeErrorNoEsValidoTipoDatoTLV})

	for _, c := range []struct {
		ingresa []byte
		espera  []ResultTLV
	}{
		{tlvIngresarByte, tlvEspera},
	} {
		result := LeerTlv(c.ingresa)
		if !reflect.DeepEqual(result[0].TLV, c.espera[0].TLV) {
			t.Errorf("LeerTlv(%v) == %v, espera %v", c.ingresa, result, c.espera)
		}
	}
}

func TestElNumeroCampoTipoTLVNoEsNumero(t *testing.T) {

	fmt.Println("----- TestElNumeroCampoTipoTLVNoEsNumero")

	tlvEspera := []ResultTLV{}
	tlvIngresarByte := []byte{49, 49, 65, 88, 88, 88, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 51, 48, 48} //11AXXAB398765UJ102N2300
	tlvEspera = append(tlvEspera, ResultTLV{tlvIngresarMap, MensajeErrorNumeroCampoTipoTLVNoEsNumero})

	for _, c := range []struct {
		ingresa []byte
		espera  []ResultTLV
	}{
		{tlvIngresarByte, tlvEspera},
	} {
		result := LeerTlv(c.ingresa)
		if !reflect.DeepEqual(result[0].TLV, c.espera[0].TLV) {
			t.Errorf("LeerTlv(%v) == %v, espera %v", c.ingresa, result, c.espera)
		}
	}
}

func TestElValorTLVNoEsAlfanumerico(t *testing.T) {

	fmt.Println("----- TestElValorTLVNoEsAlfanumerico")

	tlvEspera := []ResultTLV{}
	tlvIngresarByte := []byte{49, 49, 65, 48, 53, 49, 50, 51, 52, 53, 54, 55, 56, 57, 49, 48, 48, 50, 78, 50, 51, 48, 48} //11A051234567891002N2300
	tlvEspera = append(tlvEspera, ResultTLV{tlvIngresarMap, MensajeErrorElValorTLVNoEsAlfanumerico})

	for _, c := range []struct {
		ingresa []byte
		espera  []ResultTLV
	}{
		{tlvIngresarByte, tlvEspera},
	} {
		result := LeerTlv(c.ingresa)
		if !reflect.DeepEqual(result[0].TLV, c.espera[0].TLV) {
			t.Errorf("LeerTlv(%v) == %v, espera %v", c.ingresa, result, c.espera)
		}
	}
}

func TestElValorTLVNoEsNumerico(t *testing.T) {

	fmt.Println("----- TestElValorTLVNoEsNumerico")

	tlvEspera := []ResultTLV{}
	tlvIngresarByte := []byte{49, 49, 78, 48, 53, 65, 83, 68, 70, 71, 72, 74, 75, 76, 79, 73, 48, 50, 78, 50, 51, 48, 48} //11A051234567891002N2300
	tlvEspera = append(tlvEspera, ResultTLV{tlvIngresarMap, MensajeErrorElValorTLVNoEsNumerico})

	for _, c := range []struct {
		ingresa []byte
		espera  []ResultTLV
	}{
		{tlvIngresarByte, tlvEspera},
	} {
		result := LeerTlv(c.ingresa)
		if !reflect.DeepEqual(result[0].TLV, c.espera[0].TLV) {
			t.Errorf("LeerTlv(%v) == %v, espera %v", c.ingresa, result, c.espera)
		}
	}
}

func TestEsAlfanumerico(t *testing.T) {

	fmt.Println("----- TestEsAlfanumerico")

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
}

func TestStringToIntValidoNoValido(t *testing.T) {

	fmt.Println("----- TestStringToIntValidoNoValido")

	for _, c := range []struct {
		ingresa string
		espera  int
	}{
		{"123456789", 123456789},
		{"1Q2W3E4R", -1},
		{"1-", -1},
	} {
		result := StringToInt(c.ingresa)

		if result != c.espera {
			t.Errorf("EsAlfanumerico(%q) == %v, espera %v", c.ingresa, result, c.espera)
		}
	}
}
