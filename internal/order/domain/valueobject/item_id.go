package valueobject

import "errors"

type ItemID string

func NewItemID(val string) (ItemID, error) {
	if len(val) < 8 {
		return "", errors.New("item id too short")
	}
	return ItemID(val), nil
}
