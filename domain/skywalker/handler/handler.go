package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"rings/domain/skywalker/domain"
	"rings/domain/skywalker/service"
	"strings"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func A() string {
	return "hola"
}

func (h *Handler) HandleRequest(g *gin.Context) {
	response, err := h.service.GetMessage()
	if err != nil {
		logrus.Errorf("Error getting message, err: %s", err.Error())
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	responseStr := h.generateResponse(response)
	g.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": responseStr,
	})
	return
}

func (h *Handler) generateResponse(response *domain.Response) string {
	responseStr := ""
	for _, character := range response.Characters {
		filmsStr := strings.Join(character.Films[:len(character.Films)-1], ", ")
		if len(character.Films) != 1 {
			filmsStr = fmt.Sprintf("%s and %s", filmsStr, character.Films[len(character.Films)-1])
		}

		responseStr += fmt.Sprintf("%s participated in %s. ", character.Name, filmsStr)
	}

	return responseStr
}
