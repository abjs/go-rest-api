package commend

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorsLoadingComments = errors.New("Error loading comment by CommentId")
	ErrorNotImplemented   = errors.New("The method is not implemented as of now")
)

type Commend struct {
	CommendId string
	PostId    string
	Body      string
	Author    string
}

type CommentStore interface {
	GetComment(context.Context, string) (Commend, error)
}

type Service struct {
	CommentStore CommentStore
}

func NewService(commentStore CommentStore) *Service {
	return &Service{
		CommentStore: commentStore,
	}
}

func (s *Service) GetComment(ctx context.Context, CommendId string) (Commend, error) {
	fmt.Println("Loading The Comments...")
	cmt, err := s.CommentStore.GetComment(ctx, CommendId)
	if err != nil {
		fmt.Println(err)
		return Commend{}, ErrorsLoadingComments
	}
	return cmt, nil
}

func (s *Service) CreateComments(ctx context.Context, cmt Commend) error {
	return ErrorNotImplemented
}

func (s *Service) UpdateComments(ctx context.Context, cmt Commend) error {
	return ErrorNotImplemented
}

func (s *Service) DeleteComments(ctx context.Context, CommendId string) (Commend, error) {
	return Commend{}, ErrorNotImplemented
}
