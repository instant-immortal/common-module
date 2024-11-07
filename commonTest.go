package common

import (
	"fmt"
	"math/rand"
	"time"
)

// GetCommonData a method
func GetCommonData() string {
	return "Common Data"
}

func GetData() string {
	return "Data"
}

// MatrixMultiplication 矩阵乘法
func MatrixMultiplication(matrixA [][]int, matrixB [][]int) [][]int {
	// 获取矩阵 A 的行数和列数
	rowsA := len(matrixA)
	colsA := len(matrixA[0])

	// 获取矩阵 B 的行数和列数
	rowsB := len(matrixB)
	colsB := len(matrixB[0])

	// 检查矩阵 A 的列数是否等于矩阵 B 的行数
	if colsA != rowsB {
		panic("Matrix dimensions are not compatible for multiplication")
	}

	// 创建一个新的矩阵来存储结果
	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	// 执行矩阵乘法
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result
}

func MatrixMultiplicationTest() {
	// 示例矩阵 A 和 B
	matrixA := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	matrixB := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	// 计算矩阵乘法
	result := MatrixMultiplication(matrixA, matrixB)

	// 打印结果矩阵
	fmt.Println("Result of matrix multiplication:")
	for _, row := range result {
		fmt.Println(row)
	}
}

// GenerateLargeMatrix 生成一个大矩阵
func GenerateLargeMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Intn(100) // 生成0到99之间的随机整数
		}
	}
	return matrix
}
