package mockfilesystem

const (
	CommandChangeDirectory = "cd"
	CommandList            = "ls"
)

const (
	Directory = "dir"
)

const (
	ObjectTypeFolder = "FOLDER"
	ObjectTypeFile   = "FILE"
)

type FileSystemObject interface {
	GetName() string
	SetName(string)
	GetObjectType() string
	GetParent() *Folder
	SetParent(*Folder)
}

type Folder struct {
	Name       string
	Children   []FileSystemObject
	objectType string
	parent     *Folder
	TotalSize  int
}

func NewFolder(name string) *Folder {
	return &Folder{name, []FileSystemObject{}, ObjectTypeFolder, nil, 0}
}

func (f *Folder) AddObject(obj FileSystemObject) *Folder {
	f.Children = append(f.Children, obj)
	return f
}

func (f *Folder) FindObject(name string) FileSystemObject {
	for _, obj := range f.Children {
		if obj.GetName() == name {
			return obj
		}
	}
	return nil
}

func (f *Folder) GetName() string {
	return f.Name
}

func (f *Folder) SetName(name string) {
	f.Name = name
}

func (f *Folder) GetObjectType() string {
	return f.objectType
}

func (f *Folder) GetParent() *Folder {
	return f.parent
}

func (f *Folder) SetParent(parent *Folder) {
	f.parent = parent
}

func (f *Folder) CalculateTotalSize() int {
	f.TotalSize = 0
	for _, child := range f.Children {
		if child.GetObjectType() == ObjectTypeFolder {
			f.TotalSize += (child).(*Folder).CalculateTotalSize()
		} else if child.GetObjectType() == ObjectTypeFile {
			f.TotalSize += (child).(*File).Size
		}
	}
	return f.TotalSize
}

type File struct {
	Name       string
	Size       int
	objectType string
	parent     *Folder
}

func (f *File) GetName() string {
	return f.Name
}

func (f *File) SetName(name string) {
	f.Name = name
}

func (f *File) GetObjectType() string {
	return f.objectType
}

func (f *File) GetParent() *Folder {
	return f.parent
}

func (f *File) SetParent(parent *Folder) {
	f.parent = parent
}

func NewFile(name string, size int) *File {
	return &File{name, size, ObjectTypeFile, nil}
}
