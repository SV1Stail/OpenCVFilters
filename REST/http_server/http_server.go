package httpserver

import (
	"context"
	"io"
	"net/http"
	"strconv"

	"github.com/SV1Stail/OpenCVFilters/REST/constants"
	"github.com/SV1Stail/OpenCVFilters/REST/gen"
	"github.com/rs/zerolog/log"
)

type Client struct {
	gen.ServiceClient
}

func NewClient(client gen.ServiceClient) *Client {
	return &Client{client}
}

func (c *Client) UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Err(constants.ErrNotAllowed).Msg("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		log.Err(constants.ErrBadRequest).Msg("Failed to read image")
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imgBytes, err := io.ReadAll(file)
	if err != nil {
		log.Err(constants.ErrInternal).Msg("Failed to read image data")
		http.Error(w, "Failed to read image data", http.StatusInternalServerError)
		return
	}

	var resp *pb.ImageResponse
	switch r.FormValue("type") {
	case "binary":
		resp, err = grpcClient.ConvertToBinary(context.Background(), &pb.BinaryRequest{
			ImageData: imgBytes,
		})
	case "monochrome":
		resp, err = grpcClient.ConvertToMonochrome(context.Background(), &pb.MonochromeRequest{
			ImageData:   imgBytes,
			TargetColor: r.FormValue("color"),
		})
	case "threshold":
		var thresholdInt int

		thresholdInt, err = strconv.Atoi(r.FormValue("threshold"))
		if err != nil {
			log.Err(constants.ErrBadRequest).Msg("Invalid threshold value")
			http.Error(w, "Invalid threshold value", http.StatusBadRequest)
			return
		}

		resp, err = grpcClient.ConvertToThreshold(context.Background(), &pb.ThresholdRequest{
			ImageData: imgBytes,
			Threshold: int32(thresholdInt),
		})
	}
	if err != nil {
		log.Err(err).Msg("gRPC call failed")
		http.Error(w, "gRPC call failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(resp.ProcessedImageData)
}
