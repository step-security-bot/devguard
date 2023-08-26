// Copyright (C) 2023 Tim Bastin, l3montree UG (haftungsbeschränkt)
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AppModel struct {
	ID        uuid.UUID    `gorm:"primarykey;type:uuid;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}

func (a AppModel) GetID() uuid.UUID {
	return a.ID
}

type User struct {
	AppModel
	Email string `json:"email" gorm:"type:varchar(255);unique"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
}

type Organization struct {
	AppModel
	Name                   string    `json:"name" gorm:"type:varchar(255)"`
	ContactPhoneNumber     *string   `json:"-" gorm:"type:varchar(255)"`
	NumberOfEmployees      *int      `json:"numberOfEmployees"`
	Country                *string   `json:"country" gorm:"type:varchar(255)"`
	Industry               *string   `json:"industry" gorm:"type:varchar(255)"`
	CriticalInfrastructure bool      `json:"criticalInfrastructure"`
	ISO27001               bool      `json:"iso27001"`
	NIST                   bool      `json:"nist"`
	Grundschutz            bool      `json:"grundschutz"`
	Projects               []Project `json:"projects"`
}

type Project struct {
	AppModel
	Name             string `json:"name" gorm:"type:varchar(255)"`
	Applications     []Application
	ServiceProviders []ServiceProvider
	OrganizationID   uuid.UUID `json:"organizationId"`
}

type Application struct {
	AppModel
	Name      string `json:"name" gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`

	Envs      []Env
	ProjectID uuid.UUID `json:"projectId"`
}

type Env struct {
	AppModel
	Name          string `json:"name" gorm:"type:varchar(255)"`
	Runs          []Report
	ApplicationID uuid.UUID `json:"applicationId"`
	IsDefault     bool      `json:"isDefault"`
}

type Report struct {
	AppModel
	ApplicationID        uuid.UUID `json:"applicationId"`
	DriverName           string    `json:"driverName"`
	DriverVersion        *string   `json:"driverVersion"`
	DriverInformationUri *string   `json:"driverInformationUri"`
	Results              []Result
	EnvID                uuid.UUID `json:"envId"`
	InitiatingUserID     string    `json:"initiatingUserId"`
}

type Result struct {
	AppModel
	ReportID  uuid.UUID      `json:"reportId"`
	RuleID    *string        `json:"ruleId"`
	Level     *string        `json:"level"`
	Message   *string        `json:"message"`
	Locations datatypes.JSON `gorm:"type:jsonb;default:'[]';not null"`
	Comments  []Comment
}

type Comment struct {
	AppModel
	ResultID uuid.UUID `json:"resultId"`
	UserID   uuid.UUID `json:"userId"`
	Comment  string    `json:"comment"`
}

type MitigationComment struct {
	AppModel
	MitigationID uuid.UUID `json:"mitigationId"`
	UserID       uuid.UUID `json:"userId"`
	Comment      string    `json:"comment"`
}

type (
	MitigationType string
)

const (
	MitigationTypeAvoid    MitigationType = "avoid"
	MitigationTypeAccept   MitigationType = "accept"
	MitigationTypeFix      MitigationType = "fix"
	MitigationTypeTransfer MitigationType = "transfer"
)

type Mitigation struct {
	AppModel
	MitigationType   MitigationType `json:"mitigationType"`
	InitiatingUserID string         `json:"initiatingUserId"`
	ResultID         uuid.UUID      `json:"resultId"`

	DueDate    *time.Time     `json:"dueDate"`
	Properties datatypes.JSON `gorm:"type:jsonb;default:'{}';not null"`

	MitigationPending bool `json:"mitigationPending" gorm:"default:false"` // will be true for fix and transfer types - we are waiting for another scan report which verifies, that the related result is fixed. Will be false for avoid and accept types

	propertiesMap any

	Comments []MitigationComment
}

type ServiceProvider struct {
	AppModel
	Name         string `json:"name" gorm:"unique;primarykey;type:varchar(255)"`
	ContactEmail string `json:"contact" gorm:"type:varchar(255)"`
	ProjectID    uuid.UUID
}

type MitigationTransferProperties struct {
	ServiceProviderID uuid.UUID `json:"serviceProviderId"`
}

type MitigationFixProperties = map[string]interface{}

type MitigationAcceptProperties struct {
	Justification string `json:"justification"`
}

type MitigationAvoidProperties struct {
	Justification string `json:"justification"`
}

// it is safe to typecast the return value to the correct type
func (m *Mitigation) GetProperties() any {
	if m.propertiesMap == nil {
		switch m.MitigationType {
		case MitigationTypeAvoid:
			m.propertiesMap = &MitigationAvoidProperties{}
		case MitigationTypeAccept:
			m.propertiesMap = &MitigationAcceptProperties{}
		case MitigationTypeFix:
			m.propertiesMap = &MitigationFixProperties{}
		case MitigationTypeTransfer:
			m.propertiesMap = &MitigationTransferProperties{}
		}

		json.Unmarshal(m.Properties, &m.propertiesMap)
	}
	return m.propertiesMap
}
