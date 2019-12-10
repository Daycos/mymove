package testdatagen

import (
	"github.com/go-openapi/swag"
	"github.com/gobuffalo/pop"

	"github.com/transcom/mymove/pkg/models"
)

// MakePaymentRequest creates a single PaymentRequest and associated set relationships
func MakePaymentRequest(db *pop.Connection, assertions Assertions) models.PaymentRequest {

	// Create new PaymentRequest if not provided
	paymentRequest := models.PaymentRequest{
		IsFinal:         false,
		RejectionReason: swag.String("Not good enough"),
		Status:          models.PaymentRequestStatusPending,
	}

	// Overwrite values with those from assertions
	mergeModels(&paymentRequest, assertions.PaymentRequest)

	mustCreate(db, &paymentRequest)

	return paymentRequest
}

// MakeDefaultPaymentRequest makes an PaymentRequest with default values
func MakeDefaultPaymentRequest(db *pop.Connection) models.PaymentRequest {
	return MakePaymentRequest(db, Assertions{})
}
