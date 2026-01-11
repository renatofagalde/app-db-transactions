package main

import (
	"app-db-transactions/application/usecase"
	"app-db-transactions/infrastructure/repository"
	"context"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const connStr = "postgres://user:pass@localhost:5510/db-transactions"

func main() {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	const n = 100
	var wg sync.WaitGroup
	wg.Add(n)

	ctx := context.Background()

	estoqueRepo := repository.NewEstoqueRepository(db)
	comprarUseCase := usecase.NewComprarEstoqueUseCase(estoqueRepo)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			_ = comprarUseCase.Execute(ctx, 1)
		}()
	}

	wg.Wait()
}
