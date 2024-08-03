// Copyright (C) 2024 Tim Bastin, l3montree UG (haftungsbeschränkt)
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package asset

import (
	"github.com/l3montree-dev/devguard/internal/database/models"
)

type treeNode struct {
	Name     string      `json:"name"`
	Children []*treeNode `json:"children"`
}

type tree struct {
	Root           *treeNode `json:"root"`
	cursors        map[string]*treeNode
	insertionOrder []string
}

func newNode(name string) *treeNode {
	return &treeNode{
		Name:     name,
		Children: []*treeNode{},
	}
}

func (tree *tree) addNode(source string, dep string) {
	// check if source does exist
	if _, ok := tree.cursors[source]; !ok {
		tree.cursors[source] = newNode(source)
		tree.insertionOrder = append(tree.insertionOrder, source)
	}
	// check if dep does already exist
	if _, ok := tree.cursors[dep]; !ok {
		tree.cursors[dep] = newNode(dep)
		tree.insertionOrder = append(tree.insertionOrder, source)
	}

	// check if connection does already exist
	for _, child := range tree.cursors[source].Children {
		if child.Name == dep {
			return
		}
	}

	tree.cursors[source].Children = append(tree.cursors[source].Children, tree.cursors[dep])
}

// Helper function to detect and cut cycles
func cutCycles(node *treeNode, visited map[*treeNode]bool) {
	// Mark the current node as visited
	visited[node] = true

	// Iterate over the children
	for i := 0; i < len(node.Children); i++ {
		child := node.Children[i]
		if visited[child] {
			// If the child is already visited, we have found a cycle
			// Remove the child reference to cut the cycle
			node.Children = append(node.Children[:i], node.Children[i+1:]...)
			i-- // Adjust index due to slice modification
		} else {
			// Recursively check the child
			cutCycles(child, visited)
		}
	}

	// Unmark the current node before returning to allow different paths
	// to explore this node without falsely detecting a cycle
	delete(visited, node)
}

func CalculateDepth(node *treeNode, currentDepth int, depthMap map[string]int) {
	if _, ok := depthMap[node.Name]; !ok {
		depthMap[node.Name] = currentDepth
	} else if depthMap[node.Name] > currentDepth {
		// use the shortest path
		depthMap[node.Name] = currentDepth
	}
	for _, child := range node.Children {
		CalculateDepth(child, currentDepth+1, depthMap)
	}
}

func BuildDependencyTree(elements []models.ComponentDependency) tree {
	// create a new tree
	tree := tree{
		Root:           &treeNode{Name: "root"},
		cursors:        make(map[string]*treeNode),
		insertionOrder: make([]string, 0),
	}

	tree.cursors["root"] = tree.Root

	for _, element := range elements {
		if element.ComponentPurl == nil {
			tree.addNode("root", element.DependencyPurl)
		} else {
			tree.addNode(*element.ComponentPurl, element.DependencyPurl)
		}
	}

	cutCycles(tree.Root, make(map[*treeNode]bool))

	return tree
}
