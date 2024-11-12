package service

import (
	"DB_course_paper/server/entity"
	"context"
)

func (s service) CreateTokens(ctx context.Context, token entity.Token) (int, error) {
	tokenId, err := s.r.CreateTokens(ctx, token)
	if err != nil {
		return 0, err
	}
	return tokenId, nil
}

func (s service) DeleteTokens(ctx context.Context, tokenId int) error {
	err := s.r.DeleteTokens(ctx, tokenId)
	if err != nil {
		return err
	}
	return nil
}

func (s service) UpdateTokens(ctx context.Context, token entity.Token) (int, error) {
	tokenId, err := s.r.UpdateTokens(ctx, token)
	if err != nil {
		return 0, err
	}
	return tokenId, nil
}

func (s service) GetTokens(ctx context.Context, tokenId int) (entity.Token, error) {
	tokens, err := s.r.GetTokens(ctx, tokenId)
	if err != nil {
		return entity.Token{}, err
	}
	return tokens, nil
}

func (s service) LogOutUser(ctx context.Context, userId int) error {
	err := s.r.LogOutUser(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}
