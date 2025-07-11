package main

import "testing"

func TestSumar(t *testing.T) {
	/*
		r := Suma(3, 4)
		e := 7
		if r != e {
			t.Errorf("Error al sumar: se esperaba %d, pero se obtuvo %d", e, r)
		}
	*/

	casos := []struct {
		a, b, e int
	}{
		{2, 3, 5},
		{4, 5, 9},
		{456, 123, 579},
		{0, -1, -1},
	}

	for _, caso := range casos {
		r := Suma(caso.a, caso.b)
		if r != caso.e {
			t.Errorf("Error al sumar %d + %d: se esperaba %d, pero se obtuvo %d", caso.a, caso.b, caso.e, r)
		}
	}
}

func TestMayor(t *testing.T) {

	casos := []struct {
		a, b, e int
	}{
		{2, 3, 3},
		{4, 5, 5},
		{456, 123, 456},
		{0, -1, 0},
		{0, 0, 0},
		{-1, -2, -1},
	}

	for _, caso := range casos {
		r := Mayor(caso.a, caso.b)
		if r != caso.e {
			t.Errorf("Error al comparar %d > %d: se esperaba %d, pero se obtuvo %d", caso.a, caso.b, caso.e, r)
		}
	}

}

func TestFibonacci(t *testing.T) {

	casos := []struct {
		n, e int
	}{
		{0, 0},
		{1, 1},
		{7, 13},
		{50, 12586269025},
	}

	for _, caso := range casos {
		r := Fibonacci(caso.n)
		if r != caso.e {
			t.Errorf("Error al calcular Fibonacci: se esperaba %d, pero se obtuvo %d", caso.e, r)
		}
	}

}
