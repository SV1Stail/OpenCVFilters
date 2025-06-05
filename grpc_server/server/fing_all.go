package server

import (
	"context"

	"github.com/SV1Stail/OpenCVFilters/grpc/constants"
	"github.com/SV1Stail/OpenCVFilters/grpc/gen"
	"github.com/rs/zerolog/log"
	"gocv.io/x/gocv"
)

func (c *Server) FindAll(ctx context.Context, r *gen.ImageReq) (*gen.AllResp, error) {
	log.Debug().Msg("Start FindP")
	if r == nil || len(r.OriginalImage) == 0 {
		log.Err(constants.ErrBadRequest).Msg("No image")
		return nil, constants.ErrBadRequest
	}

	img, format, err := decodeAnyImage(r.OriginalImage)
	if err != nil {
		log.Err(err).Str("format", format).Msg("Cant decode")
		return nil, err
	}

	log.Debug().Msg("start ImageToMatRGB")
	mat, err := gocv.ImageToMatRGB(img)
	if err != nil {
		log.Err(constants.ErrInternal).Msg("ImageToMatRGB")
		return nil, err
	}
	defer mat.Close()

	log.Debug().Msg("start findContours")
	contours, err := findContours(mat)
	defer contours.Close()

	return &gen.AllResp{
		ResultP: findP(contours),
		ResultS: findS(contours),
	}, err
}
