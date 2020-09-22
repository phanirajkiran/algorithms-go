package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
	"testing"
)

func Test_buildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree2(preorder, inorder)

	var res []int
	btree.InOrder(root, &res)

	t.Log(res)
}