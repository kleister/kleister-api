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
	"github.com/kleister/kleister-api/pkg/service/teams/repository"
	teams "github.com/kleister/kleister-api/pkg/service/teams/v1"
	types "github.com/kleister/kleister-api/pkg/service/types/v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var (
	noContext = context.Background()
)

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	want := []*types.Team{
		{
			Id:        "dd5c7e72-4d42-4b9a-af85-45f792711e85",
			Slug:      "team-1",
			Name:      "Team 1",
			CreatedAt: timestamppb.New(time.Unix(1257894000, 0).UTC()),
			UpdatedAt: timestamppb.New(time.Unix(1257894000, 0).UTC()),
		},
	}

	repo := repository.NewMockTeamsRepository(ctrl)
	repo.EXPECT().
		List(gomock.Any(), "").
		Return([]*model.Team{
			{
				ID:        "dd5c7e72-4d42-4b9a-af85-45f792711e85",
				Slug:      "team-1",
				Name:      "Team 1",
				CreatedAt: time.Unix(1257894000, 0).UTC(),
				UpdatedAt: time.Unix(1257894000, 0).UTC(),
			},
		}, nil)

	server := NewTeamsServer(
		config.Load(),
		nil,
		nil,
		repo,
	)

	got, err := server.List(
		noContext,
		&connect.Request[teams.ListRequest]{
			Msg: &teams.ListRequest{
				Query: "",
			},
		},
	)

	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(
		got.Msg.Teams,
		want,
		cmpopts.IgnoreUnexported(
			types.Team{},
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

	repo := repository.NewMockTeamsRepository(ctrl)
	repo.EXPECT().
		List(gomock.Any(), "").
		Return(nil, want)

	server := NewTeamsServer(
		config.Load(),
		nil,
		nil,
		repo,
	)

	_, got := server.List(
		noContext,
		&connect.Request[teams.ListRequest]{
			Msg: &teams.ListRequest{
				Query: "",
			},
		},
	)

	if got.(*connect.Error).Unwrap() != want {
		t.Errorf("want: %s, got: %s", got, want)
	}
}
