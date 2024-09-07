package main

/**
题目描述：
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

示例1：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
*/

/**
解题思路：
1.层序遍历本质上是对二叉树进行广度优先搜索。我们需要逐层地访问树的节点，这正是BFS的特点。
2.队列是实现BFS的理想数据结构。它能确保我们按照"先进先出"（FIFO）的顺序处理节点，这正好符合层序遍历的要求。
3.我们需要清楚地区分每一层的节点。这可以通过在每一轮迭代中记录当前层的节点数来实现。
4.使用Go中的切片来存储结果，其中每个内部数组代表一层的节点值。

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	res := make([][]int, 0)

	if root == nil {
		return res
	}

	// 创建队列并将根节点入队
	queue := []*TreeNode{root}

	// 当队列不为空时，继续遍历
	for len(queue) > 0 {
		// 记录当前层的节点数
		levelLen := len(queue)
		// 用于存储当前层节点值的切片
		curLevel := make([]int, 0)

		// 遍历当前层的所有节点
		for i := 0; i < levelLen; i++ {
			// 出队一个节点
			node := queue[0]  // 取出队列的第一个节点
			queue = queue[1:] // 将这个节点从队列中移除

			// 将节点的值添加到当前层的结果中
			curLevel = append(curLevel, node.Val)

			// 如果左子节点存在，将其入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// 如果右子节点存在，将其入队
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 将当前层的结果添加到最终结果中
		res = append(res, curLevel)
	}

	return res
}
