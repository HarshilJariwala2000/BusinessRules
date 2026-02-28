package utils

type ArrayList interface {
	int | int32 | int64 | uint | string | float32 | float64
}

func Map [T any, R any] (collection []T, iteratee func(T) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = iteratee(item)
	}
	return result
}

func ArrayDifference [T ArrayList] (array1, array2 []T) []T {
	array2Map := make(map[T]struct{})
	for _, elem := range array2 {
		array2Map[elem] = struct{}{}
	}

	var difference []T
	for _, elem := range array1 {
		if _, ok := array2Map[elem]; !ok {
			difference = append(difference, elem)
		}
	}
	return difference
}

func RemoveArrayDuplicates [T ArrayList] (array []T) []T{
	var result []T
	hashMap := make(map[T]int)
	for _, value := range array{
		if _, ok := hashMap[value]; !ok {
			hashMap[value] = 1
		}else{
			hashMap[value]++
		}
	}
	for key := range hashMap{
		result = append(result, key)
	}
	return result
}