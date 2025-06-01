package server

import (
	"bytes"
	"context"
	"image"

	"github.com/SV1Stail/OpenCVFilters/grpc/constants"
	"github.com/SV1Stail/OpenCVFilters/grpc/gen"
	"github.com/rs/zerolog/log"
	"gocv.io/x/gocv"
)

type Server struct {
	gen.UnimplementedServiceServer
}
type FChResp struct {
	gen.FiltersAndChannelsResp
}

func (s *Server) AddFiltersAndChannels(ctx context.Context, r *gen.ImageReq) (*gen.FiltersAndChannelsResp, error) {
	if r == nil || len(r.OriginalImage) == 0 {
		return nil, constants.ErrBadRequest
	}
	img, format, err := image.Decode(bytes.NewReader(r.OriginalImage))
	// img, err := jpeg.Decode(bytes.NewReader(r.OriginalImage))
	if err != nil {
		return nil, err
	}

	mat, err := gocv.ImageToMatRGB(img)
	if err != nil {
		return nil, err
	}
	defer mat.Close()

	result := &FChResp{}

	err = result.splitChannels(mat, format)
	if err != nil {
		return nil, err
	}

	err = result.splitChannels(mat, format)
	if err != nil {
		return nil, err
	}

	return &result.FiltersAndChannelsResp, nil // проверить что реально возвращаются измененные поля структуры а не пустая хрень
}

func (fc *FChResp) addFilters(mat gocv.Mat, format string) error {

	gaussian := gocv.NewMat()
	defer gaussian.Close()
	median := gocv.NewMat()
	defer median.Close()
	bilateral := gocv.NewMat()
	defer bilateral.Close()

	var err error
	// Gaussian Blur
	gocv.GaussianBlur(mat, &gaussian, image.Point{X: 5, Y: 5}, 0, 0, gocv.BorderDefault)
	fc.FilteredImage1, err = convertMatToBytes(gaussian, format)
	if err != nil {
		return err
	}
	// Median Blur
	gocv.MedianBlur(mat, &median, 5)
	fc.FilteredImage2, err = convertMatToBytes(median, format)
	// Bilateral Filter
	gocv.BilateralFilter(mat, &bilateral, 9, 75, 75)
	fc.FilteredImage3, err = convertMatToBytes(bilateral, format)

	return nil
}

func convertMatToBytes(mat gocv.Mat, format string) ([]byte, error) {
	var err error
	var buf *gocv.NativeByteBuffer
	switch format {
	case "jpeg":
		buf, err = gocv.IMEncode(gocv.JPEGFileExt, mat)
	case "png":
		buf, err = gocv.IMEncode(gocv.PNGFileExt, mat)

	default:
		return nil, constants.ErrWrongFormat
	}
	if err != nil {
		log.Err(err).Msg("Convert To Monochrome failed")
		return nil, err
	}

	return buf.GetBytes(), nil
}

// splitChannels splitting image (gocv.mat) on R G B channels
func (fc *FChResp) splitChannels(mat gocv.Mat, format string) error {
	var err error
	channels := gocv.Split(mat)
	red := gocv.NewMat()
	defer red.Close()
	err = gocv.Merge([]gocv.Mat{channels[2], channels[1], channels[0]}, &red)
	if err != nil {
		return constants.ErrInternal
	}
	fc.FiltersAndChannelsResp.RedChannel, err = convertMatToBytes(red, format)
	if err != nil {
		return err
	}

	green := gocv.NewMat()
	defer green.Close()
	err = gocv.Merge([]gocv.Mat{channels[0], channels[1], channels[2]}, &green)
	if err != nil {
		return constants.ErrInternal
	}
	fc.FiltersAndChannelsResp.GreenChannel, err = convertMatToBytes(green, format)
	if err != nil {
		return err
	}

	blue := gocv.NewMat()
	defer blue.Close()
	err = gocv.Merge([]gocv.Mat{channels[0], channels[1], channels[2]}, &blue)
	if err != nil {
		return constants.ErrInternal
	}
	fc.FiltersAndChannelsResp.BlueChannel, err = convertMatToBytes(blue, format)
	if err != nil {
		return err
	}
	return nil
}
