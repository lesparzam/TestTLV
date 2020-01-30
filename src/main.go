package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type ResultTLV struct {
	TLV   map[string]string
	Error string
}

const (
	MensajeErrorLargoTLV                     = "Error en el largo del TLV"
	MensajeErrorTLVNoCumpleLongitudMinima    = "El tlv no cumple con la logitud mínima"
	MensajeErrorNoEsValidoTipoDatoTLV        = "No es válido el Tipo dato TLV"
	MensajeErrorNumeroCampoTipoTLVNoEsNumero = "El número de campo de tipo del TLV no es un número"
	MensajeErrorElValorTLVNoEsAlfanumerico   = "El valor del TLV no es alfanumérico"
	MensajeErrorElValorTLVNoEsNumerico       = "El valor del TLV no es númerico"
	MensajeErrorEstructuraTLVNoValida        = "Estructura del TLV no válida"
	MensajeErrorOk                           = "Sin error"
)

func main() {

	fmt.Println("TLV Out: ", LeerTlv([]byte{49, 49, 65, 48, 53, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 51, 48, 48}))

}

func LeerTlv(tlv []byte) []ResultTLV {

	resultTlv := []ResultTLV{}

	if len(tlv) < 6 || len(tlv) == 0 {
		resultTlv = append(resultTlv, ErrorTLV(MensajeErrorTLVNoCumpleLongitudMinima))
		fmt.Println(MensajeErrorTLVNoCumpleLongitudMinima, ": ", len(tlv))
		return resultTlv
	}

	for i := 0; i < len(tlv); i++ {

		if i < len(tlv) {

			largoTlv := string(tlv[i : i+2])
			if !EsValidoLargoTLV(largoTlv) {
				resultTlv = append(resultTlv, ErrorTLV(MensajeErrorLargoTLV))
				break
			}

			// largoTlvExtraer: corresponde a la suma largoTlv + 5 indice correspondiente a los valores:
			// largo, tipo de dato y numero de campo
			largoTlvExtraer := i + 5 + StringToInt(largoTlv)
			if len(tlv) >= largoTlvExtraer {

				tipoDatoTlv := string(tlv[i+2 : i+3])
				if !EsValidoTipoDatoTLV(tipoDatoTlv) {
					resultTlv = append(resultTlv, ErrorTLV(MensajeErrorNoEsValidoTipoDatoTLV))
					fmt.Println(MensajeErrorNoEsValidoTipoDatoTLV, ": ", tipoDatoTlv)
					break
				}

				numeroCampoTlv := string(tlv[i+3 : i+5])
				if !EsValidoTipoCampoTLV(numeroCampoTlv) {
					resultTlv = append(resultTlv, ErrorTLV(MensajeErrorNumeroCampoTipoTLVNoEsNumero))
					fmt.Println(MensajeErrorNumeroCampoTipoTLVNoEsNumero, ": ", numeroCampoTlv)
					break
				}

				valorTlv := string(tlv[i+5 : i+5+StringToInt(largoTlv)])
				if tipoDatoTlv == "A" {
					if !EsValidoValorAlfanumericoTLV(valorTlv) {
						resultTlv = append(resultTlv, ErrorTLV(MensajeErrorElValorTLVNoEsAlfanumerico))
						fmt.Println(MensajeErrorElValorTLVNoEsAlfanumerico, ": ", valorTlv)
						break
					}
				}

				if tipoDatoTlv == "N" {
					if !EsValidoValorNumericoTLV(valorTlv) {
						resultTlv = append(resultTlv, ErrorTLV(MensajeErrorElValorTLVNoEsNumerico))
						fmt.Println(MensajeErrorElValorTLVNoEsNumerico, ": ", valorTlv)
						break
					}
				}

				//siguiente tlv
				i += StringToInt(largoTlv) + 4

				tlvOut := map[string]string{"largo": largoTlv, "tipoDato": tipoDatoTlv, "numeroCampo": numeroCampoTlv, "valor": valorTlv}
				resultTlv = append(resultTlv, ResultTLV{tlvOut, MensajeErrorOk})

			} else {

				resultTlv = append(resultTlv, ErrorTLV(MensajeErrorEstructuraTLVNoValida))
				fmt.Println(MensajeErrorEstructuraTLVNoValida, resultTlv)
				break
			}
		}
	}
	return resultTlv
}

func ErrorTLV(err string) ResultTLV {
	tlvMap := map[string]string{"largo": "", "tipoDato": "", "numeroCampo": "", "valor": ""}
	return ResultTLV{tlvMap, err}
}

func EsValidoLargoTLV(largoTlv string) bool {
	esValido := true
	if !EsNumero(largoTlv) {
		esValido = false
	}
	return esValido
}

func EsValidoTipoDatoTLV(tipoDatoTlv string) bool {
	esValido := true
	if !EsValidoTipoDato(tipoDatoTlv) {
		esValido = false
	}
	return esValido
}

func EsValidoTipoCampoTLV(numeroCampoTlv string) bool {
	esValido := true
	if !EsNumero(numeroCampoTlv) {
		esValido = false
	}
	return esValido
}

func EsValidoValorAlfanumericoTLV(valorTlv string) bool {
	esValido := true
	if !EsAlfanumerico(valorTlv) {
		esValido = false
	}
	return esValido
}

func EsValidoValorNumericoTLV(valorTlv string) bool {
	esValido := true
	if !EsNumero(valorTlv) {
		esValido = false
	}
	return esValido
}

func EsNumero(s string) bool {
	re := regexp.MustCompile(`^\d*$`)
	return re.MatchString(s)
}

func EsValidoTipoDato(s string) bool {
	//A: Alfanumérico y N: Numérico
	re := regexp.MustCompile(`[AN]`)
	return re.MatchString(s)
}

func EsAlfanumerico(s string) bool {
	re := regexp.MustCompile(`^([a-zA-Z_]{1,}\d{1,})+|(\d{1,}[a-zA-Z_]{1,})+$`)
	return re.MatchString(s)
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}
