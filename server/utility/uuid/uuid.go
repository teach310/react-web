package uuid

import (
	guuid "github.com/google/uuid"
)

var newString = DefaultNewString

// NewString uuidをstringで生成
func NewString() string {
	return newString()
}

// DefaultNewString ライブラリで生成したUUIDの文字列を返す
func DefaultNewString() string {
	newUUID := guuid.New()
	return newUUID.String()
}
