package repository

// import (
// 	"context"
// 	models "github.com/cs161079/godbLib/Models"
// 	"gorm.io/gorm"
// )

// type UVersionRepository interface {
// 	Create(ctx context.Context, entity *models.UVersions) error
// 	Update(ctx context.Context, entity *models.UVersions) error
// 	//Transaction(ctx context.Context, fn func(repo UVersionRepository) error) error
// 	SelectAll(ctx context.Context) ([]models.UVersions, error)
// }

// // withTx creates a new repository instance with the given transaction
// func (r *Repository) withTx(tx *gorm.DB) UVersionRepository {
// 	return &Repository{
// 		Db: tx,
// 	}
// }

// // Update modifies an existing entity in the database
// func (r *Repository) Update(ctx context.Context, entity *models.UVersions) error {
// 	return r.Db.Table(SYNCVERSIONSTABLE).Save(entity).Error
// }

// func (r *Repository) SelectAll(ctx context.Context) ([]models.UVersions, error) {
// 	var res []models.UVersions

// 	if err := r.Db.Table(SYNCVERSIONSTABLE).Find(&res).Error; err != nil {
// 		return nil, err
// 	}
// 	return res, nil

// }

// // Transaction manages the transaction lifecycle
// func (r *Repository) Transaction(ctx context.Context, fn func(repo UVersionRepository) error) error {
// 	tx := r.Db.Begin()
// 	if tx.Error != nil {
// 		return tx.Error
// 	}
// 	repo := r.withTx(tx)
// 	err := fn(repo)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	return tx.Commit().Error
// }

// // Create adds a new entity to the database
// func (r *Repository) Create(ctx context.Context, entity *models.UVersions) error {
// 	return r.Db.Create(entity).Error
// }

// type UVersionsService struct {
// 	repo UVersionRepository
// }

// // PerformBusinessLogic performs multiple database operations within a transaction
// func (s *UVersionsService) LinePost(ctx context.Context, uversions *models.UVersions) error {
// 	return s.repo.Update(ctx, uversions)
// }
