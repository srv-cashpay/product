package mobilelegend

import (
	m "github.com/srv-cashpay/middlewares/middlewares"
	dto "github.com/srv-cashpay/product/dto"

	r "github.com/srv-cashpay/product/repositories/topup/mobilelegend"
)

type MobileLegendService interface {
	MobileLegend(req dto.MobileLegendRequest) (dto.MobileLegendResponse, error)
}

type topupService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewTopUpService(Repo r.DomainRepository, jwtS m.JWTService) MobileLegendService {
	return &topupService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
