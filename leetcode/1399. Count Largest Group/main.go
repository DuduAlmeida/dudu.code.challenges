package main

import "fmt"

func countLargestGroup(n int) int {
	groups := NumberGroupList{}

	for i := 1; i <= n; i++ {
		groups.Add(i)
	}

	return len(groups.FilterByBiggerSizes())
}

type NumberSeparetedByDigits []int

func (nsbd NumberSeparetedByDigits) Sum() (sum int) {
	sum = 0
	for _, digit := range nsbd {
		sum = sum + digit
	}

	return sum
}

func NewNumberSeparetedByDigits(number int) NumberSeparetedByDigits {
	list := NumberSeparetedByDigits{}
	if number == 0 {
		return []int{0}
	}

	for number > 0 {
		digit := number % 10
		list = append(list, digit)
		number /= 10
	}

	return list
}

type NumberGroup struct {
	Items []int
	Sum   int
}

func (ng NumberGroup) Size() int {
	return len(ng.Items)
}

type NumberGroupList []NumberGroup

func (list *NumberGroupList) FindIndexBySum(sum int) (index int) {
	for i, numberGroup := range *list {
		if sum == numberGroup.Sum {
			return i
		}
	}

	return -1
}

func (list *NumberGroupList) Add(i int) {
	iDigits := NewNumberSeparetedByDigits(i)
	sum := iDigits.Sum()

	index := list.FindIndexBySum(sum)

	if isNewGroup := index == -1; isNewGroup {
		newGroup := NumberGroup{
			Items: []int{i},
			Sum:   sum,
		}
		*list = append(*list, newGroup)
		fmt.Println(newGroup)
		return
	}

	groupUpdated := (*list)[index]
	(*list)[index].Items = append(groupUpdated.Items, i)
	fmt.Println(groupUpdated)
}

func (list *NumberGroupList) FilterByBiggerSizes() (filteredList NumberGroupList) {
	maxSize := 0
	filteredList = NumberGroupList{}

	for _, numberGroup := range *list {
		size := numberGroup.Size()

		fmt.Println("size", size)

		if size > maxSize {
			maxSize = size
			filteredList = NumberGroupList{numberGroup}
			continue
		}

		if maxSize == size {
			filteredList = append(filteredList, numberGroup)
		}
	}

	return filteredList
}

func main() {
	fmt.Println(countLargestGroup(13))
	fmt.Println(countLargestGroup(2))
}
