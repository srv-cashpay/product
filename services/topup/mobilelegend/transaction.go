package mobilelegend

import (
	dto "github.com/srv-cashpay/product/dto"
)

func (s *topupService) MobileLegend(req dto.MobileLegendRequest) (dto.MobileLegendResponse, error) {
	return s.Repo.MobileLegend(req)
}
