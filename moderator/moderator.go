package moderator

import (
	"context"
	"log"
	"math/rand/v2"
	"time"

	"github.com/Arnav-Agrawal-987/Libr_Prototype/model"
)

func Validate(ctx context.Context, message model.Message, modRespChan chan model.ModResponse, modID int) {
	delay := rand.IntN(3) + 1
	var status int
	select {
	case <-time.After(time.Second * time.Duration(delay)):
		status = rand.IntN(2)
	case <-ctx.Done():
		log.Printf("Moderator %d timed out.", modID)
		status = 0
	}
	modRes := model.ModResponse{
		ModID:        modID,
		MessageID:    message.ID,
		Status:       status,
		ResponseTime: delay,
	}
	modRespChan <- modRes
}
