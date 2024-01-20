package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/iamvineettiwari/go-web-server/data"
	"github.com/iamvineettiwari/go-web-server/http"
	"github.com/iamvineettiwari/go-web-server/http/status"
)

func GetUserById(req *http.Request, res *http.Response) {
	db := data.GetDB()

	id, present := req.Params["id"]

	if !present {
		res.SetStatus(status.HTTP_400_BAD_REQUEST)
		res.Write([]byte("User Not Found"))
		return
	}

	userId, err := strconv.Atoi(id)

	if err != nil {
		res.Send([]byte(err.Error()), status.HTTP_500_SERVER_ERROR)
		return
	}

	user := db.GetById(userId)

	data, err := json.Marshal(user)

	if err != nil {
		res.Send([]byte(err.Error()), status.HTTP_500_SERVER_ERROR)
		return
	}

	res.AddHeader("Content-Type", http.APPLICATION_JSON)
	res.Send(data, status.HTTP_200_OK)
}

func GetUsers(req *http.Request, res *http.Response) {
	db := data.GetDB()

	users := db.Get()

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

	id, present := req.Params["id"]

	if !present {
		res.SetStatus(status.HTTP_400_BAD_REQUEST)
		res.Write([]byte("User Not Found"))
		return
	}

	userId, err := strconv.Atoi(id)

	if err != nil {
		res.Send([]byte(err.Error()), status.HTTP_500_SERVER_ERROR)
		return
	}

	err = json.Unmarshal(req.Body, &user)

	if err != nil {
		res.SetStatus(status.HTTP_400_BAD_REQUEST)
		res.Write([]byte(err.Error()))
		return
	}

	db := data.GetDB()

	if db.GetById(userId) == nil {
		res.SetStatus(status.HTTP_400_BAD_REQUEST)
		res.Write([]byte("User Not Found"))
		return
	}

	user.Id = userId
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
	id, isIdPresent := req.Params["id"]

	if !isIdPresent {
		res.SetStatus(status.HTTP_400_BAD_REQUEST)
		res.Write([]byte("Please provide id in query parameter"))
		return
	}

	userId, err := strconv.Atoi(id)

	if err != nil {
		res.Send([]byte(err.Error()), status.HTTP_500_SERVER_ERROR)
		return
	}

	db := data.GetDB()
	db.Delete(userId)
	res.AddHeader("Content-Type", http.APPLICATION_JSON)
	res.Send(nil, status.HTTP_204_NO_CONTENT)
}
