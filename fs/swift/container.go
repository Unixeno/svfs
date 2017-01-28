package swift

import (
	"time"

	"github.com/ovh/svfs/fs"
	"github.com/ovh/svfs/swift"
)

type Container struct {
	*Fs
	swiftContainer *swift.LogicalContainer
}

func (c *Container) Create(nodeName string) (fs.File, error) {
	panic("not implemented")
}

func (c *Container) GetAttr() (attr *fs.Attr, err error) {
	attr = &fs.Attr{
		Atime: time.Now(),
		Mtime: c.Fs.mountTime,
	}

	return
}

func (c *Container) Hardlink(targetPath string, linkName string) error {
	panic("not implemented")
}

func (c *Container) Mkdir(dirName string) (fs.Directory, error) {
	panic("not implemented")
}

func (c *Container) Remove(node fs.Node) error {
	panic("not implemented")
}

func (c *Container) Rename(newName string, newDir fs.Directory) error {
	panic("not implemented")
}

func (c *Container) Symlink(targetPath string, linkName string) error {
	panic("not implemented")
}
