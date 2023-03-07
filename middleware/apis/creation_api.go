package apis

import (
	"context"
	"errors"
	"testserver/db"
	"testserver/db/id_gen"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

type CreationApiImpl struct {
	db    db.HumanizeDB
	idGen id_gen.ULIDGenerator
	humanize_protobuf.UnimplementedRuleApiServer
}

func NewCreationApi(db db.HumanizeDB, generator id_gen.ULIDGenerator) *CreationApiImpl {
	return &CreationApiImpl{
		db:    db,
		idGen: generator,
	}
}

// List -----------------------------------------------------------------------------------
func (c CreationApiImpl) ListGenerationConfigs(
	ctx context.Context, request *humanize_protobuf.ListGenerationConfigRequest,
) (*humanize_protobuf.ListGenerationConfigResponse, error) {
	genConfigIds, err := c.db.ListConfigs()
	if err != nil {
		errMessage := "could not retrieve list of generation configs"
		return &humanize_protobuf.ListGenerationConfigResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.ListGenerationConfigResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.ListGenerationConfigResponse{
		ConfigIds: genConfigIds,
		Status:    humanize_protobuf.ListGenerationConfigResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) ListPersonalityIds(
	ctx context.Context, request *humanize_protobuf.ListPersonalityIdsRequest,
) (*humanize_protobuf.ListPersonalityIdsResponse, error) {
	personalityIds, err := c.db.ListPersonalities()
	if err != nil {
		errMessage := "could not retrieve list of personality ids"
		return &humanize_protobuf.ListPersonalityIdsResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.ListPersonalityIdsResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.ListPersonalityIdsResponse{
		PersonalityIds: personalityIds,
		Status:         humanize_protobuf.ListPersonalityIdsResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) ListPresetStateIds(
	ctx context.Context, request *humanize_protobuf.ListPresetStateIdsRequest,
) (*humanize_protobuf.ListPresetStateIdsResponse, error) {
	stateIds, err := c.db.ListPresetStates(nil)
	if err != nil {
		errMessage := "could not retrieve list of preset state ids"
		return &humanize_protobuf.ListPresetStateIdsResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.ListPresetStateIdsResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.ListPresetStateIdsResponse{
		PresetStateIds: stateIds,
		Status:         humanize_protobuf.ListPresetStateIdsResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) ListPrompts(
	ctx context.Context, request *humanize_protobuf.ListPromptsRequest,
) (*humanize_protobuf.ListPromptsResponse, error) {
	promptIds, err := c.db.ListPrompts(request.PromptSetId, nil)
	if err != nil {
		errMessage := "could not retrieve list of preset state ids"
		return &humanize_protobuf.ListPromptsResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.ListPromptsResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.ListPromptsResponse{
		PromptIds: promptIds,
		Status:    humanize_protobuf.ListPromptsResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) ListPromptSetIds(
	ctx context.Context, request *humanize_protobuf.ListPromptSetIdsRequest,
) (*humanize_protobuf.ListPromptSetIdsResponse, error) {
	promptSetIds, err := c.db.ListPromptSetIds(nil)
	if err != nil {
		errMessage := "could not retrieve list of preset state ids"
		return &humanize_protobuf.ListPromptSetIdsResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.ListPromptSetIdsResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.ListPromptSetIdsResponse{
		PromptSetIds: promptSetIds,
		Status:       humanize_protobuf.ListPromptSetIdsResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) ListPromptSegmentSetIds(
	ctx context.Context, request *humanize_protobuf.ListPromptSegmentSetIdsRequest,
) (*humanize_protobuf.ListPromptSegmentSetIdsResponse, error) {
	promptSegmentSetIds, err := c.db.ListPromptSegmentSets(nil)
	if err != nil {
		errMessage := "could not retrieve list of preset state ids"
		return &humanize_protobuf.ListPromptSegmentSetIdsResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.ListPromptSegmentSetIdsResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.ListPromptSegmentSetIdsResponse{
		PromptSegmentSetIds: promptSegmentSetIds,
		Status:              humanize_protobuf.ListPromptSegmentSetIdsResponse_SUCCESS,
	}, nil
}

// ------------ GET -----------------------------

func (c CreationApiImpl) GetPersonality(
	ctx context.Context, request *humanize_protobuf.GetPersonalityRequest,
) (*humanize_protobuf.GetPersonalityResponse, error) {
	personality, err := c.db.GetPersonality(request.PersonalityId, nil)
	if err != nil {
		errMessage := "could not retrieve personality with id: " + request.GetPersonalityId()
		return &humanize_protobuf.GetPersonalityResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.GetPersonalityResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.GetPersonalityResponse{
		Personality: personality,
		Status:      humanize_protobuf.GetPersonalityResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) GetEmotionalState(
	ctx context.Context, request *humanize_protobuf.GetEmotionalStateRequest,
) (*humanize_protobuf.GetEmotionalStateResponse, error) {
	emotionalState, err := c.db.GetEmotionalState(request.PresetStateId, nil)
	if err != nil {
		errMessage := "could not retrieve emotionalState with id: " + request.GetPresetStateId()
		return &humanize_protobuf.GetEmotionalStateResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.GetEmotionalStateResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.GetEmotionalStateResponse{
		EmotionalState: emotionalState,
		Status:         humanize_protobuf.GetEmotionalStateResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) GetGenerationConfig(
	ctx context.Context, request *humanize_protobuf.GetGenConfigRequest,
) (*humanize_protobuf.GetGenConfigResponse, error) {
	config, err := c.db.GetGenerationConfig(request.GetGenConfigId())
	if err != nil {
		errMessage := "could not retrieve generation config with id: " + request.GetGenConfigId()
		return &humanize_protobuf.GetGenConfigResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.GetGenConfigResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.GetGenConfigResponse{
		Config: config,
		Status: humanize_protobuf.GetGenConfigResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) GetPrompt(
	ctx context.Context, request *humanize_protobuf.GetPromptRequest,
) (*humanize_protobuf.GetPromptResponse, error) {
	prompt, err := c.db.GetPrompt(request.PromptId, nil)
	if err != nil {
		errMessage := "could not retrieve prompt with id: " + request.GetPromptId()
		return &humanize_protobuf.GetPromptResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.GetPromptResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.GetPromptResponse{
		Prompt: prompt,
		Status: humanize_protobuf.GetPromptResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) GetPromptSet(
	ctx context.Context, request *humanize_protobuf.GetPromptSetRequest,
) (*humanize_protobuf.GetPromptSetResponse, error) {
	promptSet, err := c.db.GetPromptSet(request.PromptSetId, nil)
	if err != nil {
		errMessage := "could not retrieve promptSet set with id: " + request.GetPromptSetId()
		return &humanize_protobuf.GetPromptSetResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.GetPromptSetResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.GetPromptSetResponse{
		Prompts: promptSet,
		Status:  humanize_protobuf.GetPromptSetResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) GetPromptSegmentSet(
	ctx context.Context, request *humanize_protobuf.GetPromptSegmentSetRequest,
) (*humanize_protobuf.GetPromptSegmentSetResponse, error) {
	promptSegments, err := c.db.GetPromptSegmentSet(request.PromptSegmentSetId, "", nil)
	if err != nil {
		errMessage := "could not retrieve promptSegme ts segment set with id: " + request.GetPromptSegmentSetId()
		return &humanize_protobuf.GetPromptSegmentSetResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.GetPromptSegmentSetResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.GetPromptSegmentSetResponse{
		PromptSegments: promptSegments,
		Status:         humanize_protobuf.GetPromptSegmentSetResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) GetPromptSegment(
	ctx context.Context, request *humanize_protobuf.GetPromptSegmentRequest,
) (*humanize_protobuf.GetPromptSegmentResponse, error) {
	prompt, err := c.db.GetPromptSegment(request.PromptSegmentId, nil)
	if err != nil {
		errMessage := "could not retrieve prompt segment with id: " + request.GetPromptSegmentId()
		return &humanize_protobuf.GetPromptSegmentResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.GetPromptSegmentResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.GetPromptSegmentResponse{
		PromptSegment: prompt,
		Status:        humanize_protobuf.GetPromptSegmentResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) GetActuationRuleSet(
	ctx context.Context, request *humanize_protobuf.GetActuationRuleSetRequest,
) (*humanize_protobuf.GetActuationRuleSetResponse, error) {
	rules, err := c.db.GetActuationRuleSet(request.GetPersonalityId(), nil)
	if err != nil {
		errMessage := "could not retrieve rules segment with id: " + request.GetPersonalityId()
		return &humanize_protobuf.GetActuationRuleSetResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.GetActuationRuleSetResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.GetActuationRuleSetResponse{
		Rules:  rules,
		Status: humanize_protobuf.GetActuationRuleSetResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) CreateGenerationConfig(
	ctx context.Context, request *humanize_protobuf.CreateGenerationConfigRequest,
) (*humanize_protobuf.CreateGenerationConfigResponse, error) {
	configId := c.idGen.GetULIDfromtime()
	err := c.db.CreateGenerationConfig(configId, request.Config, nil)
	if err != nil {
		errMessage := "could not create the generation config"
		return &humanize_protobuf.CreateGenerationConfigResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.CreateGenerationConfigResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.CreateGenerationConfigResponse{
		GenerationConfigId: configId,
		Status:             humanize_protobuf.CreateGenerationConfigResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) CreatePrompt(
	ctx context.Context, request *humanize_protobuf.CreatePromptRequest,
) (*humanize_protobuf.CreatePromptResponse, error) {
	promptId, err := c.db.CreatePrompt(request.Prompt, nil)
	if err != nil {
		errMessage := "could not create the prompt"
		return &humanize_protobuf.CreatePromptResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.CreatePromptResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.CreatePromptResponse{
		PromptId: promptId,
		Status:   humanize_protobuf.CreatePromptResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) CreatePersonalityRule(
	ctx context.Context, request *humanize_protobuf.CreatePersonalityRuleRequest,
) (*humanize_protobuf.CreatePersonalityRuleResponse, error) {
	ruleId, err := c.db.CreateEmotionalBoundRule(request.Rule)
	if err != nil {
		errMessage := "could not create the personality rule"
		return &humanize_protobuf.CreatePersonalityRuleResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.CreatePersonalityRuleResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.CreatePersonalityRuleResponse{
		Status: humanize_protobuf.CreatePersonalityRuleResponse_SUCCESS,
		RuleId: ruleId,
	}, nil
}

func (c CreationApiImpl) CreatePersonalityRulePart(
	ctx context.Context, request *humanize_protobuf.CreatePersonalityRulePartRequest,
) (*humanize_protobuf.CreatePersonalityRulePartResponse, error) {
	rulePartId, err := c.db.CreateEmotionalBoundRulePart(request.RulePart)
	if err != nil {
		errMessage := "could not create the personality rule"
		return &humanize_protobuf.CreatePersonalityRulePartResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.CreatePersonalityRulePartResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.CreatePersonalityRulePartResponse{
		Status:     humanize_protobuf.CreatePersonalityRulePartResponse_SUCCESS,
		RulePartId: rulePartId,
	}, nil
}

func (c CreationApiImpl) CreatePresetState(
	ctx context.Context, request *humanize_protobuf.CreatePresetStateRequest,
) (*humanize_protobuf.CreatePresetStateResponse, error) {
	err := c.db.CreateEmotionalState(request.State.EntityId, true, request.State, nil)
	if err != nil {
		errMessage := "could not create the personality rule"
		return &humanize_protobuf.CreatePresetStateResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.CreatePresetStateResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.CreatePresetStateResponse{
		Status: humanize_protobuf.CreatePresetStateResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) CreateEmotionalBound(
	ctx context.Context,
	request *humanize_protobuf.CreateEmotionalBoundRequest,
) (*humanize_protobuf.CreateEmotionalBoundResponse, error) {
	boundInstanceId, err := c.db.CreateEmotionalBound(request.Bound, "", nil)
	if err != nil {
		errMessage := "could not create the personality rule"
		return &humanize_protobuf.CreateEmotionalBoundResponse{
			ErrorMessage: errMessage,
			Status:       humanize_protobuf.CreateEmotionalBoundResponse_FAILED,
		}, errors.New(errMessage)
	}
	return &humanize_protobuf.CreateEmotionalBoundResponse{
		Status:          humanize_protobuf.CreateEmotionalBoundResponse_SUCCESS,
		BoundInstanceId: boundInstanceId,
	}, nil
}

func (c CreationApiImpl) UpdateEmotionalRule(
	ctx context.Context, request *humanize_protobuf.UpdateEmotionalRuleRequest,
) (*humanize_protobuf.UpdateEmotionalRuleResponse, error) {
	//TODO implement me
	return nil, nil
}

func (c CreationApiImpl) UpdateEmotionalRulePart(
	ctx context.Context, request *humanize_protobuf.UpdateEmotionalRulePartRequest,
) (*humanize_protobuf.UpdateEmotionalRulePartResponse, error) {
	return nil, nil
}

func (c CreationApiImpl) UpdateEmotionalState(
	ctx context.Context, request *humanize_protobuf.UpdateEmotionalStateRequest,
) (*humanize_protobuf.UpdateEmotionalStateResponse, error) {
	//TODO implement me
	return nil, nil
}

func (c CreationApiImpl) UpdateEmotionalBound(
	ctx context.Context, request *humanize_protobuf.UpdateEmotionalBoundRequest,
) (*humanize_protobuf.UpdateEmotionalBoundResponse, error) {
	//TODO implement me
	return nil, nil
}

func (c CreationApiImpl) UpdatePrompt(
	ctx context.Context, request *humanize_protobuf.UpdatePromptRequest,
) (*humanize_protobuf.UpdatePromptResponse, error) {
	err := c.db.UpdatePrompt(request.Prompt, nil)
	if err != nil {
		return &humanize_protobuf.UpdatePromptResponse{
			Status:       humanize_protobuf.UpdatePromptResponse_FAILED,
			ErrorMessage: "could  not update the prompt",
		}, nil
	}
	return &humanize_protobuf.UpdatePromptResponse{
		Status: humanize_protobuf.UpdatePromptResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) UpdatePromptSegment(
	ctx context.Context, request *humanize_protobuf.UpdatePromptSegmentRequest,
) (*humanize_protobuf.UpdatePromptSegmentResponse, error) {
	err := c.db.UpdatePromptSegment(request.PromptSegment, nil)
	if err != nil {
		return &humanize_protobuf.UpdatePromptSegmentResponse{
			Status:       humanize_protobuf.UpdatePromptSegmentResponse_FAILED,
			ErrorMessage: "could not update the prompt segment",
		}, nil
	}
	return &humanize_protobuf.UpdatePromptSegmentResponse{
		Status: humanize_protobuf.UpdatePromptSegmentResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) DeleteGenerationConfig(
	ctx context.Context, request *humanize_protobuf.DeleteGenerationConfigRequest,
) (*humanize_protobuf.DeleteGenerationConfigResponse, error) {
	if err := c.db.DeleteGenerationConfig(request.GenerationConfigId); err != nil {
		return &humanize_protobuf.DeleteGenerationConfigResponse{
			Status:       humanize_protobuf.DeleteGenerationConfigResponse_FAILED,
			ErrorMessage: "could not delete the generation config",
		}, nil
	}
	return &humanize_protobuf.DeleteGenerationConfigResponse{
		Status: humanize_protobuf.DeleteGenerationConfigResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) DeletePersonality(
	ctx context.Context, request *humanize_protobuf.DeletePersonalityRequest,
) (*humanize_protobuf.DeletePersonalityResponse, error) {
	if err := c.db.DeletePersonalityRule(request.GetPersonalityId(), nil); err != nil {
		return &humanize_protobuf.DeletePersonalityResponse{
			Status:       humanize_protobuf.DeletePersonalityResponse_FAILED,
			ErrorMessage: "could not delete the personality rule",
		}, nil
	}
	return &humanize_protobuf.DeletePersonalityResponse{
		Status: humanize_protobuf.DeletePersonalityResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) DeletePersonalityRule(
	ctx context.Context, request *humanize_protobuf.DeletePersonalityRuleRequest,
) (*humanize_protobuf.DeletePersonalityRuleResponse, error) {
	if err := c.db.DeletePersonalityRule(request.GetRuleId(), nil); err != nil {
		return &humanize_protobuf.DeletePersonalityRuleResponse{
			Status:       humanize_protobuf.DeletePersonalityRuleResponse_FAILED,
			ErrorMessage: "could not delete the personality rule",
		}, nil
	}
	return &humanize_protobuf.DeletePersonalityRuleResponse{
		Status: humanize_protobuf.DeletePersonalityRuleResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) DeletePersonalityRulePart(
	ctx context.Context, request *humanize_protobuf.DeletePersonalityRulePartRequest,
) (*humanize_protobuf.DeletePersonalityRulePartResponse, error) {
	if err := c.db.DeletePersonalityRulePart(request.GetRulePartId(), nil); err != nil {
		return &humanize_protobuf.DeletePersonalityRulePartResponse{
			Status:       humanize_protobuf.DeletePersonalityRulePartResponse_FAILED,
			ErrorMessage: "could not delete the personality rule part",
		}, nil
	}
	return &humanize_protobuf.DeletePersonalityRulePartResponse{
		Status: humanize_protobuf.DeletePersonalityRulePartResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) DeleteEmotionalState(
	ctx context.Context,
	request *humanize_protobuf.DeleteEmotionalStateRequest,
) (*humanize_protobuf.DeleteEmotionalStateResponse, error) {
	if err := c.db.DeleteEmotionalState(request.EmotionalStateRequest, nil); err != nil {
		return &humanize_protobuf.DeleteEmotionalStateResponse{
			Status:       humanize_protobuf.DeleteEmotionalStateResponse_FAILED,
			ErrorMessage: "could not delete the personality rule part",
		}, nil
	}
	return &humanize_protobuf.DeleteEmotionalStateResponse{
		Status: humanize_protobuf.DeleteEmotionalStateResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) DeletePrompt(
	ctx context.Context,
	request *humanize_protobuf.DeletePromptRequest,
) (*humanize_protobuf.DeletePromptResponse, error) {
	if err := c.db.DeletePrompt(request.PromptId, nil); err != nil {
		return &humanize_protobuf.DeletePromptResponse{
			Status:       humanize_protobuf.DeletePromptResponse_FAILED,
			ErrorMessage: "could not delete the prompt",
		}, nil
	}
	return &humanize_protobuf.DeletePromptResponse{
		Status: humanize_protobuf.DeletePromptResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) DeletePromptSegment(
	ctx context.Context,
	request *humanize_protobuf.DeletePromptSegmentRequest,
) (*humanize_protobuf.DeletePromptSegmentResponse, error) {
	if err := c.db.DeletePromptSegment(request.PromptSegmentId, nil); err != nil {
		return &humanize_protobuf.DeletePromptSegmentResponse{
			Status:       humanize_protobuf.DeletePromptSegmentResponse_FAILED,
			ErrorMessage: "could not delete the prompt",
		}, nil
	}
	return &humanize_protobuf.DeletePromptSegmentResponse{
		Status: humanize_protobuf.DeletePromptSegmentResponse_SUCCESS,
	}, nil
}

func (c CreationApiImpl) DeleteEmotionalBound(
	ctx context.Context,
	request *humanize_protobuf.DeleteEmotionalBoundRequest,
) (*humanize_protobuf.DeleteEmotionalBoundResponse, error) {
	//TODO implement me
	return nil, nil
}
