package model

// FSTask
type FSTask struct {
	ID      string `firestore:"ID"`
	IsDone  bool   `firestore:"IsDone"`
	Name    string `firestore:"Name"`
	OwnerID string `firestore:"OwnerID"`
}

// type FSTaskList struct {
// 	Tasks []FSTask
// }
