package forms

import (
	"bytes"
	paperworkforms "github.com/transcom/mymove/pkg/paperwork"
)

// FormType is the type of form
type FormType int

// Form Types for CreateForm
const (
	GBL FormType = iota
)

// String returns the string value of the Form Type
func (ft FormType) String() string {
	return [...]string{"GBL", "SSW"}[ft]
}

// FormTemplate are the struct fields defined to call CreateForm service object
type FormTemplate struct {
	Buffer       *bytes.Reader
	FieldsLayout map[string]paperworkforms.FieldPos
	FormType
	FileName string
	Data     interface{}
}
