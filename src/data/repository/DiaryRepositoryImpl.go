package repository

import "GoDiary/src/data/models"

type DiaryRepoImpl struct {
	Diary map[int]models.Diary
}

func newDiaryRepoImpl() *DiaryRepoImpl {
	return &DiaryRepoImpl{
		Diary: make(map[int]models.Diary),
	}
}
