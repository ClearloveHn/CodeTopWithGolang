package main

/**
题目描述：
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：
输入：l1 = [1,2,4], l2 = [1,3,4]
*/

/**
解题思路：
1.创建一个虚拟头节点(dummy),用于简化边界情况的处理。
2.使用两个指针分别遍历两个链表。
3.比较两个指针所指节点的值,将较小的节点添加到结果链表中,并移动该指针。
4.重复步骤3,直到其中一个链表遍历完毕。
5.将剩余链表的节点直接接到结果链表的末尾。
6.返回虚拟头节点的下一个节点,即为合并后的链表头。
*/

type ListNodeMerge struct {
	Val  int
	Next *ListNodeMerge
}

func mergeTwoLists(l1 *ListNodeMerge, l2 *ListNodeMerge) *ListNodeMerge {
	// 创建虚拟头节点
	dummy := &ListNodeMerge{Val: 0}
	// 使用current指针来构建新的链表
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}
