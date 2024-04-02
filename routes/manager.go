package routes

import (
	"fmt"
	"khojkhaz-server/models"
	"khojkhaz-server/storage"
	"khojkhaz-server/utils"
	"strings"

	"github.com/kataras/iris/v12"
)

func CreateManager(ctx iris.Context) {
	const maxSize = 10 * iris.MB
	ctx.SetMaxRequestBodySize(maxSize)

	var managerInput ManagerInput
	err := ctx.ReadJSON(&managerInput)
	if err != nil {
		utils.HandleValidationErrors(err, ctx)
		return
	}

	var url string = ""
	if managerInput.Image != "" {
		res := storage.UploadBase64Image(
			managerInput.Image,
			strings.ReplaceAll(fmt.Sprint(managerInput.UserID)+ "/"+managerInput.Name, " ", ""), 
		)

		url = res["url"]
	}

	manager := models.Manager{
		Name: managerInput.Name,
		UserID: managerInput.UserID,
		Email: managerInput.Email,
		PhoneNumber: managerInput.PhoneNumber,
		Website: managerInput.Website,
		Image: url,
	}

	storage.DB.Create(&manager)
}


type ManagerInput struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phoneNumber" `
	Website     string `json:"website"`
	Image       string `json:"image" validate:"required"`
	UserID      string `json:"userID"`

}
