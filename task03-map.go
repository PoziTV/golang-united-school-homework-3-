package homework

import (
	"fmt"
	"sort"
)

type mapIntStr struct {
	key int
	val string
}

func sortMapValues(input map[int]string) (result []string) {
	sliceToSort := []mapIntStr{}

	for key, val := range input {
		sliceToSort = append(sliceToSort, mapIntStr{key, val})
	}

	fmt.Println(sliceToSort)

	sort.Slice(sliceToSort, func(i, j int) bool {
		return sliceToSort[i].key < sliceToSort[j].key
	})

	fmt.Println(sliceToSort)
	for _, val := range sliceToSort {
		result = append(result, val.val)
	}

	return
}
