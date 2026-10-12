package payment

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
)

type Service struct{
	mu       sync.RWMutex
	payments map[string]Payment
}

func NewService() *Service{
	return &Service{
		payments: make(map[string]Payment),
	}
}

func (s *Service) Create(req CreatePaymentRequest) (Payment, error){
	if err := req.Validate(); err != nil{
		return Payment{}, err
	}
	var idBytes [16]byte
	
	if _, err := rand.Read(idBytes[:]); err != nil {
		return Payment{}, err
	}

	p := Payment{
		ID:       "pay_" + hex.EncodeToString(idBytes[:]),
		Amount:   req.Amount,
		Currency: req.Currency,
		Status:   StatusCreated,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.payments[p.ID] = p
	return p, nil
}

func (s *Service) Get(id string) (Payment, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.payments[id]
	return p, ok
}


