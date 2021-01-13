package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"

	"github.com/regen-network/bec/x/blog"
)

var _ blog.MsgServer = serverImpl{}

func (s serverImpl) CreatePost(goCtx context.Context, request *blog.MsgCreatePostRequest) (*blog.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(s.storeKey), blog.KeyPrefix(blog.PostKey))

	id := uuid.New().String()
	post := blog.Post{
		Id:     id,
		Author: request.Author,
		Title:  request.Title,
		Body:   request.Body,
	}

	bz, err := s.cdc.MarshalBinaryBare(&post)
	if err != nil {
		return nil, err
	}

	store.Set([]byte(id), bz)

	return &blog.MsgCreatePostResponse{Id: id}, nil
}

func (s serverImpl) CreateComment(goCtx context.Context, request *blog.MsgCreateCommentRequest) (*blog.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(s.storeKey), blog.KeyPrefix(blog.CommentKey))
	id := uuid.New().String()

	comment := blog.Comment{
		Id:     id,
		Author: request.Author,
		PostID: request.PostID,
		Body:   request.Body,
	}

	bz, err := s.cdc.MarshalBinaryBare(&comment)
	if err != nil {
		return nil, err
	}

	store.Set([]byte(id), bz)

	return &blog.MsgCreateCommentResponse{Id: id}, nil

}
