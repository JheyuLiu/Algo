/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		root := &TreeNode{Val: nums[0]}
		return root
	}

	idx := len(nums)/2
    root := &TreeNode{Val: nums[idx]}
    
    root.Left = sortedArrayToBST(nums[0:idx])
    root.Right = sortedArrayToBST(nums[idx+1:len(nums)])

    return root
}