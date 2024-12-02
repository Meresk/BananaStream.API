package controllers

import (
	"BananaStream.API/config"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
)

func Rooms(c *fiber.Ctx) error {
	roomClient := lksdk.NewRoomServiceClient(config.LivekitServerURL, config.ApiKey, config.ApiSecret)
	rooms, _ := roomClient.ListRooms(context.Background(), &livekit.ListRoomsRequest{})

	roomDto := make([]fiber.Map, len(rooms.Rooms))
	for i, room := range rooms.Rooms {
		roomDto[i] = fiber.Map{
			"name":             room.Name,
			"num_participants": room.NumParticipants,
		}
	}

	return c.JSON(roomDto)
}
