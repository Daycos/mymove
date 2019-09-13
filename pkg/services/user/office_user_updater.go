package user

import (
	"github.com/gobuffalo/validate"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
	"github.com/transcom/mymove/pkg/services/query"
)

type officeUserUpdater struct {
	builder officeUserQueryBuilder
}

func (o *officeUserUpdater) UpdateOfficeUser(user *models.OfficeUser) (*models.OfficeUser, *validate.Errors, error) {
	var foundUser models.OfficeUser
	filters := []services.QueryFilter{query.NewQueryFilter("id", "=", user.ID.String())}
	err := o.builder.FetchOne(&foundUser, filters)

	if err != nil {
		return nil, nil, err
	}

	verrs, err := o.builder.UpdateOne(user)
	if verrs != nil || err != nil {
		return nil, verrs, err
	}

	return user, nil, nil
}

func NewOfficeUserUpdater(builder officeUserQueryBuilder) services.OfficeUserUpdater {
	return &officeUserUpdater{builder}
}
