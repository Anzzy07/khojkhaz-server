package routes

import (
	"khojkhaz-server/utils"

	"github.com/kataras/iris/v12"
)

func TestMessageNotification(ctx iris.Context) {
	data := map[string]string{
		"url": "exp://192.168.2.98:8081/--/messages/101",
	}

	err := utils.SendNotification(
		"ExponentPushToken[AkaaapCobSlZW5xJUvH0ZQ]",
		"Push Title", "Push body is this message", data)
	if err != nil {
		utils.CreateInternalServerError(ctx)
		return
	}

	ctx.JSON(iris.Map{
		"sent": true,
	})
}