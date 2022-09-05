package main

import "C"
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// #include <stdio.h>
// #include <stdlib.h>
// #include <immintrin.h>

func main() {

	in := bufio.NewReader(os.Stdin)

	n := readInt(in)

	a := make([][]uint64, n)
	b := make([][]uint64, n)
	c := make([][]uint64, n)

	for i, _ := range a {
		a[i] = make([]uint64, n)
		b[i] = make([]uint64, n)
		c[i] = make([]uint64, n)

		for j, _ := range a[i] {
			a[i][j] = uint64(rand.Intn(1000) + 1)
			b[i][j] = uint64(rand.Intn(1000) + 1)
		}
	}

	wg := sync.WaitGroup{}
	start := time.Now()
	//calcMatrix_v1(&wg, c, a, b)
	//calcMatrix_v0(c, a, b, n)
	//calcMatrix_v2(c, a, b, n)
	//calcMatrix_v3(c, a, b, n)
	c = calcMatrix_v4(a, b, n)
	//calcMatrix_v5(&wg, c, a, b, n)
	//calcMatrixByUseAvx(c, a, b, n)
	wg.Wait()

	finish := time.Since(start).Seconds()
	//printMatrix(a)
	//printMatrix(b)
	//printMatrix(c)

	fmt.Println("Work time in secs ", finish)
}

func calcMatrix_v1(wg *sync.WaitGroup, c [][]uint64, a [][]uint64, b [][]uint64) {
	for i, val := range c {
		for j, _ := range val {
			wg.Add(1)
			go func(i int, j int) {
				defer wg.Done()
				sum := uint64(0)
				for t, val := range a[i] {
					sum += val * b[t][j]
				}
				c[i][j] = sum

			}(i, j)
		}

	}
}

func calcMatrix_v0(c [][]int, a [][]int, b [][]int, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
}

func calcMatrix_v2(c [][]int, a [][]int, b [][]int, size int) {
	for i := 0; i < size; i++ {

		for k := 0; k < size; k++ {
			for j := 0; j < size; j++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
}

func calcMatrix_v3(c [][]int, a [][]int, b [][]int, size int) {
	for i := 0; i < size; i++ {
		Cc := make([]int, size)
		for k := 0; k < size; k++ {
			aa := a[i][k]
			bb := b[k]
			for j := 0; j < size; j++ {
				Cc[j] += aa * bb[j]
			}
		}
		c[i] = Cc
	}
}

func calcMatrix_v4(a [][]uint64, b [][]uint64, size int) [][]uint64 {
	c := make([][]uint64, size)
	wg := sync.WaitGroup{}
	for i := 0; i < size; i++ {

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			temp := make([]uint64, size)
			for k := 0; k < size; k++ {
				aa := a[i][k]
				bb := b[k]
				for j := 0; j < size; j++ {
					temp[j] += aa * bb[j]
				}
			}

			c[i] = temp
		}(i)
	}

	wg.Wait()
	return c
}

func calcMatrix_v5(wg *sync.WaitGroup, c [][]uint64, a [][]uint64, b [][]uint64, size int) {
	for i := 0; i < size; i++ {

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			temp := make([]uint64, size)
			for k := 0; k < size; k++ {
				aa := a[i][k]
				bb := b[k]
				for j := 0; j < size; j++ {
					temp[j] += aa * bb[j]
				}
			}

			c[i] = temp
		}(i)
	}
}

//func calcMatrixByUseAvx(c [][]int, a [][]int, b [][]int, size int) {
//	for i := 0; i < size; i++ {
//		for j := 0; j < size; j += 8 {
//			C._mm256_storeu_ps(unsafe.Pointer(&c[j][0]), C._mm256_setzero_ps())
//		}
//		for k := 0; k < size; k++ {
//			a = C._mm256_set1_ps(a[i][k])
//		}
//		for j := 0; j < size; j += 16 {
//			_ = C._mm256_storeu_ps(unsafe.Pointer(&c[j][0]), C._mm256_fmadd_ps(a, C._mm256_loadu_ps(unsafe.Pointer(&b[j][0]),
//				C._mm256_loadu_ps(unsafe.Pointer(&c[j][0])))))
//			_ = C._mm256_storeu_ps(unsafe.Pointer(&c[j][8]), C._mm256_fmadd_ps(a, C._mm256_loadu_ps(unsafe.Pointer(&b[j][8]),
//				C._mm256_loadu_ps(unsafe.Pointer(&c[j][8])))))
//		}
//	}
//}

func printMatrix(a [][]uint64) {
	for _, val := range a {
		fmt.Println(val)
	}
}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}
