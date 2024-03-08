package models

type Diary struct {
	Username string
	password string
	Id       int
	isLocked bool
}

func MyDiary(username string, password string, id int) *Diary {
	return &Diary{
		Username: username,
		password: password,
		Id:       id,
		isLocked: true,
	}
}

func (receiver *Diary) SetUsername(username string) {
	receiver.Username = username
}

func (receiver *Diary) GetUsername() string {
	return receiver.Username

}

func (receiver *Diary) SetPassword(password string) {
	receiver.password = password

}

func (receiver *Diary) SetId(id int) {
	receiver.Id = id
}

func (receiver *Diary) GetId() int {
	return receiver.Id

}
