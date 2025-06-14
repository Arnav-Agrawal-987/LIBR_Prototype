package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Arnav-Agrawal-987/Libr_Prototype/db"
	"github.com/Arnav-Agrawal-987/Libr_Prototype/model"
	"github.com/Arnav-Agrawal-987/Libr_Prototype/moderator"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func moderateMessage(messageContent string) model.Message {
	var message model.Message
	message.ID = uuid.New()
	message.Content = messageContent
	message.Timestamp = int(time.Now().Unix())
	modRespChan := make(chan model.ModResponse, 3)
	wg := sync.WaitGroup{}
	response := 0
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(modID int, message model.Message) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()
			moderator.Validate(ctx, message, modRespChan, modID)
		}(i, message)
	}
	wg.Wait()
	close(modRespChan)
	for modResp := range modRespChan {
		response += modResp.Status
	}
	if response == 2 {
		message.Status = "Approved"
	} else {
		message.Status = "Rejected"
	}
	return message
}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	var input model.Message
	json.NewDecoder(r.Body).Decode(&input)
	message := moderateMessage(input.Content)
	if message.Status == "Approved" {
		_, err := db.Pool.Exec(context.Background(), `INSERT INTO MESSAGES (id, content, timestamp, status) VALUES($1,$2,$3,$4)`, message.ID, message.Content, message.Timestamp, message.Status)
		if err != nil {
			log.Println(err)
		}
	}
	apires := model.ApiResponse{
		ID:        message.ID,
		Timestamp: message.Timestamp,
		Status:    message.Status,
	}
	json.NewEncoder(w).Encode(apires)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	timestamp, _ := strconv.Atoi(params["timestamp"])
	rows, err := db.Pool.Query(context.Background(), `SELECT * FROM MESSAGES WHERE TIMESTAMP=$1`, timestamp)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var messages []model.Message
	for rows.Next() {
		var message model.Message
		err := rows.Scan(&message.ID, &message.Content, &message.Timestamp, &message.Status)
		if err != nil {
			log.Println(err)
		}
		messages = append(messages, message)
	}
	json.NewEncoder(w).Encode(messages)
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Pool.Query(context.Background(), `SELECT * FROM MESSAGES`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var messages []model.Message
	for rows.Next() {
		var message model.Message
		err := rows.Scan(&message.ID, &message.Content, &message.Timestamp, &message.Status)
		if err != nil {
			log.Println(err)
		}
		messages = append(messages, message)
	}
	json.NewEncoder(w).Encode(messages)
}
