package piscine

type NodeIe struct {
	Data int
	Next *NodeIe
}

func SortListInsert(l *NodeIe, data_ref int) *NodeIe {
	n := &NodeIe{Data: data_ref}
	n.Next = nil

	if l == nil || l.Data >= n.Data {
		n.Next = l
		return n
	}
	temp := l
	for temp.Next != nil && temp.Next.Data < n.Data {
		temp = temp.Next
	}
	n.Next = temp.Next
	temp.Next = n

	return l
}
