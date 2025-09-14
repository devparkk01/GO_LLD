package main

type Room struct {
	roomId       string
	hotelId      string
	roomNo       string
	roomType     RoomType
	roomCalendar *RoomCalendar
	status       RoomStatus
	basePrice    float64
	capacity     int
}


func NewRoom(roomId string, roomNo string, hotelId string, roomType RoomType, status RoomStatus, basePrice float64, capacity int) *Room {
	return &Room {
		roomId: roomId,
		hotelId: hotelId,
		roomNo: roomId,
		roomType: roomType,
		status: status, 
		basePrice: basePrice,
		roomCalendar: NewRoomCalendar(roomId),
		capacity: capacity,
	}
}

