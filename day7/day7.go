package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Type int64

const (
	FOLDER Type = 0
	FILE        = 1
)

type TreeNode struct {
	parent   *TreeNode
	fileType Type
	path     string
	size     int
	files    map[string]*TreeNode
}

func initTreeNode(path string) *TreeNode {
	return &TreeNode{
		path:     path,
		fileType: FOLDER,
		files:    make(map[string]*TreeNode),
	}
}

func (t *TreeNode) insertFile(size int, path string) *TreeNode {
	return &TreeNode{
		parent:   t,
		fileType: FILE,
		size:     size,
		path:     path,
	}
}

func (t *TreeNode) insertFolder(path string) *TreeNode {
	return &TreeNode{
		parent: t,
		path:   path,
		files:  make(map[string]*TreeNode),
	}
}

func (t *TreeNode) cd(path string) *TreeNode {
	if path == ".." {
		return t.parent
	}

	return t.files[path]
}

func (t *TreeNode) handleCommand(command string) *TreeNode {
	split := strings.Split(command, " ")

	switch split[0] {
	case "$":
		if split[1] == "cd" {
			return t.cd(split[2])
		}

		if split[1] == "ls" {
			return t
		}
	default:
		return t.insert(command)
	}

	panic("Command not detected")
}

func (t *TreeNode) insert(data string) *TreeNode {
	split := strings.Split(data, " ")

	if split[0] == "dir" {
		t.files[split[1]] = t.insertFolder(split[1])
	} else {
		size, _ := strconv.Atoi(split[0])
		t.files[split[1]] = t.insertFile(size, split[1])
		t.updateSize(size)
	}

	return t
}

func (t *TreeNode) updateSize(size int) {
	t.size += size

	if t.parent != nil {
		t.parent.updateSize(size)
	}
}

func main() {
	input, err := ioutil.ReadFile("day7/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Day 7 [1]: ", partOne(string(input)))
	log.Println("Day 7 [2]: ", partTwo(string(input)))
}

var MAX_FOLDER_SIZE = 100000

func partOne(input string) int {
	lines := strings.Split(input, "\n")
	tree := initTreeNode("/")
	root := tree
	for i := 1; i < len(lines); i++ {
		tree = tree.handleCommand(lines[i])
	}

	return traverseThroughFolders(root, 0)
}

func traverseThroughFolders(tree *TreeNode, sum int) int {
	if tree.fileType == FOLDER {
		if tree.size <= MAX_FOLDER_SIZE {
			sum += tree.size
		}
	}

	if tree.files != nil {
		for _, file := range tree.files {
			if file.fileType == FOLDER {
				sum = traverseThroughFolders(file, sum)
			}
		}
	}

	return sum
}

var MAX_SIZE = 70000000
var UNUSED_SPACE = 30000000

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
	tree := initTreeNode("/")
	root := tree
	for i := 1; i < len(lines); i++ {
		tree = tree.handleCommand(lines[i])
	}

	actualRootSize := root.size
	actualUnusedSpace := MAX_SIZE - actualRootSize
	sizeNeededToDelete := UNUSED_SPACE - actualUnusedSpace

	return lookForSmallestDirToDelete(root, sizeNeededToDelete, root.size)
}

func lookForSmallestDirToDelete(tree *TreeNode, toDelete int, folderSize int) int {
	if tree.fileType == FOLDER {
		if tree.size >= toDelete && tree.size < folderSize {
			folderSize = tree.size
		}
	}

	if tree.files != nil {
		for _, file := range tree.files {
			if file.fileType == FOLDER {
				folderSize = lookForSmallestDirToDelete(file, toDelete, folderSize)
			}
		}
	}

	return folderSize
}
