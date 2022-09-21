package main

import (
	"musicstore_rest_api/controllers"
	"musicstore_rest_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// инициализируем новый маршрутизатор Gin
	route := gin.Default()
	
	// подключение к БД
	models.ConnectDB() // new

	// Далее определим GET маршрут до пустого endpoint.Чтобы определить маршрут, нам нужно указать две вещи:
	// 1.Конечную точку (endpoint) — это путь, который клиент хочет получить.
	// 2.Обработчик — он определяет, как сервер предоставляет данные клиенту.
	// Тут находится бизнес-логика, такая как получение данных из базы данных, проверка ввода пользователя и так далее.
	// Клиенту можно отправить несколько типов ответов,
	// но RESTful API обычно дают ответ в формате JSON. Для этого в Gin мы используем метод JSON, в контексте запроса.

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	// зарегистрируем функции из contollers в качестве обработчика маршрута
	   // Маршруты
	   route.GET("/tracks", controllers.GetAllTracks)
	   route.POST("/tracks", controllers.CreateTrack)
	   route.GET("/tracks/:id", controllers.GetTrack)
	   route.PATCH("/tracks/:id", controllers.UpdateTrack)
	   route.DELETE("/tracks/:id", controllers.DeleteTrack)

	// Запускаем сервер, просто вызывав метод Run нашего экземпляра Gin
	route.Run()
}

// Команда в терминале - $ go run main.go
