package repository

import (
	"GoDiary/src/data/models"
	"errors"
	"fmt"
)

type DiaryRepoImpl struct {
	Diary map[int]models.Diary
}

func NewDiaryRepoImpl() *DiaryRepoImpl {
	return &DiaryRepoImpl{
		Diary: make(map[int]models.Diary),
	}
}

func (repo *DiaryRepoImpl) Save(diary models.Diary) (models.Diary, error) {
	if diary.Id == 0 {
		diary.Id = generateUniqueId()
	}
	repo.Diary[diary.Id] = diary

	return diary, nil
}

func (repo *DiaryRepoImpl) FindById(id int) (models.Diary, error) {
	diary, exists := repo.Diary[id]
	if !exists {
		return models.Diary{}, errors.New("user not found")
	}
	return diary, nil
}

func (repo *DiaryRepoImpl) DeleteById(id int) error {
	_, exists := repo.Diary[id]

	if !exists {
		return errors.New("user not found")
	}

	delete(repo.Diary, id)
	return nil
}

func (repo *DiaryRepoImpl) List() []models.Diary {
	diaries := make([]models.Diary, 0, len(repo.Diary))
	for _, diary := range repo.Diary {
		diaries = append(diaries, diary)
	}
	return diaries
}

func (repo *DiaryRepoImpl) findDiaryByUsername(username string) (models.Diary, error) {
	for _, diary := range repo.Diary {
		if diary.Username == username {
			return models.Diary{}, nil
		}
	}

	return models.Diary{}, fmt.Errorf("user with username '%s' not found", username)
}

func generateUniqueId() int {
	staticIdCounter++
	return staticIdCounter
}

var staticIdCounter int
