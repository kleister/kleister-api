package serverv1

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/model"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	"github.com/kleister/kleister-api/pkg/service/users/repository"
	users "github.com/kleister/kleister-api/pkg/service/users/v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var (
	noContext = context.Background()
)

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	want := []*types.User{
		{
			Id:        "dd5c7e72-4d42-4b9a-af85-45f792711e85",
			Slug:      "user1",
			Username:  "user1",
			Email:     "user1@example.com",
			Firstname: "Max",
			Lastname:  "Mustermann",
			Admin:     false,
			Active:    true,
			CreatedAt: timestamppb.New(time.Unix(1257894000, 0).UTC()),
			UpdatedAt: timestamppb.New(time.Unix(1257894000, 0).UTC()),
		},
	}

	repo := repository.NewMockUsersRepository(ctrl)
	repo.EXPECT().
		List(gomock.Any(), gomock.Any()).
		Return([]*model.User{
			{
				ID:        "dd5c7e72-4d42-4b9a-af85-45f792711e85",
				Slug:      "user1",
				Username:  "user1",
				Email:     "user1@example.com",
				Firstname: "Max",
				Lastname:  "Mustermann",
				Admin:     false,
				Active:    true,
				CreatedAt: time.Unix(1257894000, 0).UTC(),
				UpdatedAt: time.Unix(1257894000, 0).UTC(),
			},
		}, nil)

	server := NewUsersServer(
		config.Load(),
		nil,
		nil,
		repo,
	)

	got, err := server.List(
		noContext,
		&connect.Request[users.ListRequest]{
			Msg: &users.ListRequest{
				Query: "",
			},
		},
	)

	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(
		got.Msg.Users,
		want,
		cmpopts.IgnoreUnexported(
			types.User{},
			timestamppb.Timestamp{},
		),
	); diff != "" {
		t.Errorf(diff)
	}
}

func TestList_ErrGeneric(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	want := fmt.Errorf("whooops")

	repo := repository.NewMockUsersRepository(ctrl)
	repo.EXPECT().
		List(gomock.Any(), gomock.Any()).
		Return(nil, want)

	server := NewUsersServer(
		config.Load(),
		nil,
		nil,
		repo,
	)

	_, got := server.List(
		noContext,
		&connect.Request[users.ListRequest]{
			Msg: &users.ListRequest{
				Query: "",
			},
		},
	)

	if got.(*connect.Error).Unwrap() != want {
		t.Errorf("want: %s, got: %s", got, want)
	}
}
