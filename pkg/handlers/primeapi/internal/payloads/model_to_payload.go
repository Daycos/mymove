package payloads

import (
	"github.com/transcom/mymove/pkg/etag"

	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/primemessages"
	"github.com/transcom/mymove/pkg/models"
)

// MoveTaskOrder payload
func MoveTaskOrder(moveTaskOrder *models.MoveTaskOrder) *primemessages.MoveTaskOrder {
	if moveTaskOrder == nil {
		return nil
	}
	paymentRequests := PaymentRequests(&moveTaskOrder.PaymentRequests)
	mtoServiceItems := MTOServiceItems(&moveTaskOrder.MTOServiceItems)
	mtoShipments := MTOShipments(&moveTaskOrder.MTOShipments)
	payload := &primemessages.MoveTaskOrder{
		ID:                 strfmt.UUID(moveTaskOrder.ID.String()),
		CreatedAt:          strfmt.Date(moveTaskOrder.CreatedAt),
		IsAvailableToPrime: &moveTaskOrder.IsAvailableToPrime,
		IsCanceled:         &moveTaskOrder.IsCanceled,
		MoveOrderID:        strfmt.UUID(moveTaskOrder.MoveOrderID.String()),
		MoveOrder:          MoveOrder(&moveTaskOrder.MoveOrder),
		ReferenceID:        moveTaskOrder.ReferenceID,
		PaymentRequests:    *paymentRequests,
		MtoServiceItems:    *mtoServiceItems,
		PpmEstimatedWeight: int64(moveTaskOrder.PPMEstimatedWeight),
		PpmType:            moveTaskOrder.PPMType,
		MtoShipments:       *mtoShipments,
		UpdatedAt:          strfmt.Date(moveTaskOrder.UpdatedAt),
	}
	return payload
}

// MoveTaskOrderWithEtag payload
func MoveTaskOrderWithEtag(moveTaskOrder *models.MoveTaskOrder) *primemessages.MoveTaskOrderWithEtag {
	if moveTaskOrder == nil {
		return nil
	}
	paymentRequests := PaymentRequests(&moveTaskOrder.PaymentRequests)
	mtoServiceItems := MTOServiceItems(&moveTaskOrder.MTOServiceItems)
	mtoShipments := MTOShipments(&moveTaskOrder.MTOShipments)
	payload := &primemessages.MoveTaskOrderWithEtag{
		MoveTaskOrder: primemessages.MoveTaskOrder{
			ID:                 strfmt.UUID(moveTaskOrder.ID.String()),
			CreatedAt:          strfmt.Date(moveTaskOrder.CreatedAt),
			IsAvailableToPrime: &moveTaskOrder.IsAvailableToPrime,
			IsCanceled:         &moveTaskOrder.IsCanceled,
			MoveOrderID:        strfmt.UUID(moveTaskOrder.MoveOrderID.String()),
			ReferenceID:        moveTaskOrder.ReferenceID,
			PaymentRequests:    *paymentRequests,
			PpmEstimatedWeight: moveTaskOrder.PPMEstimatedWeight.Int64(),
			PpmType:            moveTaskOrder.PPMType,
			MtoServiceItems:    *mtoServiceItems,
			MtoShipments:       *mtoShipments,
			UpdatedAt:          strfmt.Date(moveTaskOrder.UpdatedAt),
		},
		ETag: etag.GenerateEtag(moveTaskOrder.UpdatedAt),
	}

	return payload
}

// MoveTaskOrders payload
func MoveTaskOrders(moveTaskOrders *models.MoveTaskOrders) []*primemessages.MoveTaskOrder {
	payload := make(primemessages.MoveTaskOrders, len(*moveTaskOrders))

	for i, m := range *moveTaskOrders {
		payload[i] = MoveTaskOrder(&m)
	}
	return payload
}

// Customer payload
func Customer(customer *models.Customer) *primemessages.Customer {
	if customer == nil {
		return nil
	}
	payload := primemessages.Customer{
		FirstName:          customer.FirstName,
		LastName:           customer.LastName,
		DodID:              customer.DODID,
		ID:                 strfmt.UUID(customer.ID.String()),
		UserID:             strfmt.UUID(customer.UserID.String()),
		CurrentAddress:     Address(&customer.CurrentAddress),
		DestinationAddress: Address(&customer.DestinationAddress),
		Branch:             customer.Agency,
	}

	if customer.PhoneNumber != nil {
		payload.Phone = *customer.PhoneNumber
	}

	if customer.Email != nil {
		payload.Email = *customer.Email
	}
	return &payload
}

// MoveOrder payload
func MoveOrder(moveOrder *models.MoveOrder) *primemessages.MoveOrder {
	if moveOrder == nil {
		return nil
	}
	destinationDutyStation := DutyStation(moveOrder.DestinationDutyStation)
	originDutyStation := DutyStation(moveOrder.OriginDutyStation)
	if moveOrder.Grade != nil {
		moveOrder.Entitlement.SetWeightAllotment(*moveOrder.Grade)
	}
	reportByDate := strfmt.Date(*moveOrder.ReportByDate)
	entitlements := Entitlement(moveOrder.Entitlement)
	payload := primemessages.MoveOrder{
		CustomerID:             strfmt.UUID(moveOrder.CustomerID.String()),
		Customer:               Customer(moveOrder.Customer),
		DestinationDutyStation: destinationDutyStation,
		Entitlement:            entitlements,
		ID:                     strfmt.UUID(moveOrder.ID.String()),
		OriginDutyStation:      originDutyStation,
		OrderNumber:            moveOrder.OrderNumber,
		LinesOfAccounting:      moveOrder.LinesOfAccounting,
		Rank:                   moveOrder.Grade,
		ConfirmationNumber:     *moveOrder.ConfirmationNumber,
		ReportByDate:           reportByDate,
	}
	return &payload
}

// Entitlement payload
func Entitlement(entitlement *models.Entitlement) *primemessages.Entitlements {
	if entitlement == nil {
		return nil
	}
	var proGearWeight, proGearWeightSpouse, totalWeight int64
	if entitlement.WeightAllotment() != nil {
		proGearWeight = int64(entitlement.WeightAllotment().ProGearWeight)
		proGearWeightSpouse = int64(entitlement.WeightAllotment().ProGearWeightSpouse)
		totalWeight = int64(entitlement.WeightAllotment().TotalWeightSelf)
	}
	var authorizedWeight *int64
	if entitlement.AuthorizedWeight() != nil {
		aw := int64(*entitlement.AuthorizedWeight())
		authorizedWeight = &aw
	}
	var sit int64
	if entitlement.StorageInTransit != nil {
		sit = int64(*entitlement.StorageInTransit)
	}
	var totalDependents int64
	if entitlement.TotalDependents != nil {
		totalDependents = int64(*entitlement.TotalDependents)
	}
	return &primemessages.Entitlements{
		ID:                    strfmt.UUID(entitlement.ID.String()),
		AuthorizedWeight:      authorizedWeight,
		DependentsAuthorized:  entitlement.DependentsAuthorized,
		NonTemporaryStorage:   entitlement.NonTemporaryStorage,
		PrivatelyOwnedVehicle: entitlement.PrivatelyOwnedVehicle,
		ProGearWeight:         proGearWeight,
		ProGearWeightSpouse:   proGearWeightSpouse,
		StorageInTransit:      sit,
		TotalDependents:       totalDependents,
		TotalWeight:           totalWeight,
	}
}

// DutyStation payload
func DutyStation(dutyStation *models.DutyStation) *primemessages.DutyStation {
	if dutyStation == nil {
		return nil
	}
	address := Address(&dutyStation.Address)
	payload := primemessages.DutyStation{
		Address:   address,
		AddressID: address.ID,
		ID:        strfmt.UUID(dutyStation.ID.String()),
		Name:      dutyStation.Name,
	}
	return &payload
}

// Address payload
func Address(address *models.Address) *primemessages.Address {
	if address == nil {
		return nil
	}
	return &primemessages.Address{
		ID:             strfmt.UUID(address.ID.String()),
		StreetAddress1: &address.StreetAddress1,
		StreetAddress2: address.StreetAddress2,
		StreetAddress3: address.StreetAddress3,
		City:           &address.City,
		State:          &address.State,
		PostalCode:     &address.PostalCode,
		Country:        address.Country,
	}
}

// PaymentRequest payload
func PaymentRequest(paymentRequest *models.PaymentRequest) *primemessages.PaymentRequest {
	return &primemessages.PaymentRequest{
		ID:              strfmt.UUID(paymentRequest.ID.String()),
		Status:          primemessages.PaymentRequestStatus(paymentRequest.Status),
		IsFinal:         &paymentRequest.IsFinal,
		MoveTaskOrderID: strfmt.UUID(paymentRequest.MoveTaskOrderID.String()),
		RejectionReason: paymentRequest.RejectionReason,
	}
}

// PaymentRequests payload
func PaymentRequests(paymentRequests *models.PaymentRequests) *primemessages.PaymentRequests {
	payload := make(primemessages.PaymentRequests, len(*paymentRequests))

	for i, p := range *paymentRequests {
		payload[i] = PaymentRequest(&p)
	}
	return &payload
}

// MTOShipment payload
func MTOShipment(mtoShipment *models.MTOShipment) *primemessages.MTOShipment {
	payload := &primemessages.MTOShipment{
		ID:                       strfmt.UUID(mtoShipment.ID.String()),
		MoveTaskOrderID:          strfmt.UUID(mtoShipment.MoveTaskOrderID.String()),
		ShipmentType:             primemessages.MTOShipmentType(mtoShipment.ShipmentType),
		CustomerRemarks:          *mtoShipment.CustomerRemarks,
		RequestedPickupDate:      strfmt.Date(*mtoShipment.RequestedPickupDate),
		ScheduledPickupDate:      strfmt.Date(*mtoShipment.ScheduledPickupDate),
		PickupAddress:            Address(&mtoShipment.PickupAddress),
		Status:                   string(mtoShipment.Status),
		DestinationAddress:       Address(&mtoShipment.DestinationAddress),
		SecondaryPickupAddress:   Address(mtoShipment.SecondaryPickupAddress),
		SecondaryDeliveryAddress: Address(mtoShipment.SecondaryDeliveryAddress),
		CreatedAt:                strfmt.DateTime(mtoShipment.CreatedAt),
		UpdatedAt:                strfmt.DateTime(mtoShipment.UpdatedAt),
		ETag:                     etag.GenerateEtag(mtoShipment.UpdatedAt),
	}

	if mtoShipment.ApprovedDate != nil && !mtoShipment.ApprovedDate.IsZero() {
		approvedDate := strfmt.Date(*mtoShipment.ApprovedDate)
		payload.ApprovedDate = &approvedDate
	}

	if mtoShipment.ActualPickupDate != nil && !mtoShipment.ActualPickupDate.IsZero() {
		payload.ActualPickupDate = strfmt.Date(*mtoShipment.ActualPickupDate)
	}

	if mtoShipment.FirstAvailableDeliveryDate != nil && !mtoShipment.FirstAvailableDeliveryDate.IsZero() {
		payload.FirstAvailableDeliveryDate = strfmt.Date(*mtoShipment.FirstAvailableDeliveryDate)
	}

	if mtoShipment.PrimeEstimatedWeight != nil {
		payload.PrimeEstimatedWeight = int64(*mtoShipment.PrimeEstimatedWeight)
		payload.PrimeEstimatedWeightRecordedDate = strfmt.Date(*mtoShipment.PrimeEstimatedWeightRecordedDate)
	}

	return payload
}

// MTOShipments payload
func MTOShipments(mtoShipments *models.MTOShipments) *primemessages.MTOShipments {
	payload := make(primemessages.MTOShipments, len(*mtoShipments))

	for i, m := range *mtoShipments {
		payload[i] = MTOShipment(&m)
	}
	return &payload
}

// MTOServiceItem payload
func MTOServiceItem(mtoServiceItem *models.MTOServiceItem) *primemessages.MTOServiceItem {
	return &primemessages.MTOServiceItem{
		ID:              strfmt.UUID(mtoServiceItem.ID.String()),
		MoveTaskOrderID: strfmt.UUID(mtoServiceItem.MoveTaskOrderID.String()),
		ReServiceID:     strfmt.UUID(mtoServiceItem.ReServiceID.String()),
		ReServiceCode:   mtoServiceItem.ReService.Code,
		ReServiceName:   mtoServiceItem.ReService.Name,
	}
}

// MTOServiceItems payload
func MTOServiceItems(mtoServiceItems *models.MTOServiceItems) *primemessages.MTOServiceItems {
	payload := make(primemessages.MTOServiceItems, len(*mtoServiceItems))

	for i, p := range *mtoServiceItems {
		payload[i] = MTOServiceItem(&p)
	}
	return &payload
}
