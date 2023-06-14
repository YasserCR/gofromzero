package variables

import "fmt"

/*Para que la funcion pueda ser importada desde fuera del paquete debe
de estar escrita en uppercase*/
func ShowIntegers() {
	var commonInt int
	intOf32 := int32(10)
	intOf64 := int64(100)

	//Cuando se define una variable toma el valor más bajo, no null
	fmt.Println("Common int = ", commonInt)
	fmt.Println("Int 32 = ", intOf32)
	fmt.Println("Int 64 = ", intOf64)
}
