package mapper

import (
	"encoding/json"
	"pokedex/entity"
	"pokedex/shared/dto"
)

func MapResponseMonster(monster *entity.Monster) *dto.MonsterResponse {
	dataStats := []byte(monster.Statistics)
	var payloadStats *dto.Statistics
	_ = json.Unmarshal(dataStats, &payloadStats)
	return &dto.MonsterResponse{
		Id:          monster.ID,
		Name:        monster.Name,
		Description: monster.Description,
		Statistics:  payloadStats,
		Kind:        monster.Kind,
		CreatedAt:   monster.CreatedAt.GoString(),
		UpdatedAt:   monster.UpdatedAt.GoString(),
	}
}
func MapResponseMonsterDetail(monster *entity.MonsterDetail) *dto.MonsterResponse {
	dataStats := []byte(monster.Statistics)
	var payloadStats *dto.Statistics
	_ = json.Unmarshal(dataStats, &payloadStats)
	return &dto.MonsterResponse{
		Id:          monster.ID,
		Name:        monster.Name,
		Description: monster.Description,
		TypeId:      monster.TypeId,
		TypeName:    monster.TypeName,
		Statistics:  payloadStats,
		Kind:        monster.Kind,
		CreatedAt:   monster.CreatedAt,
		UpdatedAt:   monster.UpdatedAt,
	}
}

func MapResponseTypes(types *entity.Type) dto.ResponseType {
	return dto.ResponseType{
		Id:   types.ID,
		Name: types.Name,
	}
}
