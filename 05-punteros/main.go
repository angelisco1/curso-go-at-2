package main

import "fmt"

func Shift(datos []string) ([]string, error) {
	if len(datos) == 0 {
		return nil, fmt.Errorf("no se puede aplicar esta operación a un slice vacío")
	}
	return datos[1:], nil
}

func ShiftPtr(datos *[]string) (string, error) {
	if len(*datos) == 0 {
		return "", fmt.Errorf("no se puede aplicar esta operación a un slice vacío")
	}

	itemEliminado := (*datos)[0]
	*datos = (*datos)[1:]

	return itemEliminado, nil
}

func main() {

	var a int = 10
	fmt.Println(a)
	fmt.Println(&a)

	var ptr *int = &a
	fmt.Println(ptr)
	fmt.Println(*ptr)

	nombre := "Charly"
	nombre2 := nombre
	nombre2 = "Mike"

	fmt.Println(nombre)
	fmt.Println(nombre2)

	var nombre3 *string = &nombre
	*nombre3 = "Jason"
	fmt.Println(nombre)
	fmt.Println(*nombre3)

	var nombres []string = []string{nombre, nombre2, "Charly"}

	nombresSinElPrimero, _ := Shift(nombres)
	fmt.Println(nombresSinElPrimero)
	fmt.Println(nombres)

	nombreEliminado, err := ShiftPtr(&nombres)
	if err != nil {
		fmt.Println("Hay un error")
	}
	fmt.Println(nombres)
	fmt.Println(nombreEliminado)

}
