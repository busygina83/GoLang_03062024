/*
// Пример REST сервера с несколькими маршрутами(используем только стандартную библиотеку)

// POST   /task/              :  создаёт задачу и возвращает её ID
// GET    /task/<taskid>      :  возвращает одну задачу по её ID
// GET    /task/              :  возвращает все задачи
// DELETE /task/<taskid>      :  удаляет задачу по ID
// DELETE /task/              :  удаляет все задачи
// GET    /tag/<tagname>      :  возвращает список задач с заданным тегом
// GET    /due/<yy>/<mm>/<dd> :  возвращает список задач, запланированных на указанную дату

Структура проекта
https://github.com/golang-standards/project-layout/blob/master/README_ru.md
*/

package main

import (
	"TaskStoreAPI_Gin/internal/taskstore"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

func (ts *taskServer) createTaskHandler(context *gin.Context) {

	// Структура для создания задачи
	type RequestTask struct {
		Text string    `json:"text"`
		Tags []string  `json:"tags"`
		Due  time.Time `json:"due"`
	}

	var rt RequestTask
	// Десериализация JSON -> RequestTask
	if err := context.ShouldBindJSON(&rt); err != nil {
		context.String(http.StatusBadRequest, err.Error())
	}

	// Создаем новую задачу в хранилище и получаем ее <id>
	id := ts.store.CreateTask(rt.Text, rt.Tags, rt.Due)
	context.JSON(http.StatusOK, gin.H{"Id": id})

}

func (ts *taskServer) getAllTaskHandler(context *gin.Context) {
	allTasks := ts.store.GetAllTasks()
	context.JSON(http.StatusOK, allTasks)
}

func (ts *taskServer) getTaskHandler(context *gin.Context) {
	id, err := strconv.Atoi(context.Params.ByName("id"))

	if err != nil {
		context.String(http.StatusBadRequest, err.Error()) // код ошибки 400
		return
	}
	task, err := ts.store.GetTask(id)
	if err != nil {
		context.String(http.StatusNotFound, err.Error()) // код ошибки 404
		return
	}

	context.JSON(http.StatusOK, task)
}

func (ts *taskServer) deleteTaskHandler(context *gin.Context) {
	id, err := strconv.Atoi(context.Params.ByName("id"))

	if err != nil {
		context.String(http.StatusBadRequest, err.Error()) // код ошибки 400
		return
	}

	err = ts.store.DeleteTask(id)
	if err != nil {
		context.String(http.StatusNotFound, err.Error()) // код ошибки 404
		return
	}

}

func (ts *taskServer) deleteAllTaskHandler(context *gin.Context) {
	log.Printf("Handling delete all tasks at %s\n", context.FullPath())

	ts.store.DeleteAllTasks()
}

func main() {
	router := gin.Default()
	server := NewTaskServer()

	// Added routing for "/task/" URL
	router.GET("/task/", server.getAllTaskHandler)
	router.GET("/task/:id", server.getTaskHandler)
	router.POST("/task/", server.createTaskHandler)
	router.DELETE("/task/", server.deleteAllTaskHandler)
	router.DELETE("/task/:id", server.deleteTaskHandler)

	router.Run(":8080")
}