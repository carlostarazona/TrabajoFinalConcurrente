package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Data    []Data    `json:"data"`
	Weights []float64 `json:"weights"`
	Heights []float64 `json:"heights"`
}

type Data struct {
	Inputs []float64 `json:"inputs"`
	Output float64   `json:"output"`
}

func calc_ols_params(y []float64, x [][]float64, n_iterations int, alpha float64) []float64 {

	thetas := make([]float64, len(x))

	for i := 0; i < n_iterations; i++ {

		my_diffs := calc_diff(thetas, y, x)

		my_grad := calc_gradient(my_diffs, x)

		for j := 0; j < len(my_grad); j++ {
			thetas[j] += alpha * my_grad[j]
		}
	}
	return thetas
}

func calc_diff(thetas []float64, y []float64, x [][]float64) []float64 {
	diffs := make([]float64, len(y))
	for i := 0; i < len(y); i++ {
		prediction := 0.0
		for j := 0; j < len(thetas); j++ {
			prediction += thetas[j] * x[j][i]
		}
		diffs[i] = y[i] - prediction
	}
	return diffs
}

func calc_gradient(diffs []float64, x [][]float64) []float64 {
	gradient := make([]float64, len(x))
	for i := 0; i < len(diffs); i++ {
		for j := 0; j < len(x); j++ {
			gradient[j] += diffs[i] * x[j][i]
		}
	}
	for i := 0; i < len(x); i++ {
		gradient[i] = gradient[i] / float64(len(diffs))
	}

	return gradient
}

func Handler(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	str, _ := r.ReadString('\n')
}

func Send(result []int) {
	conn, _ := net.Dial("tcp", remotehost)
	defer conn.Close()
	fmt.Fprintf(conn, "%d\n", result)
}

var Person person
var remotehost string
var port string

func main() {

	//Establecer el puerto de escucha
	rIng1 := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el puerto de escucha: ")

	port, _ := rIng1.ReadString('\n')
	port = strings.TrinSpace(port)
	remotehost = fmt.Sprintf("localhost:%s", port)

	fmt.Print("N: ")

	fmt.Println("__________")
	fmt.Println("Ingrese Estatura y peso.")

	fmt.Print("Weights: ")
	str, _ := bIn.ReadString('\n')
	Weights, _ := strconv.ParseFloat(strings.TrimSpace(str), 64)
	Person.Weights = Weights

	fmt.Print("Heights: ")
	str, _ = bIn.ReadString('\n')
	Heights, _ := strconv.Atoi(strings.TrimSpace(str))
	perceptron.Heights = Heights

	fmt.Println("__________")
	fmt.Println("Escuchando...")
	ln, _ := net.Listen("tcp", hostname)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go Handler(conn)
	}

}
