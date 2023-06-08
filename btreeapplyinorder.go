package piscine

type TreeNodee struct {
	Left, Right, Parent *TreeNodee
	Data                string
}

func BTreeApplyInorder(root *TreeNodee, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}
