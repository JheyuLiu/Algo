/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func help(root *TreeNode, slist *[]int) {
	if root == nil {
    	return
    }

    (*slist) = append((*slist), root.Val)
    sort.Ints((*slist))

    help(root.Left, slist)
    help(root.Right, slist)
}

func findSecondMinimumValue(root *TreeNode) int {
    slist := make([]int, 0, 100)
    help(root, &slist)
    
    min := slist[0]
    ans := math.MaxInt64
    
    for i:=1; i<len(slist); i++ {
    	if min < slist[i] && slist[i] < ans {
    		ans = slist[i]
    	}
    }

    if ans < math.MaxInt64 {
    	return ans
    } else {
    	return -1
    }
}