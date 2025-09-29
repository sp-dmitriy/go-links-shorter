package link

import "go-links-shorter/pkg/db"

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(datadase *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: datadase,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
