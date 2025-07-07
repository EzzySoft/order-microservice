package valueobject

import "errors"

type DeliveryID string

func NewDeliveryID(val string) (DeliveryID, error) {
	if len(val) < 8 {
		return "", errors.New("delivery id too short")
	}
	return DeliveryID(val), nil
}
