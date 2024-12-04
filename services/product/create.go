package product

import (
	"crypto/rand"
	"fmt"
	"time"

	dto "github.com/srv-cashpay/product/dto"
	"golang.org/x/crypto/blake2b"
)

func (s *productService) Create(req dto.ProductRequest) (dto.ProductResponse, error) {
	if req.Status != 1 && req.Status != 2 {
		return dto.ProductResponse{}, fmt.Errorf("invalid status: must be 1 (active) or 2 (inactive)")
	}

	create := dto.ProductRequest{
		ProductName:  req.ProductName,
		Description:  req.Description,
		Stock:        req.Stock,
		MinimalStock: req.MinimalStock,
		Price:        req.Price,
		Status:       req.Status,
		UserID:       req.UserID,
		MerchantID:   req.MerchantID,
		CreatedBy:    req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	// Dapatkan string status berdasarkan nilai integer
	statusString, ok := statusMap[create.Status]
	if !ok {
		return dto.ProductResponse{}, fmt.Errorf("invalid status value in database")
	}

	response := dto.ProductResponse{
		ID:           created.ID,
		UserID:       created.UserID,
		ProductName:  created.ProductName,
		Description:  created.Description,
		Stock:        created.Stock,
		MinimalStock: created.MinimalStock,
		Price:        created.Price,
		Status:       statusString,
		MerchantID:   created.MerchantID,
		CreatedBy:    created.CreatedBy,
	}

	return response, nil
}

func generateSecureID() (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-="

	// Generate a salt
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Combine salt and current timestamp for uniqueness
	timestamp := time.Now().UnixNano()
	saltedID := fmt.Sprintf("%x%d", salt, timestamp)

	// Hash the combination using Blake2
	hash, err := blake2b.New512(nil)
	if err != nil {
		return "", err
	}
	hash.Write([]byte(saltedID))
	hashBytes := hash.Sum(nil)

	// Convert hash bytes into a valid string
	var secureID []byte
	for i := 0; i < 12; i++ {
		secureID = append(secureID, chars[hashBytes[i]%byte(len(chars))])
	}

	return string(secureID), nil
}
