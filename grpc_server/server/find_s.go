package server

import (
	"context"

	"github.com/SV1Stail/OpenCVFilters/grpc/constants"
	"github.com/SV1Stail/OpenCVFilters/grpc/gen"
	"github.com/rs/zerolog/log"
	"gocv.io/x/gocv"
)

func (s *Server) FindP(ctx context.Context, r *gen.ImageReq) (*gen.NumericalResp, error) {
	log.Debug().Msg("Start FindS")
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

	return &gen.NumericalResp{Result: findS(contours)}, nil
}

func findS(contours gocv.PointsVector) float64 {
	var areas float64

	for i := 0; i < contours.Size(); i++ {
		contour := contours.At(i)
		areas += gocv.ContourArea(contour)
	}
	return areas
}
