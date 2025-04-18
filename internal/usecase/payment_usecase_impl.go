package usecase

import (
	"payment-options/internal/models"
	"payment-options/internal/repository"
	"sync"
)

type paymentUsecase struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(r repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{repo: r}
}

func (u *paymentUsecase) GetPaymentOptions1() (map[string]models.PaymentMethod, error) {
	result := make(map[string]models.PaymentMethod)
	result["ovo"] = u.repo.CallOVO()
	result["dana"] = u.repo.CallDANA()
	result["gopay"] = u.repo.CallGoPay()
	result["shopee"] = u.repo.CallShopee()
	result["oneklik"] = u.repo.CallOneKlik()
	result["bridd"] = u.repo.CallBRIDD()
	return result, nil
}

func (u *paymentUsecase) GetPaymentOptions() (map[string]models.PaymentMethod, error) {
	var wg sync.WaitGroup
	result := make(map[string]models.PaymentMethod)
	mu := sync.Mutex{}

	wg.Add(6)

	go func() {
		defer wg.Done()
		mu.Lock()
		result["ovo"] = u.repo.CallOVO()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["dana"] = u.repo.CallDANA()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["gopay"] = u.repo.CallGoPay()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["shopee"] = u.repo.CallShopee()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["oneklik"] = u.repo.CallOneKlik()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["bridd"] = u.repo.CallBRIDD()
		mu.Unlock()
	}()

	wg.Wait()
	return result, nil
}
