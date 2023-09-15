package main

import (
	"fmt"
	"math"
)

type Funcion interface {
	NumV() int
	NunO() int
	NumVs() int
	NunOs() int
}

type Vc struct {
	z float32 //nivel de confianza
	e float32 //erro
	N float32 //poblacion
	o float32 //desviacion estandar
}

type VS struct {
	Vc
	p float32 //probabilidad de exito
	q float32 //probabilidad de fracaso
}

func main() {

	one := Vc{N: 500, z: 1.96, o: 0.5, e: 0.09}

	fmt.Println(one.NunO())

	//two := Vc{z: 1.96, e: 0.05, p: 0.03, q: 97.7}

	three := Vc{z: 1.96, e: 0.05, o: 0.70}
	fmt.Println(three.NumV())
}

func (a Vc) NumV() int {
	n := math.Pow(float64(a.z), 2) * math.Pow(float64(a.o), 2)
	n1 := math.Pow(float64(a.e), 2)
	sampleSize := float64(n) / float64(n1)

	return int(sampleSize)

}
func (x Vc) NunO() int {

	n := (x.N - 1)
	n1 := float64(x.N) * math.Pow(float64(x.z), 2) * math.Pow(float64(x.o), 2)
	n2 := float64(n)*math.Pow(float64(x.e), 2) + (math.Pow(float64(x.z), 2) * math.Pow(float64(x.o), 2))
	nfin := float64(n1) / float64(n2)

	return int(nfin)
}

// la siguente forma  V SIMPLE

func (c VS) NumVs() int {
	n2 := c.p * c.q
	n := math.Pow(float64(c.z), 2) * float64(n2)
	n1 := math.Pow(float64(c.e), 2)
	sampleSize := float64(n) / float64(n1)

	return int(sampleSize)
}

func (v VS) NunOs() int {

	n := (v.N - 1)
	n9 := v.p * v.q
	n1 := float64(v.N) * math.Pow(float64(v.z), 2) * float64(n9)
	n2 := float64(n)*math.Pow(float64(v.e), 2) + (math.Pow(float64(v.z), 2) * float64(n9))
	nfin := float64(n1) / float64(n2)

	return int(nfin)
}

// uso de interface

func datos_reflejados(f Funcion, v VS) string {
	var a string
	if v.p != 0 && v.o == 0 {
		a = "es un variable simple simple"
	} else if v.o == 0 {
		a = "es un variable simple compuesta"
	} else if v.N != 0 && v.p != 0 {
		a = "es un variable simple compuesta"
	} else {
		a = " es una variable compuesta compuesta "
	}

	x := f.NumV() + f.NumVs() + f.NunO() + f.NunOs()

	fmt.Println(x)
	return a
}
