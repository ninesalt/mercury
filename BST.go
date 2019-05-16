package main

type node struct {
	order  Order
	key    float32
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
func (tree *Tree) Flatten() []Order {
	var orders []Order
	fn := func(n *node) {
		orders = append(orders, n.order)
	}
	tree.root.inOrder(fn)
	return orders
}

func newNode(order Order) *node {
	var n node
	n.key = order.Price
	n.order = order
	return &n
}

// Add Order to tree
func (tree *Tree) Add(order Order) {
	tree.insert(newNode(order))
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
