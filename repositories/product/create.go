package product

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"time"

	dto "github.com/srv-cashpay/product/dto"
	"github.com/srv-cashpay/product/entity"
	"golang.org/x/crypto/blake2b"
)

func (r *productRepository) Create(req dto.ProductRequest) (dto.ProductResponse, error) {
	// Insert or update the auto_increment value based on merchant_id
	var autoIncrement int
	err := r.DB.Raw(`
		INSERT INTO merchant_auto_increments (merchant_id, next_increment)
		VALUES (?, 1)
		ON CONFLICT (merchant_id) DO UPDATE
		SET next_increment = merchant_auto_increments.next_increment + 1
		RETURNING next_increment - 1;
	`, req.MerchantID).Scan(&autoIncrement).Error

	if err != nil {
		return dto.ProductResponse{}, err
	}

	// Generate Product ID with prefix and auto increment value
	prefix := "p="
	secureID, err := generateProductID(prefix, autoIncrement)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	// Create the new product entry
	create := entity.Product{
		ID:           secureID,
		ProductName:  req.ProductName,
		Stock:        req.Stock,
		MinimalStock: req.MinimalStock,
		Price:        req.Price,
		Status:       req.Status,
		UserID:       req.UserID,
		MerchantID:   req.MerchantID,
		CreatedBy:    req.CreatedBy,
		Description:  req.Description,
	}

	// Save the new product to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.ProductResponse{}, err
	}

	// Map the status from integer to string
	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	createdStatus, err := strconv.Atoi(fmt.Sprintf("%v", create.Status))
	if err != nil {
		return dto.ProductResponse{}, fmt.Errorf("invalid status value: %v", create.Status)
	}

	statusString, ok := statusMap[createdStatus]
	if !ok {
		return dto.ProductResponse{}, fmt.Errorf("invalid status value in database")
	}

	// Build the response for the created product
	response := dto.ProductResponse{
		ID:           create.ID,
		UserID:       create.UserID,
		MerchantID:   create.MerchantID,
		ProductName:  create.ProductName,
		Description:  create.Description,
		Stock:        create.Stock,
		MinimalStock: create.MinimalStock,
		Price:        create.Price,
		Status:       statusString,
		CreatedBy:    create.CreatedBy,
	}

	return response, nil
}

// Function to generate the product ID
func generateProductID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the product ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final product ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the product ID
func generateSecurePart() (string, error) {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

	// Generate a random string of length 12
	securePart := make([]byte, 12)
	_, err := rand.Read(securePart)
	if err != nil {
		return "", err
	}

	// Map each byte to a character from the chars string
	for i := range securePart {
		securePart[i] = chars[securePart[i]%byte(len(chars))]
	}

	return string(securePart), nil
}

// Function to generate a secure hash-based ID (if needed)
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
