package common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetCommonData(t *testing.T) {
	if ans := GetCommonData(); ans != "Common Data" {
		t.Errorf("CommonData must be Common Data, but %s got", ans)
	}

}

func TestGetCommonData1(t *testing.T) {
	expected := "Common Data"
	actual := GetCommonData()

	assert.Equal(t, expected, actual, "Expected and actual data should be equal")
}

func TestGetData(t *testing.T) {
	expectedData := "Data"
	actualData := GetData()

	assert.Equal(t, expectedData, actualData, "GetData() should return the expected data")
}

func TestMatrixMultiplication(t *testing.T) {
	// 测试正常情况
	matrixA := [][]int{{1, 2}, {3, 4}}
	matrixB := [][]int{{5, 6}, {7, 8}}
	expectedResult := [][]int{{19, 22}, {43, 50}}
	result := MatrixMultiplication(matrixA, matrixB)
	assert.Equal(t, expectedResult, result, "Matrix multiplication result is incorrect")

	// 测试矩阵维度不兼容情况
	matrixC := [][]int{{1, 2}, {3, 4}}
	matrixD := [][]int{{5, 6}, {7}}
	assert.Panics(t, func() { MatrixMultiplication(matrixC, matrixD) }, "Expected panic for incompatible matrix dimensions")
}

func TestMatrixMultiplication1(t *testing.T) {
	// 生成两个1000x1000的大矩阵
	matrixA := GenerateLargeMatrix(1000, 1000)
	matrixB := GenerateLargeMatrix(1000, 1000)

	// 执行矩阵乘法
	startTime := time.Now()
	result := MatrixMultiplication(matrixA, matrixB)
	// 打印矩阵的前10行和前10列
	fmt.Println("Matrix result (first 10 rows and columns):")
	for i := 0; i < 10; i++ {
		fmt.Println(result[i][:10])
	}

	elapsedTime := time.Since(startTime)

	// 打印执行时间
	fmt.Printf("Matrix multiplication took: %v\n", elapsedTime)

	// 验证结果的正确性（这里可以根据需要添加验证逻辑）
	// 例如，检查结果的维度是否正确，或者检查特定位置的值是否符合预期

}

func TestGenerateLargeMatrix(t *testing.T) {
	// 生成两个1000x1000的大矩阵
	matrixA := GenerateLargeMatrix(1000, 1000)
	matrixB := GenerateLargeMatrix(1000, 1000)

	// 打印矩阵A的前10行和前10列
	fmt.Println("Matrix A (first 10 rows and columns):")
	for i := 0; i < 10; i++ {
		fmt.Println(matrixA[i][:10])
	}

	// 打印矩阵B的前10行和前10列
	fmt.Println("\nMatrix B (first 10 rows and columns):")
	for i := 0; i < 10; i++ {
		fmt.Println(matrixB[i][:10])
	}
}
