package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/iamvineettiwari/go-web-server/data"
	"github.com/iamvineettiwari/go-web-server/http"
	"github.com/iamvineettiwari/go-web-server/http/status"
)

func GetUser(req *http.Request, res *http.Response) {
	db := data.GetDB()

	var users any

	id, isIdPresent := req.Query["id"]

	if isIdPresent {
		id, err := strconv.Atoi(id.(string))

		if err != nil {
			res.Send([]byte(err.Error()), status.HTTP_500_SERVER_ERROR)
			return
		}

		users = db.GetById(id)
	} else {
		users = db.Get()
	}

	data, err := json.Marshal(users)

	if err != nil {
		res.Send([]byte(err.Error()), status.HTTP_500_SERVER_ERROR)
		return
	}

	res.AddHeader("Content-Type", http.APPLICATION_JSON)
	res.Send(data, status.HTTP_200_OK)
}

func CreateUser(req *http.Request, res *http.Response) {
	var user data.User

	err := json.Unmarshal(req.Body, &user)

	if err != nil {
		res.SetStatus(status.HTTP_400_BAD_REQUEST)
		res.Write([]byte(err.Error()))
		return
	}

	db := data.GetDB()

	userInserted := db.Insert(&user)
	respData, err := json.Marshal(userInserted)

	if err != nil {
		res.SetStatus(status.HTTP_400_BAD_REQUEST)
		res.Write([]byte(err.Error()))
		return
	}

	res.AddHeader("Content-Type", http.APPLICATION_JSON)
	res.Send(respData, status.HTTP_201_CREATED)
}

func UpdateUser(req *http.Request, res *http.Response) {
	var user data.User

	err := json.Unmarshal(req.Body, &user)

	if err != nil {
		res.SetStatus(status.HTTP_400_BAD_REQUEST)
		res.Write([]byte(err.Error()))
		return
	}

	db := data.GetDB()

	if db.GetById(user.Id) == nil {
		res.SetStatus(status.HTTP_400_BAD_REQUEST)
		res.Write([]byte("User Not Found"))
		return
	}

	userUpdated := db.Update(&user)
	respData, err := json.Marshal(userUpdated)

	if err != nil {
		res.SetStatus(status.HTTP_400_BAD_REQUEST)
		res.Write([]byte(err.Error()))
		return
	}

	res.AddHeader("Content-Type", http.APPLICATION_JSON)
	res.Send(respData, status.HTTP_200_OK)
}

func DeleteUser(req *http.Request, res *http.Response) {
	id, isIdPresent := req.Query["id"]

	if !isIdPresent {
		res.SetStatus(status.HTTP_400_BAD_REQUEST)
		res.Write([]byte("Please provide id in query parameter"))
		return
	}

	userId, err := strconv.Atoi(id.(string))

	if err != nil {
		res.Send([]byte(err.Error()), status.HTTP_500_SERVER_ERROR)
		return
	}

	db := data.GetDB()
	db.Delete(userId)
	res.AddHeader("Content-Type", http.APPLICATION_JSON)
	res.Send(nil, status.HTTP_204_NO_CONTENT)
}
