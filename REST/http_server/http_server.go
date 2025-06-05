package httpserver

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"

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

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	switch r.FormValue("operation") {
	case "AddFiltersAndChannels":
		log.Info().Str("method", "AddFiltersAndChannels").Msg("gRPC call begins")
		resp, err := c.AddFiltersAndChannels(ctx, &gen.ImageReq{
			OriginalImage: imgBytes,
		})
		if err != nil {
			log.Err(err).Str("method", "AddFiltersAndChannels").Msg("gRPC call failed")
			http.Error(w, "gRPC call failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		result := map[string]interface{}{
			"redChannel":     base64.StdEncoding.EncodeToString(resp.RedChannel),
			"greenChannel":   base64.StdEncoding.EncodeToString(resp.GreenChannel),
			"blueChannel":    base64.StdEncoding.EncodeToString(resp.BlueChannel),
			"filteredImage3": base64.StdEncoding.EncodeToString(resp.FilteredImage3),
			"filteredImage1": base64.StdEncoding.EncodeToString(resp.FilteredImage1),
			"filteredImage2": base64.StdEncoding.EncodeToString(resp.FilteredImage2),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)

	case "FindContours":
		log.Info().Str("method", "FindContours").Msg("gRPC call begins")
		resp, err := c.FindContours(ctx, &gen.ImageReq{
			OriginalImage: imgBytes,
		})
		if err != nil {
			log.Err(err).Str("method", "FindContours").Msg("gRPC call failed")
			http.Error(w, "gRPC call failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		result := map[string]interface{}{
			"FindContoursImage": base64.StdEncoding.EncodeToString(resp.GetFinalImageData()),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	case "FindP":
	case "FindS":
	case "FindAll":
	}
}
