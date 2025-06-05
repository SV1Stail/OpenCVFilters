package server

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

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
	log.Debug().Msg("Start AddFiltersAndChannels")
	if r == nil || len(r.OriginalImage) == 0 {
		log.Err(constants.ErrBadRequest).Msg("No image")
		return nil, constants.ErrBadRequest
	}

	img, format, err := decodeAnyImage(r.OriginalImage)
	if err != nil {
		log.Err(err).Str("format", format).Msg("Cant decode")
		return nil, err
	}

	mat, err := gocv.ImageToMatRGB(img)
	if err != nil {
		log.Err(constants.ErrInternal).Msg("ImageToMatRGB")
		return nil, err
	}
	defer mat.Close()

	result := &FChResp{}

	err = result.splitChannels(mat, format)
	if err != nil {
		log.Err(constants.ErrInternal).Msg("splitChannels")
		return nil, err
	}

	err = result.addFilters(mat, format)
	if err != nil {
		log.Err(constants.ErrInternal).Msg("addFilters")
		return nil, err
	}

	return &result.FiltersAndChannelsResp, nil // проверить что реально возвращаются измененные поля структуры а не пустая хрень
}

// understand format of picture
func decodeAnyImage(data []byte) (img image.Image, format string, err error) {
	if img, err = jpeg.Decode(bytes.NewReader(data)); err == nil {
		log.Debug().Str("format", "jpeg").Msg("data")
		return img, "jpeg", nil
	}

	if img, err = png.Decode(bytes.NewReader(data)); err == nil {
		log.Debug().Str("format", "png").Msg("data")
		return img, "png", nil
	}

	return nil, "", fmt.Errorf("unsupported image format")
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
	gocv.GaussianBlur(mat, &gaussian, image.Point{X: 15, Y: 15}, 0, 0, gocv.BorderDefault)
	fc.FilteredImage1, err = convertMatToBytes(gaussian, format)
	if err != nil {
		return err
	}
	// Median Blur
	gocv.MedianBlur(mat, &median, 15)
	fc.FilteredImage2, err = convertMatToBytes(median, format)
	// Bilateral Filter
	gocv.BilateralFilter(mat, &bilateral, 25, 150, 150)
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
	err = gocv.Merge([]gocv.Mat{channels[0], channels[0], channels[2]}, &red)
	if err != nil {
		return constants.ErrInternal
	}
	fc.RedChannel, err = convertMatToBytes(red, format)
	if err != nil {
		return err
	}

	green := gocv.NewMat()
	defer green.Close()
	err = gocv.Merge([]gocv.Mat{channels[0], channels[2], channels[0]}, &green)
	if err != nil {
		return constants.ErrInternal
	}
	fc.GreenChannel, err = convertMatToBytes(green, format)
	if err != nil {
		return err
	}

	blue := gocv.NewMat()
	defer blue.Close()
	err = gocv.Merge([]gocv.Mat{channels[2], channels[0], channels[0]}, &blue)
	if err != nil {
		return constants.ErrInternal
	}
	fc.BlueChannel, err = convertMatToBytes(blue, format)
	if err != nil {
		return err
	}
	return nil
}
