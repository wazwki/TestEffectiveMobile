package handlers

import "net/http"

func GetSongHandler(w http.ResponseWriter, r *http.Request) {
	//вывод всей библиотеки(только поля group и song) с пагинацией и фильтрами
}

func GetDetailSongHandler(w http.ResponseWriter, r *http.Request) {
	//вывод всех полей конкретного объекта
}

func PostSongHandler(w http.ResponseWriter, r *http.Request) {
	//создание объекта с полями group и song и добавление оставшихся полей из внешнего api
}

func UpdateSongHandler(w http.ResponseWriter, r *http.Request) {
	//обновление объекта по id
}

func DeleteSongHandler(w http.ResponseWriter, r *http.Request) {
	//удаление объекта по id
}
