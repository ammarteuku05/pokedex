package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"pokedex/entity"
	"pokedex/repository"
	"pokedex/shared/dto"
	"pokedex/shared/mapper"
	"strconv"
	"time"
)

//go:generate mockery --name MonsterServices --case snake --output ../mocks/mocksServices --disable-version-string
type MonsterServices interface {
	CreateMonster(req *dto.MonsterInput) (string, error)
	UpdateMonster(id int, req *dto.MonsterInputUpdate) (string, error)
	DeleteMonster(monsterId, typeId int) (string, error)
	FindByIdMonster(id int) (*dto.MonsterResponse, error)
	FindAllType() ([]dto.ResponseType, error)
	FindAllMonsters(filter *dto.FilterMonster) ([]dto.MonsterResponse, error)
}

type monsterService struct {
	monsterRepository repository.MonsterRepository
}

func NewMonsterService(monsterRepository repository.MonsterRepository) *monsterService {
	return &monsterService{monsterRepository}
}

func (m *monsterService) CreateMonster(req *dto.MonsterInput) (string, error) {
	metaDataStatistics := map[string]interface{}{
		"hp":     req.Statistics.Hp,
		"def":    req.Statistics.Def,
		"attack": req.Statistics.Attack,
		"speed":  req.Statistics.Speed,
	}

	metadataJSON, _ := json.Marshal(metaDataStatistics)

	newMonster := entity.Monster{
		Name:        req.Name,
		Description: req.Description,
		Statistics:  metadataJSON,
		Kind:        req.Kind,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	monsterCreate, err := m.monsterRepository.CreateMonster(&newMonster)

	if err != nil {
		return "", err
	}

	for _, id := range req.TypeId {
		typeId, err := m.monsterRepository.FindByIdType(id)
		if err != nil {
			return "", err
		}

		if typeId.ID == 0 {
			return "", errors.New("not found type")
		}

		typeOfMonster := &entity.TypeOfMonster{
			MonsterId: monsterCreate.ID,
			TypeId:    typeId.ID,
		}
		_, err = m.monsterRepository.CreateTypeOfMonster(typeOfMonster)

		if err != nil {
			return "", err
		}

	}

	fmt.Println(len(req.TypeId))

	return "success to create monster", nil
}

func (m *monsterService) UpdateMonster(id int, req *dto.MonsterInputUpdate) (string, error) {
	dataUpdate := map[string]interface{}{}

	metaDataStatistics := map[string]interface{}{
		"hp":     req.Statistics.Hp,
		"def":    req.Statistics.Def,
		"attack": req.Statistics.Attack,
		"speed":  req.Statistics.Speed,
	}

	metadataJSON, _ := json.Marshal(metaDataStatistics)

	monsterId, err := m.monsterRepository.FindMonsterById(id)

	if monsterId.ID == 0 {
		newError := fmt.Sprintf("category id %v not found", monsterId)
		return "", errors.New(newError)
	}

	if err != nil {
		return "", err
	}

	if req.Name != "" {
		dataUpdate["name"] = req.Name
	}
	if req.Description != "" {
		dataUpdate["description"] = req.Description
	}
	if req.Statistics != nil {
		dataUpdate["statistics"] = metadataJSON
	}
	if req.Kind != "" {
		dataUpdate["kind"] = req.Kind
	}

	_, err = m.monsterRepository.UpdateMonster(id, dataUpdate)

	if err != nil {
		return "", err
	}

	typeId, err := m.monsterRepository.FindByIdType(req.TypeIdOld)
	if err != nil {
		return "", err
	}

	if typeId.ID == 0 {
		return "", errors.New("not found type")
	}

	strMonsterId := strconv.Itoa(monsterId.ID)

	_, err = m.monsterRepository.UpdateTypeOfMonster(strMonsterId, req.TypeId, req.TypeIdOld)

	fmt.Print(err)

	if err != nil {
		return "", err
	}

	return "succes to update", err
}

func (m *monsterService) DeleteMonster(monsterId, typeId int) (string, error) {
	resp, err := m.monsterRepository.DeleteMonster(monsterId, typeId)

	if err != nil {
		return "", err
	}

	return resp, nil
}

func (m *monsterService) FindByIdMonster(id int) (*dto.MonsterResponse, error) {
	res, err := m.monsterRepository.FindMonsterById(id)

	if err != nil {
		return nil, err
	}

	return mapper.MapResponseMonster(res), nil

}

func (m *monsterService) FindAllType() ([]dto.ResponseType, error) {
	var response []dto.ResponseType
	types, err := m.monsterRepository.FindAllType()

	if err != nil {
		return nil, err
	}

	for _, typeDetail := range *types {
		res := mapper.MapResponseTypes(&typeDetail)
		response = append(response, res)
	}

	return response, nil
}

func (m *monsterService) FindAllMonsters(filter *dto.FilterMonster) ([]dto.MonsterResponse, error) {
	var (
		response []dto.MonsterResponse
	)
	res, err := m.monsterRepository.FindAllMonsterByFilter(filter)

	if err != nil {
		return nil, errors.New("error disini")
	}

	for i := 0; i < len(res); i++ {
		responses := mapper.MapResponseMonsterDetail(&res[i])
		response = append(response, *responses)
	}

	return response, nil
}
