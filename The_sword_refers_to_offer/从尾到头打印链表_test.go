package The_sword_refers_to_offer

func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	var bk func(head *ListNode)
	bk = func(head *ListNode) {
		if head == nil {
			return
		}
		if head.Next != nil {
			bk(head.Next)
		}
		res = append(res, head.Val)
	}
	bk(head)
	return res
}
