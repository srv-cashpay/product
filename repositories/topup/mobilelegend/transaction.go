package mobilelegend

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	dto "github.com/srv-cashpay/product/dto"
)

func (u *topupRepository) MobileLegend(req dto.MobileLegendRequest) (dto.MobileLegendResponse, error) {

	vendorURL := os.Getenv("TopUp")
	apiKey := os.Getenv("PKEY") // ini API KEY, bukan sign langsung

	if apiKey == "" {
		return dto.MobileLegendResponse{}, errors.New("PKEY is not set in env")
	}
	if vendorURL == "" {
		return dto.MobileLegendResponse{}, errors.New("TopUp vendor URL missing")
	}
	if req.UserID == "" || req.ZoneID == "" {
		return dto.MobileLegendResponse{}, errors.New("UserID & ZoneID wajib diisi")
	}

	// build customer_no
	customerNo := req.UserID + req.ZoneID

	if _, err := strconv.Atoi(customerNo); err != nil {
		return dto.MobileLegendResponse{}, errors.New("CustomerNo harus angka")
	}

	// generate ref_id jika kosong
	if req.RefID == "" {
		req.RefID = "TRX" + strconv.FormatInt(time.Now().UnixNano(), 10)
	}

	// SIGN = md5(username + apiKey + ref_id)
	sign := generateMD5(req.Username + apiKey + req.RefID)

	// Payload FINAL
	payload := map[string]string{
		"username":       req.Username,
		"buyer_sku_code": req.BuyerSkuCode,
		"customer_no":    customerNo,
		"ref_id":         req.RefID,
		"sign":           sign,
	}

	body, _ := json.Marshal(payload)

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

func generateMD5(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}
