package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
)

type File struct {
	Name      string
	Parent    *File
	Children  []*File
	Size      int
	totalSize int
}

func (f *File) TotalSize() int {
	if !f.IsDir() {
		return f.Size
	}

	if f.totalSize > 0 {
		return f.totalSize
	}

	total := 0
	for _, c := range f.Children {
		total += c.TotalSize()
	}
	f.totalSize = total

	return f.totalSize
}

func (f *File) IsDir() bool {
	return f.Size == 0
}

func (f *File) VisitSubDirs(action func(sz int)) {
	if !f.IsDir() {
		return
	}

	action(f.TotalSize())
	for _, c := range f.Children {
		c.VisitSubDirs(action)
	}
}

func parseFileSystem(input []string) (root *File) {
	root = &File{Name: "/"}
	cwd := root

	for _, line := range input {
		if line == "" || line == "$ ls" {
			continue
		}

		var next string
		if _, err := fmt.Sscanf(line, "$ cd %s", &next); err == nil {
			switch next {
			case "/":
				cwd = root
			case "..":
				cwd = cwd.Parent
			default:
				for _, f := range cwd.Children {
					if f.Name == next {
						cwd = f
						break
					}
				}
			}

			continue
		}

		if _, err := fmt.Sscanf(line, "dir %s", &next); err == nil {
			d := &File{Parent: cwd, Name: next}
			cwd.Children = append(cwd.Children, d)
			continue
		}

		f := &File{Parent: cwd}
		_ = utils.Must(fmt.Sscanf(line, "%d %s", &f.Size, &f.Name))
		cwd.Children = append(cwd.Children, f)
	}

	return root
}

func solveA(input []string) int {
	root := parseFileSystem(input)

	total := 0
	root.VisitSubDirs(func(sz int) {
		if sz <= 100000 {
			total += sz
		}
	})

	return total
}

func solveB(input []string) int {
	root := parseFileSystem(input)

	sizes := make([]int, 0)
	root.VisitSubDirs(func(sz int) {
		sizes = append(sizes, sz)
	})

	needed := root.TotalSize() - 40000000
	max := 70000000

	for _, sz := range sizes {
		if sz >= needed && sz < max {
			max = sz
		}
	}

	return max
}

func main() {
	problem := utils.Must(utils.NewProblem(2022, 7))
	input := utils.Must(problem.LoadInput())

	var result int
	if problem.IsA() {
		result = solveA(input)
	}

	if problem.IsB() {
		result = solveB(input)
	}

	_ = problem.WriteOutput(strconv.Itoa(result))
}
