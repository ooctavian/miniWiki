package repository

import (
	"context"

	"miniWiki/internal/domain/resource/model"
	"miniWiki/pkg/utils"

	"gorm.io/gorm"
)

var (
	rModel = model.Resource{}
)

type ResourceRepository struct {
	db *gorm.DB
}

func NewResourceRepository(db *gorm.DB) *ResourceRepository {
	return &ResourceRepository{
		db: db,
	}
}

func (r ResourceRepository) CountCategoryResources(ctx context.Context, id int) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(rModel).
		Where("category_id = ?", id).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r ResourceRepository) MakeResourcesPrivate(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).
		Model(rModel).
		Where("author_id = ?", id).
		Update("state", "PRIVATE").
		Error
	return err
}

func (r ResourceRepository) GetResourceById(ctx context.Context, id int) (*model.Resource, error) {
	var resource model.Resource
	err := r.db.WithContext(ctx).
		Take(&resource, id).
		Error
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func (r ResourceRepository) GetResources(ctx context.Context, accountId int, pagination utils.Pagination, filters model.GetResourcesFilters) (utils.Pagination, error) {
	var resources []model.Resource
	err := r.db.WithContext(ctx).
		Where(`(?::TEXT IS NULL OR title LIKE '%' || ? || '%') 
AND (?::TEXT IS NULL OR link LIKE '%'|| ? ||'%')
AND (?::TEXT IS NULL or category_id = ANY(?))
AND (state = 'PUBLIC' OR author_id = ?)`,
			filters.Title, filters.Title, filters.Link, filters.Link,
			filters.Categories, filters.Categories, accountId).
		Scopes(pagination.Paginate(resources, r.db)).
		Find(&resources).
		Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = resources
	return pagination, nil
}

func (r ResourceRepository) DeleteResourceById(ctx context.Context, resourceId uint, accountId uint) error {
	err := r.db.WithContext(ctx).
		Where("resource_id = ? AND author_id = ?", resourceId, accountId).
		Delete(&rModel).
		Error
	return err
}

func (r ResourceRepository) InsertResource(ctx context.Context, resource model.CreateResource) (uint, error) {
	err := r.db.WithContext(ctx).
		Create(&resource).
		Error
	if err != nil {
		return 0, err
	}
	return resource.ID, nil
}

func (r ResourceRepository) UpdateResource(ctx context.Context, request model.UpdateResourceRequest) error {
	err := r.db.WithContext(ctx).
		Where("resource_id = ? AND author_id = ?",
			request.ResourceId, request.AccountId).
		Updates(&request.Resource).
		Error
	return err
}

func (r ResourceRepository) UpdateResourcePicture(ctx context.Context, resourceId int, accountId int, path string) error {
	err := r.db.WithContext(ctx).
		Model(&rModel).
		Where("resource_id = ? AND author_id = ?",
			resourceId, accountId).
		Update("picture_url", path).
		Error
	return err
}
