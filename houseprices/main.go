package main

import (
	"bufio"
	"fmt"
	"gonum.org/v1/gonum/stat"
	"log"
	"os"

	"gonum.org/v1/gonum/mat"
)

func main() {
	m, err := readData("data/houseprices.csv")
	if err != nil {
		log.Fatal(err)
	}

	r, c := m.Dims()

	var X mat.Dense
	X.Augment(Ones(r, 1), m.Slice(0, r, 0, c-1))
	normalizeFeatures(&X)
	fmt.Println(mat.Formatted(&X))

	y := m.Slice(0, r, c-1, c)

	theta := mat.NewDense(c, 1, nil)

	num_iters := 400
	alpha := 0.01

	for i := 0; i < num_iters; i++ {
		delta := gradient(&X, y, theta)
		delta.Scale(alpha, delta)
		theta.Sub(theta, delta)
		fmt.Printf("Iteration %d, cost %f\n", i, cost(&X, y, theta))
	}

	fmt.Printf("theta:\n%v\ncost: %f\n", mat.Formatted(theta, mat.Squeeze()), cost(&X, y, theta))
}

func normalizeFeatures(m *mat.Dense) {
	_, c := m.Dims()

	var mus, sigmas []float64
	for j := 0; j < c; j++ {
		xs := mat.Col(nil, j, m)
		mu, sigma := stat.MeanStdDev(xs, nil)
		mus = append(mus, mu)
		sigmas = append(sigmas, sigma)
	}

	m.Apply(func(i, j int, v float64) float64 {
		mu := mus[j]
		sigma := sigmas[j]
		if sigma != 0.0 {
			return (v - mu) / sigma
		}
		return v
	},
		m)

	return
}

func gradient(X mat.Matrix, y mat.Matrix, theta mat.Matrix) *mat.Dense {
	m, _ := X.Dims()
	var T, U mat.Dense
	T.Mul(X, theta)
	T.Sub(&T, y)
	U.Mul(X.T(), &T)
	U.Scale(1.0/float64(m), &U)
	return &U
}

func cost(X mat.Matrix, y mat.Matrix, theta mat.Matrix) float64 {
	m, _ := X.Dims()

	var T mat.Dense
	T.Mul(X, theta)
	T.Sub(&T, y)
	T.Apply(func(i, j int, v float64) float64 { return v * v }, &T)

	return mat.Sum(&T) / (2.0 * float64(m))
}

func Ones(r, c int) mat.Matrix {
	data := make([]float64, r*c)
	for i := range data {
		data[i] = 1
	}
	return mat.NewDense(r, c, data)
}

func readData(path string) (*mat.Dense, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data := make([]float64, 0)
	r := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		r++
		var a, b, c float64
		_, err := fmt.Sscanf(scanner.Text(), "%f,%f,%f", &a, &b, &c)
		if err != nil {
			return nil, fmt.Errorf("Reading %s, parse error at line %d: %v", r, err)
		}
		data = append(data, a, b, c)
	}
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("Reading %s, scanner error: %v", path, err)
	}

	return mat.NewDense(r, 3, data), nil
}
