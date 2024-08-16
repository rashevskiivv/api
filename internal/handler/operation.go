package handler

import "tax-api/internal/repository"

type OperationHandler struct {
	repo repository.OperationRepo
}

func NewOperationHandler(repo repository.OperationRepo) OperationHandler {
	return OperationHandler{
		repo: repo,
	}
}
