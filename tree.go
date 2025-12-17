/* A red black tree is a binary search tree which follows the following 4 rules:
1. Every node is either red or black.
2. The root is black
3. Every path from a node to any of its descedant NIL nodes have the same number of black nodes
4. Every red node must have two black child nodes
*/

package main

type Node struct {
	Value  int
	Color  string
	Left   *Node
	Right  *Node
	Parent *Node
}

type RedBlackTree struct {
	Root *Node
}

// O(n) is max time complexity
func (t *RedBlackTree) Insert(root *Node, value int) {
	if root == nil {
		return
	}

	// New nodes will always be appened to the "end" of the tree
}
