package repository

import "backend/internal/models"

type DatabaseRepo interface {
	AllLinks() ([]*models.Link, error)
}
