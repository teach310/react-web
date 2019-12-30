package model

// Task  gormに保存するよう
type Task struct {
	ID      string `gorm:"primary_key;type:varchar(64)"`
	IsDone  bool
	Name    string
	OwnerID string `gorm:"index:idx_task_01;type:varchar(64)"` // 所有者
}
