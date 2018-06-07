package main

import "fmt"

func searchroot(in []int, root int, size int) int {
	for i := 0; i < size; i++ {
		if in[i] == root {
			return i;
		}
	}

	return -1;
}

func printPostOrder(in []int, pre []int, size int) {
	root := searchroot(in, pre[0], size)
    //fmt.Println(in)
    //fmt.Println(pre)
    //fmt.Println("size: ", size)
    //fmt.Println("root: ", root)
    if root != 0 {
    	tmp_pre := pre[1:]
    	printPostOrder(in, tmp_pre, root)
    }

    if root != size-1 {
    	tmp_in := in[root+1:]
    	tmp_pre := pre[root+1:]
    	printPostOrder(tmp_in, tmp_pre, size-root-1)
    }

    fmt.Print(pre[0], " ")
}

func main() {
	in := []int{4, 2, 5, 1, 3, 6}
	pre := []int{1, 2, 4, 5, 3, 6}
	n := len(in)

	fmt.Println("Postorder traversal ")
	printPostOrder(in, pre, n)
	fmt.Println()
}