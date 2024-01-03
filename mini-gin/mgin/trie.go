package mgin

import "strings"

type node struct {
	pattern  string  //url to be pattern,eg: /p/:lang, only if this node is lear node ,pattern is not null
	part     string  // current url, eg: /:lang
	children []*node // childNode
	isWide   bool    //if pattern exactly, if part contail ':' or '*',value is true
}

// find first match node
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWide {
			return child
		}
	}
	return nil
}

// find all match node
func (n *node) matchAllChild(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWide {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// add node to trie tree
func (n *node) insert(pattern string, parts []string, heignt int) {
	if len(parts) == heignt {
		n.pattern = pattern
		return
	}
	part := parts[heignt]
	child := n.matchChild(part)

	if child == nil {
		child = &node{part: part, isWide: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, heignt+1)
}

//search from trie tree
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchAllChild(part)
	for _, child := range children {
		res := child.search(parts, height+1)
		if res != nil {
			return res
		}
	}
	return nil
}
