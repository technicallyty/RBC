package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// CreateComment - implements commenting functionality. Returns comsos sdk error if a post with given ID does not exist.
func (s serverImpl) CreateComment(goCtx context.Context, request *blog.MsgCreateCommentRequest) (*blog.MsgCreateCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(s.storeKey), blog.KeyPrefix(blog.CommentKey))
	id := uuid.New().String()

	/*
	* The below error doesn't propagate throughout the system
	* Here, it is aware when a post doesn't exist,
	* however, the error is not picked up by the test case.
	* Not really sure where it ends up.
	 */

	storePosts := prefix.NewStore(ctx.KVStore(s.storeKey), blog.KeyPrefix(blog.PostKey))

	// check if store has post with given ID
	ok := storePosts.Has([]byte(request.PostID))

	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Post with id %v does not exist", request.PostID)
	}

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
