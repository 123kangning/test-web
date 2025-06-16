package service

import (
	"context"
	"test/book/dal/model"
	"test/book/dal/query"
)

type UserService struct {
	q *query.Query
}

func NewUserService() *UserService {
	return &UserService{q: query.Q}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.UsersDO) error {
	return s.q.UsersDO.WithContext(ctx).Create(user)
}

func (s *UserService) UpdateUserStatus(ctx context.Context, userID int, status int32) error {
	_, err := s.q.UsersDO.WithContext(ctx).
		Where(s.q.UsersDO.ID.Eq(int64(userID))).
		Update(s.q.UsersDO.Status, status)
	return err
}

func (s *UserService) GetUserBooks(ctx context.Context, userID int) ([]*model.BooksDO, error) {
	usersDO := s.q.UsersDO
	uBooksDO := s.q.UserBooksDO
	user, err := usersDO.WithContext(ctx).Preload(usersDO.Books).
		Where(usersDO.ID.Eq(int64(userID))).
		Join(uBooksDO, usersDO.ID.EqCol(uBooksDO.UserID)).
		First()
	if err != nil {
		return nil, err
	}

	result := make([]*model.BooksDO, 0)
	for _, ub := range user.Books {
		book := &model.BooksDO{
			ID:     ub.ID,
			Title:  ub.Title,
			Author: ub.Author,
			Price:  ub.Price,
		}
		result = append(result, book)
	}
	return result, nil
}
