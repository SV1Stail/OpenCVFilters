package server

import (
	"context"
	"image/color"

	"github.com/SV1Stail/OpenCVFilters/grpc/constants"
	"github.com/SV1Stail/OpenCVFilters/grpc/gen"
	"github.com/rs/zerolog/log"
	"gocv.io/x/gocv"
)

type FCResp struct {
	gen.FindContoursResp
}

func (s *Server) FindContours(ctx context.Context, r *gen.ImageReq) (*gen.FindContoursResp, error) {
	log.Debug().Msg("Start FindContours")
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

	blackBackground := gocv.NewMatWithSize(mat.Rows(), mat.Cols(), gocv.MatTypeCV8UC3)
	defer blackBackground.Close()

	if contours.Size() == 0 {
		log.Warn().Msg("No contours found")
		finalImageBytes, err := convertMatToBytes(mat, format)
		if err != nil {
			return nil, err
		}
		return &gen.FindContoursResp{FinalImageData: finalImageBytes}, nil
	}

	for i := 0; i < contours.Size(); i++ {
		gocv.DrawContours(
			&blackBackground,
			contours,
			i,
			color.RGBA{0, 255, 0, 255},
			2,
		)
	}

	resp := &FCResp{}

	log.Debug().Msg("start convertMatToBytes")
	resp.FinalImageData, err = convertMatToBytes(blackBackground, format)
	if err != nil {
		log.Err(constants.ErrInternal).Msg("convertMatToBytes")
		return nil, err
	}

	log.Debug().Int("resp image len", len(resp.FinalImageData)).Msg("result")
	return &resp.FindContoursResp, nil
}

func findContours(mat gocv.Mat) (gocv.PointsVector, error) {

	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(mat, &gray, gocv.ColorBGRToGray)

	binary := gocv.NewMat()
	defer binary.Close()
	gocv.Threshold(gray, &binary, 128, 255, gocv.ThresholdBinary)

	contours := gocv.FindContours(binary, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	defer contours.Close()

	return contours, nil
}
