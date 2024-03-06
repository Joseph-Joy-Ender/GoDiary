package repository

import "GoDiary/src/data/models"

type DiaryRepo interface {
	save(diary models.Diary) models.Diary
	findById(id string)
	count() float64
	delete(diary models.Diary)
	deleteById(id string)
	findDiaryByUsername(username string) models.Diary
	list() []models.Diary
}
