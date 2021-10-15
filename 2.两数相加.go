/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	t := 0
	root := &ListNode{
		Val:  0,
		Next: nil,
	}
	node := root
	for l1 != nil && l2 != nil {
		t = l1.Val + l2.Val + flag
		if t >= 10 {
			flag = 1
			t = t - 10
		} else {
			flag = 0
		}
		node.Next = &ListNode{
			Val:  t,
			Next: nil,
		}
		node = node.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	//相同长度
	if l1 == nil && l2 == nil {
		if flag != 0 {
			node.Next = &ListNode{
				Val:  flag,
				Next: nil,
			}
			flag = 0
		}
	}
	//不同长度
	var tmpNode *ListNode
	if l1 != nil && l2 == nil {
		tmpNode = l1
	} else {
		tmpNode = l2
	}
	for tmpNode != nil {
		t = tmpNode.Val + flag
		if t >= 10 {
			flag = 1
			t = t - 10
		} else {
			flag = 0
		}
		node.Next = &ListNode{
			Val:  t,
			Next: nil,
		}
		node = node.Next
		tmpNode = tmpNode.Next
	}
	if flag != 0 {
		node.Next = &ListNode{
			Val:  flag,
			Next: nil,
		}
		flag = 0
	}
	return root.Next
}

// @lc code=end

