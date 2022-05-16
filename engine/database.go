package engine

import (
	"fmt"
)

// Content.
type Board struct {
	BoardID uint64
	Name    string

	Threads    []Thread
	LastPostID uint64
}

type Thread struct {
	Posts []Post
}

type Post struct {
	PostID uint64
	OP     bool

	Date  string
	Theme string
	Name  string

	Files   []File
	Comment string
}

type File struct {
	FileID  uint64
	Content []byte
	FileExt string
}

func (f File) String() string {
	return fmt.Sprintf("%d.%s", f.FileID, f.FileExt)
}

// Users.
const (
	ADMIN = iota
	MODER
	USER
)

// Ключ инвайта.
const KEY = "123"

type User struct {
	UID             uint64
	Login, Password string
	AccessType      uint8

	// ID досок, на которых юзер имеет права модератора.
	ModRights []uint64
}

type Base map[string]User

func (m Base) Has(key string) bool {
	_, ok := m[key]
	return ok
}

var (
	users = make(Base)
)
