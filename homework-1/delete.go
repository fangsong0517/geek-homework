package main

import "fmt"

// 删除指定index
func delete(array []int, idx int) ([]int, error) {
	if idx < 0 || idx >= len(array) {
		return nil, fmt.Errorf("idx invalid")
	}
	arrRes := make([]int, 0, len(array)-1)
	arrRes = append(arrRes, array[:idx]...)
	arrRes = append(arrRes, array[idx+1:]...)
	return arrRes, nil
}

// 高性能删除指定元素
func deleteNew(array []int, idx int) ([]int, error) {
	if idx < 0 || idx >= len(array) {
		return nil, fmt.Errorf("idx invalid")
	}
	copy(array[idx:], array[idx+1:])
	return array[:len(array)-1], nil
}

// 改造为泛型
func deleteGeneric[T any, V uint | uint32](array []T, idx V) ([]T, error) {
	arrLen := len(array)
	if idx < 0 || idx >= V(arrLen) {
		return nil, fmt.Errorf("idx invalid")
	}
	arrRes := make([]T, 0, arrLen-1)
	arrRes = append(arrRes, array[:idx]...)
	arrRes = append(arrRes, array[idx+1:]...)
	return arrRes, nil
}

func main() {
	arr := []int{0, 1, 2, 3, 4}
	w, err := delete(arr, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(w, arr)

	arr = []int{0, 1, 2, 3, 4}
	w, err = deleteNew(arr, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(w, arr)

	arr = []int{0, 1, 2, 3, 4}
	w, err = deleteGeneric[int, uint32](arr, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(w, arr)
}
