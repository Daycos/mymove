package internalapi

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	
	ppmop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/ppm"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/rateengine"
	"github.com/transcom/mymove/pkg/unit"
)

// ShowPPMEstimateHandler returns PPM SIT estimate for a weight, move date,
type ShowPPMEstimateHandler struct {
	handlers.HandlerContext
}

// Handle calculates a PPM reimbursement range.
func (h ShowPPMEstimateHandler) Handle(params ppmop.ShowPPMEstimateParams) middleware.Responder {
	logger := h.LoggerFromRequest(params.HTTPRequest)
	engine := rateengine.NewRateEngine(h.DB(), logger)

	distanceMilesFromOriginPickupZip, err := h.Planner().Zip5TransitDistance(params.OriginZip, params.DestinationZip)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	distanceMilesFromOriginDutyStationZip, err := h.Planner().Zip5TransitDistance(params.OriginDutyStationZip, params.DestinationZip)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	cost, err := engine.ComputeLowestCostPPMMove(
		unit.Pound(params.WeightEstimate),
		params.OriginZip,
		params.OriginDutyStationZip,
		params.DestinationZip,
		distanceMilesFromOriginPickupZip,
		distanceMilesFromOriginDutyStationZip,
		time.Time(params.OriginalMoveDate),
		0, // We don't want any SIT charges
	)
	if err != nil {
		return handlers.ResponseForError(logger, err)
	}

	min := cost.GCC.MultiplyFloat64(0.95)
	max := cost.GCC.MultiplyFloat64(1.05)

	ppmEstimate := internalmessages.PPMEstimateRange{
		RangeMin: swag.Int64(min.Int64()),
		RangeMax: swag.Int64(max.Int64()),
	}
	return ppmop.NewShowPPMEstimateOK().WithPayload(&ppmEstimate)
}
