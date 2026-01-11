package usecase

import (
	"app-db-transactions/infrastructure/repository"
	"context"
)

type ComprarEstoqueUseCase struct {
	estoqueRepo *repository.EstoqueRepository
}

func NewComprarEstoqueUseCase(
	estoqueRepo *repository.EstoqueRepository,
) *ComprarEstoqueUseCase {
	return &ComprarEstoqueUseCase{
		estoqueRepo: estoqueRepo,
	}
}

func (uc *ComprarEstoqueUseCase) Execute(
	ctx context.Context,
	id int64,
) error {

	estoque, err := uc.estoqueRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if !estoque.PodeComprar() {
		return nil
	}

	return uc.estoqueRepo.AtualizarQuantidade(ctx, id)
}
