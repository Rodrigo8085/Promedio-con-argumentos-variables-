package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func promedio(numeros ...float64) float64 {
	var sum float64 = 0
	for _, num := range numeros {
		sum += num
	}
	return sum / float64(len(numeros))
}

func main() {
	argumentos := os.Args[1:] // empieza en 1  po que el cero es nombre de la variable
	var numeros []float64
	for _, argumento := range argumentos {
		numero, err := strconv.ParseFloat(argumento, 64)
		if err != nil {
			log.Fatal(err)
		}
		numeros = append(numeros, numero)
	}
	fmt.Printf("El promedio es: %0.2f\n", promedio(numeros...))
}
