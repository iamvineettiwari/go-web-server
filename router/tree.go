package router

import (
	"strings"

	"github.com/iamvineettiwari/go-web-server/http"
)

type Node struct {
	path      string
	children  []*Node
	isEnd     bool
	handler   http.HandlerFunc
	isParam   bool
	haveParam bool
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{
		root: &Node{
			children: []*Node{},
			isEnd:    false,
		},
	}
}

func (rt *Tree) Insert(path string, handler http.HandlerFunc) {
	if !strings.HasPrefix(path, "/") {
		panic("Path should start with /")
	}

	pathItems := strings.Split(path, "/")

	curNode := rt.root

	for idx := 0; idx < len(pathItems); idx++ {
		item := pathItems[idx]

		if item == "" {
			if idx == len(pathItems)-1 {
				item = "/"
			} else {
				continue
			}
		}

		itemNode := &Node{
			children: []*Node{},
		}

		if strings.HasPrefix(item, ":") {
			curNode.haveParam = true
			itemNode.isParam = true
			item = item[1:]
		}

		itemNode.path = item

		nextNode := rt.find(curNode, item)

		if nextNode == nil {
			curNode.children = append(curNode.children, itemNode)
			nextNode = itemNode
		}

		curNode = nextNode
	}

	curNode.isEnd = true
	curNode.handler = handler
}

func (rt *Tree) Resolve(path string) (http.HandlerFunc, http.Params) {
	pathItems := strings.Split(path, "/")

	curNode := rt.root
	params := make(http.Params)

	for idx := 0; idx < len(pathItems); idx++ {
		if curNode == nil {
			break
		}

		item := pathItems[idx]

		if item == "" {
			if idx == len(pathItems)-1 {
				item = "/"
			} else {
				continue
			}
		}

		nextNode := rt.find(curNode, item)

		if nextNode != nil {
			curNode = nextNode
			continue
		}

		if curNode.haveParam {
			nextNode = rt.findParamNode(curNode)
		}

		if nextNode != nil {
			params[nextNode.path] = item
		}

		curNode = nextNode
	}

	if curNode == nil || !curNode.isEnd {
		return nil, nil
	}

	return curNode.handler, params
}

func (rt *Tree) findParamNode(root *Node) *Node {
	for _, node := range root.children {
		if node.isParam {
			return node
		}
	}

	return nil
}

func (rt *Tree) find(root *Node, item string) *Node {
	if root == nil {
		return nil
	}

	if root.path == item {
		return root
	}

	for _, child := range root.children {
		found := rt.find(child, item)

		if found != nil {
			return found
		}
	}

	return nil
}
