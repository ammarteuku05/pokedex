package dto

type MonsterInput struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Statistics  *Statistics `json:"statistics"`
	Kind        string      `json:"kind"`
	TypeId      []string    `json:"typeId"`
}

type MonsterInputUpdate struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Statistics  *Statistics `json:"statistics"`
	Kind        string      `json:"kind"`
	TypeId      string      `json:"typeId"`
	TypeIdOld   string      `json:"typeIdOld"`
}

type MonsterResponse struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Statistics  *Statistics `json:"statistics"`
	Kind        string      `json:"kind"`
	TypeId      int         `json:"type_id"`
	TypeName    string      `json:"type_name"`
	CreatedAt   string
	UpdatedAt   string
}

type Statistics struct {
	Hp     int `json:"hp"`
	Def    int `json:"def"`
	Speed  int `json:"speed"`
	Attack int `json:"attack"`
}

type FilterType struct {
	Id []int `json:"id"`
}

type FilterMonster struct {
	Sort  string
	Order string
	Types []string
	Id    string
}

type ResponseType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
