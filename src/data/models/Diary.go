package models

type Diary struct {
	username string
	password string
	id       string
	isLocked bool
}

func myDiary(username string, password string, id string) *Diary {
	return &Diary{
		username: username,
		password: password,
		id:       id,
		isLocked: true,
	}
}

func (receiver *Diary) setUsername(username string) {
	receiver.username = username
}

func (receiver *Diary) getUsername() string {
	return receiver.username

}

func (receiver *Diary) setPassword(password string) {
	receiver.password = password

}

func (receiver *Diary) setId(id string) {
	receiver.id = id
}

func (receiver *Diary) getId() string {
	return receiver.id

}
