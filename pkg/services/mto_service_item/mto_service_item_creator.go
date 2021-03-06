package mtoserviceitem

import (
	"github.com/gobuffalo/validate"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
)

type createMTOServiceItemQueryBuilder interface {
	CreateOne(model interface{}) (*validate.Errors, error)
}

type mtoServiceItemCreator struct {
	builder createMTOServiceItemQueryBuilder
}

// CreateMTOServiceItem creates an MTO Service Item
func (o *mtoServiceItemCreator) CreateMTOServiceItem(serviceItem *models.MTOServiceItem) (*models.MTOServiceItem, *validate.Errors, error) {
	verrs, err := o.builder.CreateOne(serviceItem)
	if verrs != nil || err != nil {
		return nil, verrs, err
	}

	return serviceItem, nil, nil
}

// NewMTOServiceItemCreator returns a new MTO service item creator
func NewMTOServiceItemCreator(builder createMTOServiceItemQueryBuilder) services.MTOServiceItemCreator {
	return &mtoServiceItemCreator{builder}
}
