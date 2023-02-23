package repository

import (
	"fmt"
	"pokedex/entity"
	"pokedex/shared/dto"
	"strings"

	"gorm.io/gorm"
)

//go:generate mockery --name MonsterRepository --case snake --output ../mocks/mocksRepository --disable-version-string
type MonsterRepository interface {
	CreateMonster(monster *entity.Monster) (*entity.Monster, error)
	FindByIdType(string) (*entity.Type, error)
	UpdateMonster(idMonster int, dataUpdates map[string]interface{}) (*entity.Monster, error)
	UpdateTypeOfMonster(mosterId, typeId, typeIdOld string) (string, error)
	FindMonsterById(monsterId int) (*entity.Monster, error)
	DeleteMonster(idMonster, typeId int) (string, error)
	FindAllType() (*[]entity.Type, error)
	CreateTypeOfMonster(monster *entity.TypeOfMonster) (string, error)
	FindAllMonsterByFilter(filter *dto.FilterMonster) ([]entity.MonsterDetail, error)
}

type monsterRepository struct {
	db *gorm.DB
}

func NewMonsterRepository(db *gorm.DB) *monsterRepository {
	return &monsterRepository{db}
}

func (m *monsterRepository) CreateMonster(monster *entity.Monster) (*entity.Monster, error) {
	if err := m.db.Create(&monster).Error; err != nil {
		return nil, err
	}
	return monster, nil
}

func (m *monsterRepository) CreateTypeOfMonster(monster *entity.TypeOfMonster) (string, error) {
	if err := m.db.Create(&monster).Error; err != nil {
		return "", err
	}
	return "success to create type of monster", nil
}

func (m *monsterRepository) FindByIdType(id string) (*entity.Type, error) {
	var types *entity.Type

	if err := m.db.Where("id = ?", id).Find(&types).Error; err != nil {
		return nil, err
	}

	return types, nil
}

func (m *monsterRepository) FindMonsterById(monsterId int) (*entity.Monster, error) {
	var monster *entity.Monster

	if err := m.db.Where("id = ?", monsterId).Find(&monster).Error; err != nil {
		return nil, err
	}

	return monster, nil
}

func (m *monsterRepository) UpdateMonster(idMonster int, dataUpdates map[string]interface{}) (*entity.Monster, error) {
	var (
		monster *entity.Monster
	)

	if err := m.db.Model(monster).Where("id = ?", idMonster).Updates(dataUpdates).Error; err != nil {
		return nil, err
	}

	if err := m.db.Where("id = ?", idMonster).Find(&monster).Error; err != nil {
		return nil, err
	}

	return monster, nil
}

func (m *monsterRepository) DeleteMonster(idMonster, idType int) (string, error) {
	var (
		typeOfMonster *entity.TypeOfMonster
		monster       *entity.Monster
	)

	if err := m.db.Model(typeOfMonster).Where("monster_id = ? and type_id = ? ", idMonster, idType).Delete(&typeOfMonster).Error; err != nil {
		return "", err
	}

	if err := m.db.Model(typeOfMonster).Where("monster_id = ?", typeOfMonster.MonsterId).Scan(&typeOfMonster).Error; err == nil {
		if typeOfMonster.MonsterId == 0 {
			if err := m.db.Model(monster).Where("id = ?", idMonster).Delete(&monster).Error; err == nil {
				return "", err
			}
		}
	}

	return "success to delete", nil
}

func (m *monsterRepository) FindAllType() (*[]entity.Type, error) {
	var (
		types *[]entity.Type
	)
	err := m.db.Find(&types).Error
	if err != nil {
		return nil, err
	}

	return types, nil
}

func (m *monsterRepository) UpdateTypeOfMonster(mosterId, typeId, typeIdOld string) (string, error) {
	var typeOfMonster *entity.TypeOfMonster
	if err := m.db.Model(&typeOfMonster).Where("type_id = ? AND monster_id = ?", typeIdOld, mosterId).Update("type_id", typeId).Error; err != nil {
		return "", err
	}
	return "success update type of monster", nil
}

func (m *monsterRepository) FindAllMonsterByFilter(filter *dto.FilterMonster) ([]entity.MonsterDetail, error) {
	var (
		payload []entity.MonsterDetail
		idStr   string
		typeStr string
	)

	if filter.Order == "" {
		filter.Order = "id"
	}
	if filter.Sort == "" {
		filter.Sort = "ASC"
	}

	if filter.Id != "" {
		idStr = fmt.Sprintf(" AND x.monster_id = %s ", filter.Id)
	} else {
		idStr = ""
	}
	if len(filter.Types) > 0 {
		typeFlag := strings.Join(filter.Types, ",")
		typeStr = ` WHERE x.type_id IN (` + typeFlag + `) `
	} else {
		typeStr = ""
	}

	strOrder := fmt.Sprintf("Monsters.%s %s", filter.Order, filter.Sort)

	if err := m.db.Raw("SELECT x.*, Types.name AS TypeName, Monsters.* FROM TypeOfMonster x JOIN Monsters ON Monsters.id = x.monster_id JOIN Types ON Types.id = x.type_id" + typeStr + idStr + " ORDER BY " + strOrder).Find(&payload).Error; err != nil {
		return nil, err
	}

	return payload, nil
}
