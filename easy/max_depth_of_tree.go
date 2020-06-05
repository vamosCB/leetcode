package easy

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

//MaxDepth 递归的解法
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := MaxDepth(root.Left)
	r := MaxDepth(root.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}

// 使用深度优先遍历 和广度优先遍历 实现改问题
