package repository

import "GoDiary/src/data/models"

type DiaryRepo interface {
	Save(diary models.Diary) (models.Diary, error)
	FindById(id int) (models.Diary, error)
	count() float64
	delete(diary models.Diary)
	DeleteById(id int) error
	findDiaryByUsername(username string) (models.Diary, error)
	List() []models.Diary
}
