package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	pg "github.com/clubo-app/protobuf/profile"
)

func (s *profileServer) GetManyProfilesMap(ctx context.Context, req *pg.GetManyProfilesRequest) (*pg.GetManyProfilesMapResponse, error) {
	ps, err := s.ps.GetMany(ctx, req.Ids, int32(len(req.Ids)))
	if err != nil {
		return nil, utils.HandleError(err)
	}

	profiles := make(map[string]*pg.Profile, len(ps))

	for _, p := range ps {
		profiles[p.ID] = p.ToGRPCProfile()
	}

	return &pg.GetManyProfilesMapResponse{Profiles: profiles}, nil

}
