package models

type Diary struct {
	username string
	password string
	id       string
	isLocked bool
}

func myDiary(username string, password string, id string, isLocked bool) *Diary {
	return &Diary{
		username: username,
		password: password,
		id:       id,
		isLocked: true,
	}
}
