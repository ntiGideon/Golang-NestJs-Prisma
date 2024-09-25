package post

import (
	"NestJsStyle/data"
	"NestJsStyle/helper"
	"NestJsStyle/model"
	"NestJsStyle/prisma/db"
	"context"
	"fmt"
	"net/http"
	"os"
)

type PrismaInjection struct {
	Db *db.PrismaClient
}

func NewPrismaInjection(db *db.PrismaClient) PrismaInjection {
	return PrismaInjection{
		Db: db,
	}
}

func (p PrismaInjection) CreatePostService(ctx context.Context, post model.Post) *data.WebResponse {

	err := helper.RequestValidators(post)
	if err != nil {
		return &data.WebResponse{
			Data:   err.Error(),
			Code:   http.StatusBadRequest,
			Status: "Validation errors",
		}
	}

	exitingPostTitle, err := p.Db.Post.FindFirst(db.Post.Title.Equals(post.Title)).Exec(ctx)
	if exitingPostTitle != nil && err == nil {
		return &data.WebResponse{
			Data:   nil,
			Code:   http.StatusBadRequest,
			Status: "Post title already in use!",
		}
	}
	_, err = p.Db.Post.CreateOne(
		db.Post.Title.Set(post.Title),
		db.Post.Description.Set(post.Description),
		db.Post.Published.Set(post.Published),
		db.Post.User.Link(
			db.User.ID.Set(post.UserId),
		),
	).Exec(ctx)
	if err != nil {
		helper.PanicAllErrors(err)
	}

	return &data.WebResponse{
		Data:   nil,
		Code:   http.StatusCreated,
		Status: "OK",
	}
}

func (p PrismaInjection) UpdatePost(ctx context.Context, post model.PostUpdate) *data.WebResponse {
	err := helper.RequestValidators(post)
	if err != nil {
		return &data.WebResponse{
			Data:   err.Error(),
			Code:   http.StatusBadRequest,
			Status: "Validation errors",
		}
	}

	_, err = p.Db.Post.FindUnique(db.Post.ID.Equals(post.Id)).Update(
		db.Post.Title.Set(post.Title),
		db.Post.Published.Set(post.Published),
		db.Post.Description.Set(post.Description),
		db.Post.User.Link(
			db.User.ID.Set(post.UserId),
		),
	).Exec(ctx)
	if err != nil {
		helper.PanicAllErrors(err)
	}
	return &data.WebResponse{
		Data:   nil,
		Code:   http.StatusOK,
		Status: "Updated",
	}
}

func (p PrismaInjection) DeletePost(ctx context.Context, postId string) *data.WebResponse {
	existingPost, err := p.Db.Post.FindFirst(db.Post.ID.Equals(postId)).Exec(ctx)
	if existingPost == nil && err != nil {
		return &data.WebResponse{
			Data:   nil,
			Code:   http.StatusNotFound,
			Status: "Post not found",
		}
	}
	_, err = p.Db.Post.FindUnique(db.Post.ID.Equals(postId)).Delete().Exec(ctx)
	if err != nil {
		helper.PanicAllErrors(err)
	}
	return &data.WebResponse{
		Data:   nil,
		Code:   http.StatusOK,
		Status: "Deleted",
	}
}

func (p *PrismaInjection) GetAllPost(ctx context.Context, page int, limit int) *data.WebResponsePagination {

	totalCount, err := p.Db.Post.FindMany().Exec(ctx)

	posts, err := p.Db.Post.FindMany().Select(
		db.Post.ID.Field(),
		db.Post.Title.Field(),
		db.Post.Published.Field(),
		db.Post.Description.Field(),
		db.Post.CreatedAt.Field(),
	).Take(limit).Skip((page - 1) * limit).OrderBy(db.Post.CreatedAt.Order(db.SortOrderDesc)).Exec(ctx)

	if err != nil {
		helper.PanicAllErrors(err)
	}

	totalPages := (len(totalCount) + limit - 1) / limit

	route := fmt.Sprintf("%v/api/post", os.Getenv("FRONTEND_URL"))
	first := fmt.Sprintf("%v?limit=%v&page=1", route, limit)
	last := fmt.Sprintf("%v?limit=%v&page=%v", route, limit, totalPages)

	var prev, next string
	if page > 1 {
		prev = fmt.Sprintf("%v?limit=%v&page=%v", route, limit, page-1)
	}
	if page < totalPages {
		next = fmt.Sprintf("%v?limit=%v&page=%v", route, limit, page+1)
	}

	return &data.WebResponsePagination{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   posts,
		Meta: &data.Meta{
			CurrentPage:  page,
			ItemsPerPage: limit,
			ItemCount:    len(posts),
			TotalCount:   len(totalCount),
			TotalPages:   totalPages,
		},
		Links: &data.Links{
			First:    first,
			Previous: prev,
			Next:     next,
			Last:     last,
		},
	}
}

func (p *PrismaInjection) FindPostById(ctx context.Context, postId string) *data.WebResponse {
	post, err := p.Db.Post.FindFirst(db.Post.ID.Equals(postId)).Exec(ctx)
	if err != nil {
		return &data.WebResponse{
			Data:   nil,
			Code:   http.StatusNotFound,
			Status: "Post not found",
		}
	}
	if post == nil {
		return &data.WebResponse{
			Code:   http.StatusNotFound,
			Status: "No Content",
			Data:   nil,
		}
	}

	return &data.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: &model.PostResponse{
			Id:          post.UserID,
			Title:       post.Title,
			Description: post.Description,
		},
	}
}
