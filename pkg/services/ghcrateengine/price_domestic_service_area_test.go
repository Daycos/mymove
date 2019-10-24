package ghcrateengine

import (
	"fmt"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
	"github.com/transcom/mymove/pkg/unit"
	"testing"
	"time"
)

const (
	dsaTestServiceArea = "005"
	dsaTestWeight      = unit.Pound(3500)
)

func (suite *GHCRateEngineServiceSuite) TestPriceDomesticServiceArea() {
	suite.setUpDomesticServiceAreaPricesData()
	pricer := NewDomesticServiceAreaPricer(suite.DB(), suite.logger, testdatagen.DefaultContractCode)

	services := map[string]string {
		"DODP": "Origin/Destination",
		"DFSIT": "SIT 1st Day",
		"DASIT": "SIT Add'l Days",
	}

	for serviceCode, serviceName := range services {
		suite.T().Run(fmt.Sprintf("success %s cost within peak period", serviceName), func(t *testing.T) {

			cost, err := pricer.PriceDomesticServiceArea (
				time.Date(testdatagen.TestYear, peakStart.month, peakStart.day, 0, 0, 0, 0, time.UTC),
				dsaTestWeight,
				dsaTestServiceArea,
				serviceCode)

			suite.NoError(err)
			suite.Equal(1234, cost.Int())
		})

		//suite.T().Run(fmt.Sprintf("success %s cost within non-peak period", serviceName), func(t *testing.T) {
		//
		//})
		//
		//suite.T().Run(fmt.Sprintf("%s cost weight below minimum", serviceName), func(t *testing.T) {
		//
		//})
		//
		//suite.T().Run(fmt.Sprintf("%s date outside of valid contract year", serviceName), func(t *testing.T) {
		//
		//})
	}

	//suite.T().Run("validation errors", func(t *testing.T) {
	//
	//}
}

func (suite *GHCRateEngineServiceSuite) setUpDomesticServiceAreaPricesData() {
	// create contractYear, domesticServiceArea, services data
	contractYear := testdatagen.MakeReContractYear(suite.DB(),
		testdatagen.Assertions{
			ReContractYear: models.ReContractYear{
				Escalation:           1.0197,
				EscalationCompounded: 1.04071,
			},
		})

	serviceArea := testdatagen.MakeReDomesticServiceArea(suite.DB(),
		testdatagen.Assertions{
			ReDomesticServiceArea: models.ReDomesticServiceArea{
				ServiceArea: dlhTestServiceArea,
			},
		})

	originDestinationService := testdatagen.MakeReService(suite.DB(),
		testdatagen.Assertions{
			ReService: models.ReService{
				Code: "DODP",
				Name: "Dom. O/D Price",
			},
		})

	sit1Service := testdatagen.MakeReService(suite.DB(),
		testdatagen.Assertions{
			ReService: models.ReService{
				Code: "DFSIT",
				Name: "Dom. O/D 1st Day SIT",
			},
		})

	addlSITService := testdatagen.MakeReService(suite.DB(),
		testdatagen.Assertions{
			ReService: models.ReService{
				Code: "DASIT",
				Name: "Dom. O/D Add'l SIT",
			},
		})

		baseDomesticServiceAreaPrice := models.ReDomesticServiceAreaPrice {
			ContractID: contractYear.Contract.ID,
			DomesticServiceAreaID: serviceArea.ID,
			IsPeakPeriod: true,
		}

		// Origin/Destination Price
		oDPrice := baseDomesticServiceAreaPrice
		oDPrice.ServiceID = originDestinationService.ID

		oDPeakPrice := oDPrice
		oDPeakPrice.PriceCents = 68
		suite.MustSave(&oDPeakPrice)

		oDNonpeakPrice := oDPrice
		oDNonpeakPrice.IsPeakPeriod = false
		oDNonpeakPrice.PriceCents = 76
		suite.MustSave(&oDNonpeakPrice)

		// SIT Day 1
		sit1Price := baseDomesticServiceAreaPrice
		sit1Price.ServiceID = sit1Service.ID

		sit1PeakPrice := sit1Price
		sit1PeakPrice.PriceCents = 3400
		suite.MustSave(&sit1PeakPrice)

		sit1NonpeakPrice := sit1Price
		sit1NonpeakPrice.IsPeakPeriod = false
		sit1NonpeakPrice.PriceCents = 112
		suite.MustSave(&sit1NonpeakPrice)

		// SIT Additional Days
		addlSITPrice := baseDomesticServiceAreaPrice
		addlSITPrice.ServiceID = addlSITService.ID

		addlSITPeakPrice := addlSITPrice
		addlSITPeakPrice.PriceCents = 3400
		suite.MustSave(&addlSITPeakPrice)

		addlSITNonpeakPrice := addlSITPrice
		addlSITNonpeakPrice.IsPeakPeriod = false
		addlSITNonpeakPrice.PriceCents = 112
		suite.MustSave(&addlSITNonpeakPrice)
}
