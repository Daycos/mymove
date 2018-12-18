package publicapi

import (
	"log"
	"net/http"

	"github.com/go-openapi/loads"

	"github.com/transcom/mymove/pkg/gen/restapi"
	publicops "github.com/transcom/mymove/pkg/gen/restapi/apioperations"
	"github.com/transcom/mymove/pkg/handlers"
)

// NewPublicAPIHandler returns a handler for the public API
func NewPublicAPIHandler(context handlers.HandlerContext) http.Handler {

	// Wire up the handlers to the publicAPIMux
	apiSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	publicAPI := publicops.NewMymoveAPI(apiSpec)

	// Blackouts

	// Documents
	publicAPI.MoveDocsCreateGenericMoveDocumentHandler = CreateGenericMoveDocumentHandler{context}
	publicAPI.MoveDocsIndexMoveDocumentsHandler = IndexMoveDocumentsHandler{context}
	publicAPI.MoveDocsUpdateMoveDocumentHandler = UpdateMoveDocumentHandler{context}
	publicAPI.UploadsCreateUploadHandler = CreateUploadHandler{context}
	publicAPI.UploadsDeleteUploadHandler = DeleteUploadHandler{context}

	// Shipments
	publicAPI.ShipmentsIndexShipmentsHandler = IndexShipmentsHandler{context}
	publicAPI.ShipmentsGetShipmentHandler = GetShipmentHandler{context}
	publicAPI.ShipmentsPatchShipmentHandler = PatchShipmentHandler{context}
	publicAPI.ShipmentsAcceptShipmentHandler = AcceptShipmentHandler{context}
	publicAPI.ShipmentsRejectShipmentHandler = RejectShipmentHandler{context}
	publicAPI.ShipmentsTransportShipmentHandler = TransportShipmentHandler{context}
	publicAPI.ShipmentsDeliverShipmentHandler = DeliverShipmentHandler{context}

	publicAPI.ShipmentsCompletePmSurveyHandler = CompletePmSurveyHandler{context}
	publicAPI.ShipmentsCreateGovBillOfLadingHandler = CreateGovBillOfLadingHandler{context}

	// Accessorials
	publicAPI.AccessorialsGetShipmentLineItemsHandler = GetShipmentLineItemsHandler{context}
	publicAPI.AccessorialsUpdateShipmentLineItemHandler = UpdateShipmentLineItemHandler{context}
	publicAPI.AccessorialsCreateShipmentLineItemHandler = CreateShipmentLineItemHandler{context}
	publicAPI.AccessorialsDeleteShipmentLineItemHandler = DeleteShipmentLineItemHandler{context}
	publicAPI.AccessorialsApproveShipmentLineItemHandler = ApproveShipmentLineItemHandler{context}

	publicAPI.AccessorialsGetTariff400ngItemsHandler = GetTariff400ngItemsHandler{context}
	publicAPI.AccessorialsGetInvoiceHandler = GetInvoiceHandler{context}

	// Service Agents
	publicAPI.ServiceAgentsIndexServiceAgentsHandler = IndexServiceAgentsHandler{context}
	publicAPI.ServiceAgentsCreateServiceAgentHandler = CreateServiceAgentHandler{context}
	publicAPI.ServiceAgentsPatchServiceAgentHandler = PatchServiceAgentHandler{context}

	// TSPs
	publicAPI.TspsIndexTSPsHandler = TspsIndexTSPsHandler{context}
	publicAPI.TspsGetTspShipmentsHandler = TspsGetTspShipmentsHandler{context}

	return publicAPI.Serve(nil)
}
