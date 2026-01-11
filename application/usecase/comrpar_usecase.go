package usecase

import (
	"app-db-transactions/infrastructure/repository"
	"context"
)

type ComprarEstoqueUseCase struct {
	estoqueRepository *repository.EstoqueRepository
}

func NewComprarEstoqueUseCase(
	estoqueRepository *repository.EstoqueRepository,
) *ComprarEstoqueUseCase {
	return &ComprarEstoqueUseCase{
		estoqueRepository: estoqueRepository,
	}
}

func (uc *ComprarEstoqueUseCase) Vender(
	ctx context.Context,
	id int64,
) error {

	const maxTentativas = 50

	for i := 0; i < maxTentativas; i++ {
		estoque, err := uc.estoqueRepository.GetByID(ctx, id)
		if err != nil {
			return err
		}

		if !estoque.PodeComprar() {
			return nil
		}

		ok, err := uc.estoqueRepository.AtualizarQuantidade(ctx, id, estoque.Version)
		if err != nil {
			return err
		}

		if ok {
			return nil
		}

	}

	return nil
}
