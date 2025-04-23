package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	mergeKLists(nil)
	mergeKLists([]*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}},
	})

}

func mergeKLists(lists []*ListNode) *ListNode {
	numbers := []int{}
	for _, list := range lists {
		numbers = append(numbers, getValuesFromListNode(list)...)
	}

	numbers = sortNumbers(numbers)

	return transformNumberstInListNode(numbers)
}

func getValuesFromListNode(ln *ListNode) []int {
	if ln == nil {
		return []int{}
	}

	if ln.Next != nil {
		return append(getValuesFromListNode(ln.Next), ln.Val)
	}

	return []int{ln.Val}
}

func sortNumbers(numbers []int) []int {
	size := len(numbers)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			firstNumber := numbers[i]
			secondNumber := numbers[j]

			if firstNumber > secondNumber {
				numbers[i] = secondNumber
				numbers[j] = firstNumber
			}
		}
	}

	return numbers
}

func transformNumberstInListNode(numbers []int) (ln *ListNode) {
	fmt.Println(numbers)
	if len(numbers) < 1 {
		return nil
	}

	var lastNode *ListNode
	for index, number := range numbers {
		if index == 0 {
			ln = &ListNode{
				Val: number,
			}
			lastNode = ln
			continue
		}

		lastNode.Next = &ListNode{
			Val: number,
		}
		lastNode = lastNode.Next
	}

	return ln
}
