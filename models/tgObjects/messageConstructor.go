package tgObjects

import (
	"github.com/Vlad06013/apiGin/models"
)

type MessageConstructor struct {
	Text     string
	Type     string
	Keyboard models.Keyboard
}
