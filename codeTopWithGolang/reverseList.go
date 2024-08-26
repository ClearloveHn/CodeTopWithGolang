package main

import "fmt"

/**
题目描述：
给你单链表的头节点head ，请你反转链表，并返回反转后的链表。

示例1：
输入：head = 1->2->3->4->5
输出：5->4->3->2->1
*/

/**
解题思路：
1. 初始化一个反转后的空链表，初始化一个用于遍历原链表的头节点
2. 遍历结束条件 = 原链表的节点为nil
3. 反转核心：当前遍历链表的头节点的next指针指向反转后的链表的头节点，然后更新反转后链表的头节点为当前遍历链表节点的val，依次遍历指向更新。最终反转。
4. 需要创建一个节点用于保存当前遍历节点的下一个节点。用于遍历用，因为改变了当前遍历节点next指针的指向。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode // 初始化反转后的链表。
	pre := head           // 指向原链表的头节点,用于遍历原链表。

	for pre != nil {
		next := pre.Next   // 保存当前节点的下一个节点,因为我们即将改变 pre.Next 的指向
		pre.Next = newHead // 将当前节点的 Next 指针存到新链表头部。
		newHead = pre      // 更新新链表的头部为当前节点,因为我们刚刚将它插入到了最前面。
		pre = next         //移动到原链表的下一个节点,继续遍历。
	}

	return newHead
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {

	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	fmt.Println("Original list:")
	printList(head)

	reverseHead := reverseList(head)

	fmt.Println("Reversed list:")
	printList(reverseHead)
}
