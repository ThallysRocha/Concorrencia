package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var CntPair = 0

func createRandomMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}
	return matrix
}

func serie(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j]%2 == 0 {
				CntPair++
			}
		}
	}
	time.Sleep(1 * time.Nanosecond)
}

func parallel(matrix [][]int, id int, nThreads int, mutex *sync.Mutex, waitGroup *sync.WaitGroup) {
	lines := len(matrix)
	column := len(matrix[0])
	slice := lines * column / nThreads
	myCntPair := 0
	for i := slice * id; i < slice*(id+1); i++ {
		if matrix[i/column][i%column]%2 == 0 {
			myCntPair++
		}
	}
	time.Sleep(1 * time.Nanosecond)
	mutex.Lock()
	CntPair += myCntPair
	mutex.Unlock()

	waitGroup.Done()
}

func expSerie(qtdExp int, matrix [][]int) {
	fmt.Println("Exp Serie: ")
	for i := 0; i < qtdExp; i++ {
		CntPair = 0
		start := time.Now()
		serie(matrix)
		end := time.Now()
		fmt.Println(end.Sub(start).Nanoseconds())
	}
	fmt.Println("Serie: ", CntPair)
}

func expParallel(qtdExp int, nThreads int, matrix [][]int) {
	fmt.Println("\nExp Parallel: %d", nThreads)
	for i := 0; i < qtdExp; i++ {
		CntPair = 0
		mutex := &sync.Mutex{}
		waitGroup := &sync.WaitGroup{}
		waitGroup.Add(nThreads)

		start := time.Now()
		for j := 0; j < nThreads; j++ {
			go parallel(matrix, j, nThreads, mutex, waitGroup)
		}

		waitGroup.Wait()
		end := time.Now()
		fmt.Println(end.Sub(start).Nanoseconds())
	}
	fmt.Println("Parallel: ", nThreads, CntPair)
}

func main() {
	matrix := createRandomMatrix(10000)
	qtdExp := 100

	//Serie
	// fmt.Println("Exp Serie: ")
	// for i := 0; i < qtdExp; i++ {
	// 	CntPair = 0
	// 	start := time.Now()
	// 	serie(matrix)
	// 	end := time.Now()
	// 	fmt.Println(end.Sub(start).Nanoseconds())
	// }
	// fmt.Println("Serie: ", CntPair)
	expSerie(qtdExp, matrix)

	//Parallel
	// fmt.Println("\nExp Parallel 2: ")
	// for i := 0; i < qtdExp; i++ {
	// 	CntPair = 0
	// 	mutex := &sync.Mutex{}
	// 	waitGroup := &sync.WaitGroup{}
	// 	waitGroup.Add(nThreads)

	// 	start := time.Now()
	// 	for j := 0; j < nThreads; j++ {
	// 		go parallel(matrix, j, nThreads, mutex, waitGroup)
	// 	}

	// 	waitGroup.Wait()
	// 	end := time.Now()
	// 	fmt.Println(end.Sub(start).Nanoseconds())
	// }
	// fmt.Println("Parallel 2: ", CntPair)

	expParallel(qtdExp, 2, matrix)
	expParallel(qtdExp, 8, matrix)
	expParallel(qtdExp, 100, matrix)
	expParallel(qtdExp, 1000, matrix)

}
