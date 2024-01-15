package handler

import (
	"clean-project-exp/internal/usecase"
	"encoding/json"
	"errors"
	"net/http"
)

type UserWebHandler struct {
	CreateUserUseCase *usecase.CreateUserUseCase
	ListUsersUseCase  *usecase.ListUsersUseCase
}

func NewUserWebHandler(createUserUseCase *usecase.CreateUserUseCase, listUsersUseCase *usecase.ListUsersUseCase) *UserWebHandler {
	return &UserWebHandler{CreateUserUseCase: createUserUseCase, ListUsersUseCase: listUsersUseCase}
}

func (h *UserWebHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users" {
		panic(errors.New("invalid URL for User Web Handler"))
	}

	switch r.Method {
	case "POST":
		h.create(w, r)
		break
	case "GET":
		h.list(w, r)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *UserWebHandler) create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInputDTO

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateUserUseCase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(output); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *UserWebHandler) list(w http.ResponseWriter, _ *http.Request) {
	output, err := h.ListUsersUseCase.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(output.Users) == 0 {
		output.Users = make([]usecase.UserForListingDTO, 0)
	}

	if err := json.NewEncoder(w).Encode(output); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
