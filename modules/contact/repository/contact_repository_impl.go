package repository

import (
	"context"
	"errors"
	"golang-contact-management-restful-api/modules/contact/domain"
	"golang-contact-management-restful-api/modules/contact/entities"
	"golang-contact-management-restful-api/modules/contact/models"
	"strings"

	"gorm.io/gorm"
)

type contactRepositoryImpl struct {
	DB *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepositoryImpl{DB: db}
}

func (repository *contactRepositoryImpl) Save(ctx context.Context, username string, contact entities.Contact) (entities.Contact, error) {
	contact.Username = username
	if err := repository.DB.WithContext(ctx).Create(&contact).Error; err != nil {
		return entities.Contact{}, err
	}

	return contact, nil
}

func (repository *contactRepositoryImpl) UpdateByID(ctx context.Context, username string, id int, contact entities.Contact) (entities.Contact, error) {
	updateMap := map[string]any{}

	if contact.Email != nil {
		updateMap["email"] = *contact.Email
	}

	if contact.Phone != nil {
		updateMap["phone"] = *contact.Phone
	}

	if contact.LastName != nil {
		updateMap["last_name"] = *contact.LastName
	}

	if contact.FirstName != "" {
		updateMap["first_name"] = contact.FirstName
	}

	if len(updateMap) == 0 {
		return entities.Contact{}, nil
	}

	result := repository.DB.WithContext(ctx).Model(&entities.Contact{}).
		Where("id = ?", id).Where("username = ?", username).
		Updates(updateMap)

	if result.Error != nil {
		return entities.Contact{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entities.Contact{}, domain.ErrContactNotFound
	}

	var updated entities.Contact
	if err := repository.DB.WithContext(ctx).Take(&updated, "id = ? AND username = ?", id, username).Error; err != nil {
		return entities.Contact{}, err
	}

	return updated, nil
}

func (repository *contactRepositoryImpl) FindByID(ctx context.Context, username string, id int) (entities.Contact, error) {
	var contact entities.Contact
	if err := repository.DB.WithContext(ctx).Where("id = ? AND username = ?", id, username).Take(&contact).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Contact{}, domain.ErrContactNotFound
		}
		return entities.Contact{}, err
	}

	return contact, nil
}

func (repository *contactRepositoryImpl) DeleteByID(ctx context.Context, username string, id int) error {
	result := repository.DB.WithContext(ctx).Where("id = ? AND username = ?", id, username).Delete(&entities.Contact{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrContactNotFound
	}
	return nil
}

func (repository *contactRepositoryImpl) Search(ctx context.Context, username string, query models.ContactSearchQuery) ([]entities.Contact, int, error) {
	page := query.Page
	if page < 1 {
		page = 1
	}

	size := query.Size
	if size <= 0 {
		size = 10
	}

	if size > 100 {
		size = 100
	}

	db := repository.DB.WithContext(ctx).Model(&entities.Contact{}).Where("username = ?", username)

	if space := strings.TrimSpace(query.Name); space != "" {
		like := "%" + space + "%"
		db = db.Where("first_name ILIKE ? OR last_name ILIKE ?", like, like)
	}

	if space := strings.TrimSpace(query.Email); space != "" {
		like := "%" + space + "%"
		db = db.Where("email ILIKE ?", like)
	}

	if space := strings.TrimSpace(query.Phone); space != "" {
		like := "%" + space + "%"
		db = db.Where("phone ILIKE ?", like)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var contacts []entities.Contact
	if err := db.Order("id DESC").Limit(size).Offset((page - 1) * size).
		Find(&contacts).Error; err != nil {
		return nil, 0, err
	}

	return contacts, int(total), nil

}
