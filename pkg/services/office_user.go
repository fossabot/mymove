package services

import (
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/gen/adminmessages"
	"github.com/transcom/mymove/pkg/models"
)

// OfficeUserFetcher is the exported interface for fetching a single office user
//go:generate mockery -name OfficeUserFetcher
type OfficeUserFetcher interface {
	FetchOfficeUser(filters []QueryFilter) (models.OfficeUser, error)
}

// OfficeUserCreator is the exported interface for creating an office user
//go:generate mockery -name OfficeUserCreator
type OfficeUserCreator interface {
	CreateOfficeUser(user *models.OfficeUser, transportationIDFilter []QueryFilter) (*models.OfficeUser, *validate.Errors, error)
}

// OfficeUserUpdater is the exported interface for creating an office user
//go:generate mockery -name OfficeUserUpdater
type OfficeUserUpdater interface {
	UpdateOfficeUser(id uuid.UUID, payload *adminmessages.OfficeUserUpdatePayload) (*models.OfficeUser, *validate.Errors, error)
}
