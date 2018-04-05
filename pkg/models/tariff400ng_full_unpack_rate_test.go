package models_test

import (
	"time"

	. "github.com/transcom/mymove/pkg/models"
)

func (suite *ModelSuite) Test_UnpackEffectiveDateValidation() {
	now := time.Now()

	validUnpackRate := Tariff400ngFullUnpackRate{
		EffectiveDateLower: now,
		EffectiveDateUpper: now.AddDate(1, 0, 0),
	}

	expErrors := map[string][]string{}
	suite.verifyValidationErrors(&validUnpackRate, expErrors)

	invalidUnpackRate := Tariff400ngFullUnpackRate{
		EffectiveDateLower: now,
		EffectiveDateUpper: now.AddDate(-1, 0, 0),
	}

	expErrors = map[string][]string{
		"effective_date_upper": []string{"EffectiveDateUpper must be after EffectiveDateLower."},
	}
	suite.verifyValidationErrors(&invalidUnpackRate, expErrors)
}

func (suite *ModelSuite) Test_UnpackRateValidation() {
	validUnpackRate := Tariff400ngFullUnpackRate{
		RateMillicents: 100,
	}

	expErrors := map[string][]string{}
	suite.verifyValidationErrors(&validUnpackRate, expErrors)

	invalidUnpackRate := Tariff400ngFullUnpackRate{
		RateMillicents: -1,
	}

	expErrors = map[string][]string{
		"rate_millicents": []string{"-1 is not greater than -1."},
	}
	suite.verifyValidationErrors(&invalidUnpackRate, expErrors)
}

func (suite *ModelSuite) Test_UnpackCreateAndSave() {
	now := time.Now()

	validUnpackRate := Tariff400ngFullUnpackRate{
		RateMillicents:     100,
		Schedule:           1,
		EffectiveDateLower: now,
		EffectiveDateUpper: now.AddDate(0, 1, 0),
	}

	suite.mustSave(&validUnpackRate)
}
