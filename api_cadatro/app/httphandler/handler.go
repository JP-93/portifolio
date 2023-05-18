package httphandler

import (
	"context"
	"encoding/json"
	"syscall/js"

	"net/http"

	"github.com/portifolio/api_cadatro/app/service"
)

type creatorService interface {
	Create(ctx context.Context, s service.Request) error
}

type ResgisterData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type handlerService struct {
	svc creatorService
}

func NewHandler(svc creatorService) *handlerService {
	return &handlerService{svc: svc}
}



func (s *handlerService) handler(w http.ResponseWriter, r *http.Request) {
	var re ResgisterData
	ctx := context.Background()
	json.NewDecoder(r.Body).Decode(re)
	defer r.Body.Close()

	value := service.Request{
		Email: re.Email,
		Name: re.Name,
		Phone: re.Phone,
	}
	err := s.svc.Create(ctx, value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
