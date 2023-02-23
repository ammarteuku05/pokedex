package entity

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(32);not null"`
	Password  string `gorm:"type:varchar(100);not null"`
	Role      string `gorm:"type:varchar(8);not null"`
	Email     string `gorm:"type:varchar(50);not null;uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "Users"
}

type Monster struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(32);not null;;uniqueIndex"`
	Description string
	Statistics  JSON
	Kind        string `gorm:"type:varchar(50);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Monster) TableName() string {
	return "Monsters"
}

type Type struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(32);not null"`
}

func (Type) TableName() string {
	return "Types"
}

type TypeOfMonster struct {
	MonsterId int
	Monster   Monster
	TypeId    int
	Type      Type
}

func (TypeOfMonster) TableName() string {
	return "TypeOfMonster"
}

type MonsterDetail struct {
	ID          int
	Name        string
	Description string
	Statistics  JSON
	Kind        string
	TypeId      int
	TypeName    string
	CreatedAt   string
	UpdatedAt   string
}
