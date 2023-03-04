package db

import (
	"fmt"
	"strings"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) CreateGenerationConfig(
	id string,
	config *humanize_protobuf.GenerationConfig,
	upperTx *gorm.DB,
) error {
	err := h.mainDB.Transaction(func(tx *gorm.DB) error {
		if upperTx != nil {
			tx = upperTx
		}
		genConfig := &GenerationConfig{
			GenerationConfigId:  id,
			Temperature:         config.Temperature,
			TopP:                config.TopP,
			TopK:                config.TopK,
			RepetitionPenalty:   config.RepetitionPenalty,
			Length:              config.Length,
			StopSequences:       strings.Join(config.StopSequences, ","),
			CharacterPrompt:     config.CharacterPrompt,
			NumberOfSynonyms:    config.WordsInEmotionalPrimer,
			MaxMemoryLogEntries: config.MaxMemoryLogEntries,
		}
		err := h.mainDB.Create(genConfig)
		if err != nil {
			if err.Error != nil {
				tx.Rollback()
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *HumanizeDbImpl) GetGenerationConfig(configId string,
) (*humanize_protobuf.GenerationConfig, error) {
	config := &GenerationConfig{}
	err := h.mainDB.
		Where("generation_config_id = ?", configId).
		First(&config).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println(err.Error())
		}
		return nil, err
	}
	protoGenConfig := &humanize_protobuf.GenerationConfig{
		GenerationConfigId:     configId,
		Temperature:            config.Temperature,
		TopP:                   config.TopP,
		TopK:                   config.TopK,
		RepetitionPenalty:      config.RepetitionPenalty,
		Length:                 config.Length,
		StopSequences:          strings.Split(config.StopSequences, ","),
		CharacterPrompt:        config.CharacterPrompt,
		WordsInEmotionalPrimer: config.NumberOfSynonyms,
		MaxMemoryLogEntries:    config.MaxMemoryLogEntries,
	}

	return protoGenConfig, nil
}

func (h *HumanizeDbImpl) ListConfigs() ([]string, error) {
	var arr []string
	err := h.mainDB.Model(GenerationConfig{}).Select("generation_config_id").Find(&arr).Error
	if err != nil {
		return nil, err
	}
	return arr, nil
}

func (h *HumanizeDbImpl) DeleteGenerationConfig(genConfigId string) error {
	err := h.mainDB.Where("generation_config_id = ?", genConfigId).Delete(&GenerationConfig{})
	if err != nil {
		if err.Error != nil {
			return err.Error
		}
	}
	return nil
}
