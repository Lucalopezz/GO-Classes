package main

import "fmt"

func main() {
	// par de chave e valor

	myMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	mapAny := map[string]any{ // JSON
		"key1": "value1",
		"key2": 2,
	}
	fmt.Println(myMap)
	fmt.Println(mapAny)

	/*
		+-----------+-------------------------+
		| Balde 0   | [("a", 1), ("e", 5)]    |
		+-----------+-------------------------+
		| Balde 1   | [("b", 2)]              |
		+-----------+-------------------------+
		| Balde 2   | [("c", 3), ("f", 6)]    |
		+-----------+-------------------------+
		| Balde 3   | [("d", 4)]              |
		+-----------+-------------------------+
	*/

	// Essa impressão pode sair em outra ordem porque maps são coleções sem ordem definida em Go.
	for key, value := range myMap {
		fmt.Println("Chave:", key, "Valor:", value)
	}

	// var myMap2 map[string]any // map nil
	// myMap2["key1"] = "value1" // panic: atribuição a uma entrada em map nil
	myMap2 := map[string]any{} // map vazio
	myMap2["key1"] = "value1"  // funciona normalmente
	// com make
	myMap3 := make(map[string]any)
	myMap3["key1"] = "value1" // funciona normalmente

	fmt.Println(myMap2["dontExist"]) // imprime <nil>
	// pode causar panic se você tentar acessar uma chave inexistente e depois usá-la como valor

	myValue, ok := myMap2["dontExist"]
	fmt.Println(myValue, ok) // imprime <nil> false

	myValue2, ok2 := myMap2["key99"].(int) // asserção de tipo; causará panic se a chave não existir ou o valor não for do tipo int
	if !ok2 {
		fmt.Println("A chave 'key99' não existe ou não é do tipo int")
	} else {
		fmt.Println("Valor de 'key99':", myValue2)
	}
}
