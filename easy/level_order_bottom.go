package easy

/*
	给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

	例如：
	给定二叉树 [3,9,20,null,null,15,7],

	     3
		/ \
	   9  20
	     /  \
		15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/

//LevelOrderBottom ...
func LevelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		l := len(queue)
		var list []int
		for cur := 0; cur < l; cur++ {
			list = append(list, queue[cur].Val)
			if queue[cur].Left != nil {
				queue = append(queue, queue[cur].Left)
			}
			if queue[cur].Right != nil {
				queue = append(queue, queue[cur].Right)
			}
		}
		result = append([][]int{list}, result...)
		queue = queue[l:]
	}
	return result
}
