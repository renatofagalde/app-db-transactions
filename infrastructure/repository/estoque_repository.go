package repository

import (
	"app-db-transactions/domain"
	"context"

	"gorm.io/gorm"
)

type EstoqueRepository struct {
	db *gorm.DB
}

func NewEstoqueRepository(db *gorm.DB) *EstoqueRepository {
	return &EstoqueRepository{
		db: db,
	}
}

func (r *EstoqueRepository) GetByID(
	ctx context.Context,
	id int64,
) (*domain.Estoque, error) {

	var estoque *domain.Estoque = &domain.Estoque{}

	if err := r.db.WithContext(ctx).
		Table("estoque").
		Where("id = ?", id).
		First(estoque).
		Error; err != nil {
		return nil, err
	}

	return estoque, nil
}

func (r *EstoqueRepository) AtualizarQuantidade(
	ctx context.Context,
	id int64,
) error {

	return r.db.WithContext(ctx).
		Table("estoque").
		Where("id = ?", id).
		Update("quantidade", gorm.Expr("quantidade - 1")).
		Error
}
