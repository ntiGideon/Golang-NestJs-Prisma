package post

import (
	"NestJsStyle/helper"
	"NestJsStyle/prisma/db"
	"context"
)

type PostRepository struct {
	Db *db.PrismaClient
}

func NewPostRepository(db *db.PrismaClient) *PostRepository {
	return &PostRepository{Db: db}
}

func (p PostRepository) ExistingPostByTitle(ctx context.Context, title string) bool {
	existingPost, err := p.Db.Post.FindFirst(db.Post.Title.Equals(title)).Exec(ctx)
	if err != nil {
		helper.PanicAllErrors(err)
	}
	if existingPost == nil {
		return false
	}
	return true
}
