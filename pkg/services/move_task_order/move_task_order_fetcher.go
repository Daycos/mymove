package movetaskorder

import (
	"database/sql"

	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
)

type moveTaskOrderFetcher struct {
	db *pop.Connection
}

func (f moveTaskOrderFetcher) ListMoveTaskOrders(moveOrderID uuid.UUID) ([]models.MoveTaskOrder, error) {
	var moveTaskOrders []models.MoveTaskOrder
	err := f.db.Where("move_order_id = $1", moveOrderID).Eager().All(&moveTaskOrders)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return []models.MoveTaskOrder{}, services.NotFoundError{}
		default:
			return []models.MoveTaskOrder{}, err
		}
	}
	return moveTaskOrders, nil
}

// NewMoveTaskOrderFetcher creates a new struct with the service dependencies
func NewMoveTaskOrderFetcher(db *pop.Connection) services.MoveTaskOrderFetcher {
	return &moveTaskOrderFetcher{db}
}

//FetchMoveTaskOrder retrieves a MoveTaskOrder for a given UUID
func (f moveTaskOrderFetcher) FetchMoveTaskOrder(moveTaskOrderID uuid.UUID) (*models.MoveTaskOrder, error) {
	mto := &models.MoveTaskOrder{}
	if err := f.db.Eager().Find(mto, moveTaskOrderID); err != nil {
		switch err {
		case sql.ErrNoRows:
			return &models.MoveTaskOrder{}, services.NewNotFoundError(moveTaskOrderID)
		default:
			return &models.MoveTaskOrder{}, err
		}
	}

	f.createDefaultServiceItems(mto)

	return mto, nil
}

func (f moveTaskOrderFetcher) createDefaultServiceItems(mto *models.MoveTaskOrder) error {
	var reServices []models.ReService
	err := f.db.Where("code in (?)", []string{"MS", "CS"}).All(&reServices)

	if err != nil {
		return err
	}

	defaultServiceItems := make(map[uuid.UUID]models.MTOServiceItem)
	for _, reService := range reServices {
		defaultServiceItems[reService.ID] = models.MTOServiceItem{
			ReServiceID:     reService.ID,
			MoveTaskOrderID: mto.ID,
		}
	}

	// Remove the ones that exist on the mto
	for _, item := range mto.MTOServiceItems {
		for _, reService := range reServices {
			if item.ReServiceID == reService.ID {
				delete(defaultServiceItems, reService.ID)
			}
		}
	}

	for _, serviceItem := range defaultServiceItems {
		_, err := f.db.ValidateAndCreate(&serviceItem)

		if err != nil {
			return err
		}
	}

	return nil
}
