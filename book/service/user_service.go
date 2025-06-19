package service

import (
	"context"
	"test/book/dal/model"
	"test/book/dal/query"
	"test/book/entity"
	"time"
)

type UserService struct {
	q *query.Query
}

func NewUserService() *UserService {
	return &UserService{q: query.Q}
}

func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	userDO := ToUserDO(user)
	return s.q.UsersDO.WithContext(ctx).Create(userDO)
}

func (s *UserService) UpdateUserStatus(ctx context.Context, userID int, status int32) error {
	_, err := s.q.UsersDO.WithContext(ctx).
		Where(s.q.UsersDO.ID.Eq(int64(userID))).
		Update(s.q.UsersDO.Status, status)
	return err
}

func (s *UserService) GetUserBooks(ctx context.Context, userID int) ([]*entity.Book, error) {
	usersDO := s.q.UsersDO
	uBooksDO := s.q.UserBooksDO
	user, err := usersDO.WithContext(ctx).Preload(usersDO.Books).
		Where(usersDO.ID.Eq(int64(userID))).
		Join(uBooksDO, usersDO.ID.EqCol(uBooksDO.UserID)).
		First()
	if err != nil {
		return nil, err
	}

	result := make([]*entity.Book, 0)
	for _, ub := range user.Books {
		book := &entity.Book{
			ID:          ub.ID,
			Title:       ub.Title,
			Author:      ub.Author,
			Price:       ub.Price,
			PublishDate: ub.PublishDate,
		}
		result = append(result, book)
	}
	return result, nil
}

func FromUserDO(user *model.UsersDO) *entity.User {
	return &entity.User{
		ID:         user.ID,
		Name:       user.Name,
		Password:   user.Password,
		Status:     user.Status,
		CreateTime: user.CreateTime,
	}
}

func ToUserDO(user *entity.User) *model.UsersDO {
	if user.CreateTime.IsZero() {
		user.CreateTime = time.Now()
	}
	return &model.UsersDO{
		ID:         user.ID,
		Name:       user.Name,
		Password:   user.Password,
		Status:     user.Status,
		CreateTime: user.CreateTime,
	}
}
