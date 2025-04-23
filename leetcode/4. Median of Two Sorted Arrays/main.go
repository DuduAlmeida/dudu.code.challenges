package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	list := mergeSortedArrays(nums1, nums2)

	return medianFromArray(list)
}

func mergeSortedArrays(nums1 []int, nums2 []int) (mergedList []int) {
	mergedList = []int{}

	m := len(nums1)
	n := len(nums2)
	m_n := m + n

	i := 0
	j := 0

	for len(mergedList) < m_n {
		isJ_OutOfRange := j >= n
		isI_OutOfRange := i >= m

		if isI_OutOfRange && isJ_OutOfRange {
			break
		}

		number1 := 1001 //default it is above the maximum, for out of range cases
		if !isI_OutOfRange {
			number1 = nums1[i]
		}
		number2 := 1001 //default it is above the maximum, for out of range cases
		if !isJ_OutOfRange {
			number2 = nums2[j]
		}

		if number1 < number2 {
			mergedList = append(mergedList, number1)
			i++
			continue
		}

		mergedList = append(mergedList, number2)
		j++
	}

	fmt.Println(mergedList)

	return mergedList
}

func medianFromArray(list []int) float64 {
	size := len(list)

	if size == 0 {
		return 0
	}
	if size == 1 {
		return float64(list[0])
	}

	firstIndex := int((size - 1) / 2)
	secondIndex := firstIndex + 1

	firstNumber := float64(list[firstIndex])
	secondNumber := float64(list[secondIndex])

	fmt.Println("firstNumber", firstNumber, "secondNumber", secondNumber)

	if isOdd := size%2 == 1; isOdd {
		return firstNumber
	}

	return (firstNumber + secondNumber) / 2
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
