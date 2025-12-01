package mobilelegend

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	dto "github.com/srv-cashpay/product/dto"
)

func (u *topupRepository) MobileLegend(req dto.MobileLegendRequest) (dto.MobileLegendResponse, error) {
	var vendorURL = os.Getenv("TopUp")
	req.Sign = os.Getenv("PSign")

	if req.Sign == "" {
		return dto.MobileLegendResponse{}, errors.New("PSign is not set in environment variables")
	}

	// pastikan URL tidak kosong
	if vendorURL == "" {
		return dto.MobileLegendResponse{}, errors.New("VENDOR_URL is not set")
	}

	body, _ := json.Marshal(req)

	httpReq, _ := http.NewRequest("POST", vendorURL, bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return dto.MobileLegendResponse{}, err
	}

	defer httpResp.Body.Close()

	var result dto.MobileLegendResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return dto.MobileLegendResponse{}, err
	}

	return result, nil
}
