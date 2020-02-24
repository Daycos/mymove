package models

import (
	"time"

	"github.com/transcom/mymove/pkg/auth"

	"github.com/transcom/mymove/pkg/unit"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// WeightTicketSetType represents types of weight ticket sets
type WeightTicketSetType string

const (
	// WeightTicketSetTypeCAR captures enum value "CAR"
	WeightTicketSetTypeCAR WeightTicketSetType = "CAR"

	// WeightTicketSetTypeCARTRAILER captures enum value "CAR_TRAILER"
	WeightTicketSetTypeCARTRAILER WeightTicketSetType = "CAR_TRAILER"

	// WeightTicketSetTypeBOXTRUCK captures enum value "BOXTRUCK"
	WeightTicketSetTypeBOXTRUCK WeightTicketSetType = "BOXTRUCK"

	// WeightTicketSetTypePROGEAR captures enum value "PROGEAR"
	WeightTicketSetTypePROGEAR WeightTicketSetType = "PROGEAR"
)

// WeightTicketSetDocument weight ticket documents payload
type WeightTicketSetDocument struct {
	ID                       uuid.UUID           `json:"id" db:"id"`
	MoveDocumentID           uuid.UUID           `json:"move_document_id" db:"move_document_id"`
	MoveDocument             MoveDocument        `belongs_to:"move_documents"`
	EmptyWeight              *unit.Pound         `json:"empty_weight,omitempty" db:"empty_weight"`
	EmptyWeightTicketMissing bool                `json:"empty_weight_ticket_missing,omitempty" db:"empty_weight_ticket_missing"`
	FullWeight               *unit.Pound         `json:"full_weight,omitempty" db:"full_weight"`
	FullWeightTicketMissing  bool                `json:"full_weight_ticket_missing,omitempty" db:"full_weight_ticket_missing"`
	VehicleNickname          *string             `json:"vehicle_nickname,omitempty" db:"vehicle_nickname"`
	VehicleMake              *string             `json:"vehicle_make,omitempty" db:"vehicle_make"`
	VehicleModel             *string             `json:"vehicle_model,omitempty" db:"vehicle_model"`
	WeightTicketSetType      WeightTicketSetType `json:"weight_ticket_set_type,omitempty" db:"weight_ticket_set_type"`
	WeightTicketDate         *time.Time          `json:"weight_ticket_date,omitempty" db:"weight_ticket_date"`
	TrailerOwnershipMissing  bool                `json:"trailer_ownership_missing,omitempty" db:"trailer_ownership_missing"`
	CreatedAt                time.Time           `json:"created_at" db:"created_at"`
	UpdatedAt                time.Time           `json:"updated_at" db:"updated_at"`
	DeletedAt                *time.Time          `db:"deleted_at"`
}

// WeightTicketSetDocuments slice of WeightTicketSetDocuments
type WeightTicketSetDocuments []WeightTicketSetDocuments

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (m *WeightTicketSetDocument) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: m.MoveDocumentID, Name: "MoveDocumentID"},
		&validators.StringIsPresent{Field: string(m.WeightTicketSetType), Name: "WeightTicketSetType"},
		&MustBeBothNilOrBothNotNil{
			FieldName1:  "VehicleMake",
			FieldValue1: m.VehicleMake,
			FieldName2:  "VehicleModel",
			FieldValue2: m.VehicleModel,
		},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (m *WeightTicketSetDocument) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (m *WeightTicketSetDocument) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// SumWeightTicketSetsForPPM iterates through move documents that are weight ticket sets and accumulates
// the net weight if the weight ticket set has an OK status
func SumWeightTicketSetsForPPM(db *pop.Connection, session *auth.Session, ppmID uuid.UUID) (*unit.Pound, error) {
	status := MoveDocumentStatusOK
	var totalWeight unit.Pound
	weightTicketSets, err := FetchMoveDocuments(db, session, ppmID, &status, MoveDocumentTypeWEIGHTTICKETSET, false)

	if err != nil {
		return &totalWeight, err
	}

	for _, weightTicketSet := range weightTicketSets {
		wt := weightTicketSet.WeightTicketSetDocument
		if wt != nil && wt.FullWeight != nil && wt.EmptyWeight != nil {
			totalWeight += *wt.FullWeight - *wt.EmptyWeight
		}
	}
	return &totalWeight, nil
}
