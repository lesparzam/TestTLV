package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type ResultTLV struct {
	TLV   map[string]string
	Error string
}

func main() {

	tlvStr := "11A05AB398765UJ102N2300"
	fmt.Println("TLV ingresado string: ", tlvStr)

	tlvByte := []byte(tlvStr)
	fmt.Println("TLV ingresado byte: ", tlvByte)

	tlv := []byte{49, 49, 65, 48, 53, 65, 66, 51, 57, 56, 55, 54, 53, 85, 74, 49, 48, 50, 78, 50, 51, 48, 48}

	result := LeerTlv(tlv)
	fmt.Println("TLV : ", result)
	//fmt.Println("TLV : ", reflect.ValueOf(result[0].TLV))
}

func LeerTlv(tlv []byte) []ResultTLV {

	resultTlv := []ResultTLV{}

	if len(tlv) < 6 || len(tlv) == 0 {
		result := ErrorTLV("El tlv no cumple con la logitud mínima")
		resultTlv = append(resultTlv, result)
		fmt.Println("El tlv no cumple con la logitud mínima", len(tlv))
		return resultTlv
	}

	for i := 0; i < len(tlv); i++ {

		if i < len(tlv) {

			largoTlv := string(tlv[i : i+2])
			if !EsValidoLargoTLV(largoTlv) {
				result := ErrorTLV("Error en el largo del TLV")
				resultTlv = append(resultTlv, result)
				fmt.Println("Error en el largo del TLV 2")
				break
			}

			// largoTlvExtraer: corresponde a la suma largoTlv + 5 indice correspondiente a los valores:
			// largo, tipo de dato y numero de campo
			largoTlvExtraer := i + 5 + StringToInt(largoTlv)
			if len(tlv) >= largoTlvExtraer {

				tipoDatoTlv := string(tlv[i+2 : i+3])
				if !EsValidoTipoDatoTLV(tipoDatoTlv) {
					result := ErrorTLV("No es valido el Tipo dato TLV")
					resultTlv = append(resultTlv, result)
					fmt.Println("No es valido el Tipo dato TLV: ", tipoDatoTlv)
					break
				}

				numeroCampoTlv := string(tlv[i+3 : i+5])
				if !EsValidoTipoCampoTLV(numeroCampoTlv) {
					result := ErrorTLV("El numero de campo de tipo del TLV no es un número")
					resultTlv = append(resultTlv, result)
					fmt.Println("El numero de campo de tipo del TLV no es un número: ", numeroCampoTlv)
					break
				}

				valorTlv := string(tlv[i+5 : i+5+StringToInt(largoTlv)])
				if tipoDatoTlv == "A" {
					if !EsValidoValorAlfanumericoTLV(valorTlv) {
						result := ErrorTLV("El valor del TLV no es alfanumérico")
						resultTlv = append(resultTlv, result)
						fmt.Println("El valor del TLV no es alfanumérico: ", valorTlv)
						break
					}
				}

				if tipoDatoTlv == "N" {
					if !EsValidoValorNumericoTLV(valorTlv) {
						result := ErrorTLV("El valor del TLV no es numérico")
						resultTlv = append(resultTlv, result)
						fmt.Println("El valor del TLV no es numérico: ", valorTlv)
						break
					}
				}

				//siguiente tlv
				i += StringToInt(largoTlv) + 4

				tlvOut := map[string]string{"largo": largoTlv, "tipoDato": tipoDatoTlv, "numeroCampo": numeroCampoTlv, "valor": valorTlv}
				result := ResultTLV{tlvOut, "Sin error"}
				resultTlv = append(resultTlv, result)

			} else {

				result := ErrorTLV("Estructura del TLV no válida.")
				resultTlv = append(resultTlv, result)
				fmt.Println("Estructura del TLV no válida.", resultTlv)
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
	if len(largoTlv) < 2 {
		fmt.Println("El Largo del TLV es menor a 2: ", len(largoTlv))
		esValido = false
	}
	if !EsNumero(largoTlv) {
		fmt.Println("El Largo del TLV no es un número: ", largoTlv)
		esValido = false
	}
	return esValido
}

func EsValidoTipoDatoTLV(tipoDatoTlv string) bool {
	esValido := true
	if !EsValidoTipoDato(tipoDatoTlv) {
		fmt.Println("No es valido el Tipo dato TLV:", tipoDatoTlv)
		esValido = false
	}
	return esValido
}

func EsValidoTipoCampoTLV(numeroCampoTlv string) bool {
	esValido := true
	if !EsNumero(numeroCampoTlv) {
		fmt.Println("El numero de campo de tipo del TLV no es un número: ", numeroCampoTlv)
		esValido = false
	}
	return esValido
}

func EsValidoValorAlfanumericoTLV(valorTlv string) bool {
	esValido := true
	if !EsAlfanumerico(valorTlv) {
		fmt.Println("El valor del TLV no es alfanumérico: ", valorTlv)
		esValido = false
	}
	return esValido
}

func EsValidoValorNumericoTLV(valorTlv string) bool {
	esValido := true
	if !EsNumero(valorTlv) {
		fmt.Println("El valor del TLV no es numérico: ", valorTlv)
		esValido = false
	}
	return esValido
}

func Salir() {
	os.Exit(0)
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
		fmt.Println(err)
		Salir()
	}
	return i
}
