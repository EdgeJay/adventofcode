package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/EdgeJay/adventofcode/common/utils/files"
	"github.com/EdgeJay/adventofcode/common/utils/mockfilesystem"
	"github.com/EdgeJay/adventofcode/common/utils/str"
)

const totalDiskSpace = 70000000
const unusedSpaceNeeded = 30000000

var unusedSpaceInDisk int

func isChangeDirCommand(line string) bool {
	return str.StartsWith(line, "$ "+mockfilesystem.CommandChangeDirectory)
}

func isListDirCommand(line string) bool {
	return str.StartsWith(line, "$ "+mockfilesystem.CommandList)
}

func buildFolder(root *mockfilesystem.Folder, input []string) {
	if len(input) > 0 {
		command := input[0]
		if isChangeDirCommand(command) {
			folderName := command[5:]
			if root.GetName() == "" {
				// root folder "/"
				root.SetName(folderName)
				buildFolder(root, input[1:])
			} else if folderName == ".." {
				// go up one level
				buildFolder(root.GetParent(), input[1:])
			} else {
				// go into folder
				folder := root.FindObject(folderName).(*mockfilesystem.Folder)
				buildFolder(folder, input[1:])
			}
		} else if isListDirCommand(command) {
			buildFolder(root, input[1:])
		} else {
			// folder directory items
			cmdParts := strings.Split(command, " ")
			if cmdParts[0] == "dir" {
				// folder
				dir := mockfilesystem.NewFolder(cmdParts[1])
				dir.SetParent(root)
				root.AddObject(dir)
			} else {
				// file
				size, _ := strconv.Atoi(cmdParts[0])
				file := mockfilesystem.NewFile(cmdParts[1], size)
				file.SetParent(root)
				root.AddObject(file)
			}
			buildFolder(root, input[1:])
		}
	}
}

func buildFileSystemFromCommands(lines []string) *mockfilesystem.Folder {
	root := mockfilesystem.NewFolder("")
	buildFolder(root, lines)
	return root
}

func getFolderSize(root *mockfilesystem.Folder, folders *[]*mockfilesystem.Folder) {
	size := root.CalculateTotalSize()
	if root.Name == "/" {
		unusedSpaceInDisk = totalDiskSpace - size
		fmt.Println("Total disk space:", totalDiskSpace)
		fmt.Println("Space used by /:", size)
		fmt.Println("Unused space in disk:", unusedSpaceInDisk)
	} else {
		if size+unusedSpaceInDisk >= unusedSpaceNeeded {
			*folders = append(*folders, root)
		}
	}

	for _, child := range root.Children {
		if child.GetObjectType() == mockfilesystem.ObjectTypeFolder {
			getFolderSize((child).(*mockfilesystem.Folder), folders)
		}
	}
}

func printResults(folders *[]*mockfilesystem.Folder) {
	sort.Slice(*folders, func(i, j int) bool {
		return (*folders)[i].TotalSize < (*folders)[j].TotalSize
	})
	for _, folder := range *folders {
		fmt.Println(folder.Name, folder.TotalSize)
	}
}

func calculate(root *mockfilesystem.Folder) {
	folders := make([]*mockfilesystem.Folder, 0)
	getFolderSize(root, &folders)
	printResults(&folders)
}

func main() {
	lines := files.ReadInputsFile("./input.txt")
	// lines := files.ReadInputsFile("./test.txt")
	fs := buildFileSystemFromCommands(lines)
	calculate(fs)
}
