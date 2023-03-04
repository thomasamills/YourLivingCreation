package db

import "gorm.io/gorm"

func (h *HumanizeDbImpl) ListPresetStates(
	upperTx *gorm.DB,
) ([]string, error) {
	var arr []string
	listPresetStates := func(tx *gorm.DB) error {
		err := tx.Model(EmotionalState{}).Where(
			"is_preset = ?", true,
		).Select("entity_id").Find(&arr).Error
		if err != nil {
			return err
		}
		return nil
	}
	var err error
	if upperTx != nil {
		err = listPresetStates(upperTx)
	} else {
		err = h.mainDB.Transaction(func(tx *gorm.DB) error {
			err = listPresetStates(tx)
			if err != nil {
				tx.Rollback()
				return err
			}

			return nil
		})
	}
	if err != nil {
		return nil, err
	}
	return arr, nil
}
