package main

func merge2SortedLists(l1, l2 *LinkedListNode) *LinkedListNode {
	dummy := &LinkedListNode{data: -1}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.data <= l2.data {
			prev.next = l1
			l1 = l1.next
		} else {
			prev.next = l2
			l2 = l2.next
		}
		prev = prev.next
	}

	if l1 == nil {
		prev.next = l2
	} else {
		prev.next = l1
	}

	return dummy.next
}

func mergeKSortedLists(lists []*LinkedListNode) *LinkedListNode {
	if len(lists) > 0 {
		res := lists[0]

		for i := 1; i < len(lists); i++ {
			res = merge2SortedLists(res, lists[i])
		}
		return res
	}

	return &LinkedListNode{data: -1}
}

func main() {
	a := LinkedList([]int{11, 41, 51})
	b := LinkedList([]int{21, 23, 42})
	c := LinkedList([]int{25, 56, 66, 72})
	list1 := []*LinkedListNode{a, b, c}
	display(mergeKSortedLists(list1))
}
