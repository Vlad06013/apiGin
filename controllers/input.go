package controllers

import (
	"github.com/Vlad06013/apiGin/models/tgObjects/Input"
	"github.com/Vlad06013/apiGin/pkg/telegram"
	"github.com/gin-gonic/gin"
)

func InputTGRequest(c *gin.Context) {

	var i = Input.TextMessage{}
	//var callBackQueryMessage tgObjects.CallBackQuery
	if err := c.ShouldBindJSON(&i); err == nil {

		var inputTextMessage Input.Telegram = i
		var answer = inputTextMessage.GenerateAnswer()
		telegram.Start(answer)
		c.JSON(200, gin.H{
			//"input":  inputTextMessage,
			"answer": &answer,
		})
	}
	//if err := c.ShouldBindJSON(&callBackQueryMessage); err != nil {
	//}

	//setInputObject(&input)

	//if input.Message != nil {
	//	var obj = *input.Message
	//c.JSON(200, gin.H{"message": obj})
	//telegram.Start(&obj)

	//}

	//c.JSON(200, gin.H{"input": i})

	//var message []migrations.Message
	////models.DB.Find(&bots)
	//models.ConnectDB().Find(&message)
	//
	//c.JSON(http.StatusOK, gin.H{"message": message})
}

//
//func setInputObject(i *tgObjects.InputObject) {
//
//	fmt.Println(i)
//}
