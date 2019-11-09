package cerrors

import (
	"errors"
)

// 共通エラー common errors

var (
	// ErrNotFound なんかがみつかんなかったとき
	ErrNotFound = errors.New("Err not found")
)
