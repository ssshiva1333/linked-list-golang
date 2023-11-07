package linkedlist

type node struct {
	value int
	next  *node
}

type linked_list struct {
	head *node
}

func (ll *linked_list) add(value int) {
	new_node := node{value: value, next: nil}

	if ll.head == nil {
		ll.head = &new_node
	} else {
		node := ll.head
		for node.next != nil {
			node = node.next
		}
		node.next = &new_node
	}

	return
}

func (ll *linked_list) remove(value int) {
	node := ll.head
	for node.next != nil {
		if node.value == value {
			node.next = node.next.next
		} else {
			node = node.next
		}
	}

	return
}

func (ll *linked_list) show() (elements []int) {
	node := ll.head
	for node.next != nil {
		elements = append(elements, node.value)
		node = node.next
	}

	return
}
