package valueobject

import "errors"

type ItemRID string

func NewItemRID(val string) (ItemRID, error) {
	if len(val) < 8 {
		return "", errors.New("item rid too short")
	}
	return ItemRID(val), nil
}
