package main

type node struct {
	key    float32
	value  interface{}
	parent *node
	left   *node
	right  *node
}

// Tree type is exported.
type Tree struct {
	root *node
}

// Walk tree in order calling fn for each node
func (n *node) inOrder(fn func(n *node)) {
	if n != nil {
		n.left.inOrder(fn)
		fn(n)
		n.right.inOrder(fn)
	}
}

// Flatten tree and return in-order array of orders
func (tree *Tree) Flatten() []interface{} {
	var values []interface{}
	fn := func(n *node) {
		values = append(values, n.value)
	}
	tree.root.inOrder(fn)
	return values
}

// Create new node with (key, value) pair.
func newNode(key float32, value interface{}) *node {
	var n node
	n.key = key
	n.value = value
	return &n
}

// Add (key, value) pair to tree
func (tree *Tree) Add(key float32, value interface{}) {
	tree.insert(newNode(key, value))
}

// Insert node into tree
func (tree *Tree) insert(new *node) {
	if tree.root == nil {
		tree.root = new
		return
	}
	// Find insert position
	var p *node
	cur := tree.root
	for cur != nil {
		p = cur
		if new.key < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	// insert node
	if new.key < p.key {
		p.left = new
	} else {
		p.right = new
	}
	new.parent = p
}
