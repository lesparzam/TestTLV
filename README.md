# Test FIF Integraciones TLV

Test para inspeccionar una cadena de caracteres la cual contiene multiples campos en el formato TLV para generar un map[string]string con los campos TLV encontrados.

## Compilación

Para compilar el proyecto es necesario tener instalado **go1.13.6**, una vez realizado la instalación  acceder a la carpeta **src** del proyecto a traves de linea de comando o desde un IDE, para luego ejecutar el comando **go run .\main.go** dónde se ejecutará el método **main()** que inspeccionará un TLV de ejemplo.

## Utilización

El proyecto contiene dos archivos en la carpeta **src**:

* **main.go** : para inspeccionar un TLV es necesario reemplazar en este archivo la variable de entrada del método **LeerTlv(tlv []byte)** en el método **main()** por el tlv a inspeccionar.

* **main_test.go** : contiene métodos para realizar algunas pruebas de la funciones contenidas en el archivo **main.go** 

## CI
https://circleci.com/gh/lesparzam/TestTLV
