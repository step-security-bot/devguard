package asset

import (
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/l3montree-dev/devguard/internal/database/models"
)

type createRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`

	Importance            int  `json:"importance"`
	ReachableFromInternet bool `json:"reachableFromInternet"`

	ConfidentialityRequirement string `json:"confidentialityRequirement" validate:"required"`
	IntegrityRequirement       string `json:"integrityRequirement" validate:"required"`
	AvailabilityRequirement    string `json:"availabilityRequirement" validate:"required"`
}

func sanitizeRequirementLevel(level string) models.RequirementLevel {
	switch level {
	case "low", "medium", "high":
		return models.RequirementLevel(level)
	default:
		return "medium"
	}
}

func (a *createRequest) toModel(projectID uuid.UUID) models.Asset {
	return models.Asset{
		Name:        a.Name,
		Slug:        slug.Make(a.Name),
		ProjectID:   projectID,
		Description: a.Description,

		Importance:            a.Importance,
		ReachableFromInternet: a.ReachableFromInternet,

		ConfidentialityRequirement: sanitizeRequirementLevel(a.ConfidentialityRequirement),
		IntegrityRequirement:       sanitizeRequirementLevel(a.IntegrityRequirement),
		AvailabilityRequirement:    sanitizeRequirementLevel(a.AvailabilityRequirement),
	}
}

type patchRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`

	ReachableFromInternet *bool `json:"reachableFromInternet"`

	ConfidentialityRequirement *models.RequirementLevel `json:"confidentialityRequirement"`
	IntegrityRequirement       *models.RequirementLevel `json:"integrityRequirement"`
	AvailabilityRequirement    *models.RequirementLevel `json:"availabilityRequirement"`

	RepositoryID *string `json:"repositoryId"`

	IsPublic *bool `json:"isPublic"`
}

func (a *patchRequest) applyToModel(
	asset *models.Asset,
) bool {
	updated := false
	if a.Name != nil {
		updated = true
		asset.Name = *a.Name
		asset.Slug = slug.Make(*a.Name)
	}

	if a.Description != nil {
		updated = true
		asset.Description = *a.Description
	}

	if a.ReachableFromInternet != nil {
		updated = true
		asset.ReachableFromInternet = *a.ReachableFromInternet
	}

	if a.RepositoryID != nil {
		updated = true
		if *a.RepositoryID == "" {
			asset.RepositoryID = nil
		} else {
			asset.RepositoryID = a.RepositoryID
		}
	}

	if a.IsPublic != nil {
		updated = true
		asset.IsPublic = *a.IsPublic
	}

	return updated
}

type assetMetrics struct {
	// keyed by the scanType
	EnabledScanners []string `json:"enabledScanners"`
}
