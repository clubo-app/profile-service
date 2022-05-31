package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	pg "github.com/clubo-app/protobuf/profile"
)

func (s *profileServer) GetManyProfiles(ctx context.Context, req *pg.GetManyProfilesRequest) (*pg.GetManyProfilesResponse, error) {
	ps, err := s.ps.GetMany(ctx, req.Ids, int32(len(req.Ids)))
	if err != nil {
		return nil, utils.HandleError(err)
	}

	profiles := make([]*pg.Profile, len(ps))
	for _, p := range ps {
		profiles = append(profiles, p.ToGRPCProfile())
	}

	return &pg.GetManyProfilesResponse{Profiles: profiles}, nil

}
