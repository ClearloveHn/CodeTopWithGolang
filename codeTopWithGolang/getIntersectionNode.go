package main

/**
题目描述：
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

示例1：
输入 intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
输出 Intersected at '8'
*/

/**
解题思路：
1.分别计算两个链表的长度。
2.让较长的链表先走差值步。
3.然后两个链表同时走，直到找到相交点或到达末尾。
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := getLength(headA), getLength(headB)

	// 对齐链表起点
	for lenA > lenB {
		headA = headA.Next
		lenA--
	}

	for lenB > lenA {
		headB = headB.Next
		lenB--
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}
