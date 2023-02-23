package controller

import (
	"pokedex/service"
	"pokedex/shared/dto"
	"pokedex/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type monsterController struct {
	monsterService service.MonsterServices
}

func NewMonsterController(userMonster service.MonsterServices) *monsterController {
	return &monsterController{userMonster}
}

func (h *monsterController) CreateMonsterController(c *gin.Context) {
	var inputMonster *dto.MonsterInput

	if err := c.ShouldBindJSON(&inputMonster); err != nil {
		splitError := utils.SplitErrorInformation(err)
		responseError := utils.APIResponse(400, "Input data required", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newMonster, err := h.monsterService.CreateMonster(inputMonster)
	if err != nil {
		responseError := utils.APIResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := utils.APIResponse(201, "Create new user succeed", newMonster)
	c.JSON(201, response)
}

func (h *monsterController) UpdateMonsterController(c *gin.Context) {
	monsterId := c.Param("id")
	var updateMonster *dto.MonsterInputUpdate

	if err := c.ShouldBindJSON(&updateMonster); err != nil {
		splitError := utils.SplitErrorInformation(err)
		responseError := utils.APIResponse(400, "Input data required", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	intMonsterId, _ := strconv.Atoi(monsterId)

	newMonster, err := h.monsterService.UpdateMonster(intMonsterId, updateMonster)
	if err != nil {
		responseError := utils.APIResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := utils.APIResponse(200, "success updated monster", newMonster)
	c.JSON(200, response)
}

func (h *monsterController) DeleteMonsterController(c *gin.Context) {
	monsterId := c.Query("monster_id")
	typeId := c.Query("type_id")

	intMonsterId, _ := strconv.Atoi(monsterId)
	intTypeId, _ := strconv.Atoi(typeId)

	checkMonster, err := h.monsterService.FindByIdMonster(intMonsterId)

	if err != nil {
		responseError := utils.APIResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	if checkMonster.Id == 0 {
		responseError := utils.APIResponse(404, "not found", nil)

		c.JSON(404, responseError)
		return
	}

	resMonster, err := h.monsterService.DeleteMonster(intMonsterId, intTypeId)

	if err != nil {
		responseError := utils.APIResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := utils.APIResponse(200, "success deleted monster", resMonster)
	c.JSON(200, response)
}

func (h *monsterController) FindAllType(c *gin.Context) {
	checkMonster, err := h.monsterService.FindAllType()

	if err != nil {
		responseError := utils.APIResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	if len(checkMonster) == 0 {
		responseError := utils.APIResponse(404, "not found", nil)

		c.JSON(404, responseError)
		return
	}

	response := utils.APIResponse(200, "Success", checkMonster)
	c.JSON(201, response)
}

func (h *monsterController) FindAllMonsters(c *gin.Context) {
	sort := c.Query("sort")
	order := c.Query("order")
	id := c.Query("id")
	types := c.QueryArray("types")
	req := dto.FilterMonster{
		Sort:  sort,
		Order: order,
		Types: types,
		Id:    id,
	}

	checkMonster, err := h.monsterService.FindAllMonsters(&req)

	if err != nil {
		responseError := utils.APIResponse(500, "Internal server error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	if len(checkMonster) == 0 {
		responseError := utils.APIResponse(404, "not found", nil)

		c.JSON(404, responseError)
		return
	}

	response := utils.APIResponse(200, "Success", checkMonster)
	c.JSON(200, response)
}
