package structural

import "fmt"

// FileSystemFunc defines the common behavior for files and directories
type FileSystemFunc interface {
	ShowStructure(prefix string)
}

// FileSystem represents the common properties for files and directories
type FileSystem struct {
	name string
}

// FileDetail represents a file in the file system
type FileDetail struct {
	fileSystem FileSystem
}

func NewFileDetail(name string) *FileDetail {
	return &FileDetail{
		fileSystem: FileSystem{
			name: name,
		},
	}
}

func (f *FileDetail) ShowStructure(prefix string) {
	fmt.Printf("%s/%s\n", prefix, f.fileSystem.name)
}

// Directory represents a directory in the file system
type Directory struct {
	fileSystem FileSystem
	children   []FileSystemFunc
}

func NewDirectory(name string) *Directory {
	return &Directory{
		fileSystem: FileSystem{
			name: name,
		},
	}
}

func (d *Directory) AddComponent(component FileSystemFunc) {
	d.children = append(d.children, component)
}

func (d *Directory) ShowStructure(prefix string) {
	fmt.Printf("%s/%s\n", prefix, d.fileSystem.name)
	for _, child := range d.children {
		if fileDetail, ok := child.(*FileDetail); ok {
			fileDetail.ShowStructure(fmt.Sprintf("%s/%s", prefix, d.fileSystem.name))
		} else if directory, ok := child.(*Directory); ok {
			directory.ShowStructure(fmt.Sprintf("%s/%s", prefix, d.fileSystem.name))
		}
	}
}

func RunComposite() {
	file1 := NewFileDetail("file1.txt")
	file2 := NewFileDetail("file2.txt")
	file3 := NewFileDetail("file3.txt")

	rootDir := NewDirectory("root")
	subDir1 := NewDirectory("sub_dir_1")
	subDir2 := NewDirectory("sub_dir_2")

	subDir1.AddComponent(file1)
	subDir2.AddComponent(file2)
	subDir2.AddComponent(file3)

	rootDir.AddComponent(subDir1)
	rootDir.AddComponent(subDir2)

	rootDir.ShowStructure("")
}
