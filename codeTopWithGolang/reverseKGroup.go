package main

/**
题目描述：
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
*/

/**
解题思路：
1. 我们的目标是每 k 个节点为一组进行翻转，如果最后一组不足 k 个节点，则保持原样。
2. 我们需要遍历整个链表，对每组 k 个节点进行翻转，然后将翻转后的组重新连接到链表中。
*/

func reverseKGroup(head *ListNode, k int) *ListNode {

	// 创建一个虚拟头节点，简化对头节点的处理
	dummy := &ListNode{Next: head}
	// prev 指向当前要处理的k个节点的前一个节点
	prev := dummy

	for head != nil {
		// tail 指向从 head 开始的第 k 个节点
		tail := prev

		// 寻找第 k 个节点
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				// 如果 tail 为空，说明剩余节点不足 k 个
				return dummy.Next
			}
		}
		// 保存下一组的起始节点
		next := tail.Next

		// 翻转从 head 到 tail 的这一组节点
		head, tail = reverseListByK(head, tail)

		// 将翻转后的这一组节点连接回原链表
		prev.Next = head
		tail.Next = next

		// 更新 prev 和 head 到下一组
		prev = tail
		head = tail.Next
	}
	return dummy.Next
}

// reverseList 翻转从 head 到 tail 的这一组节点
func reverseListByK(head, tail *ListNode) (*ListNode, *ListNode) {
	// prev 初始化为 tail 的下一个节点
	prev := tail.Next
	curr := head

	// 当 prev 不等于 tail 时继续翻转
	for prev != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return tail, head
}
