package main

import (
	"errors"
	"sort"
	"sync"
	"time"
)

type RoomCalendar struct {
	roomId    string
	intervals []*ReservationInterval
	mu        sync.Mutex
}


type ReservationInterval struct {
	start time.Time 
	end time.Time
	reservationId string 
	status ReservationStatus
}

func NewRoomCalendar(roomID string) *RoomCalendar {
	return &RoomCalendar{
		roomId: roomID,
		intervals: make([]*ReservationInterval, 0),
	}
}

// findInsertIndex returns first interval index whose start >= given start
func(rc *RoomCalendar) findInsertIndex(start time.Time) int {
	return sort.Search(len(rc.intervals), func(i int) bool {
		return !rc.intervals[i].start.Before(start)
	})
}

// IsAvailable checks if [start, end) overlaps any existing interval.
// if it's possible to book this room for given time. 
func (rc *RoomCalendar) IsAvailable(start, end time.Time) bool {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	if !start.Before(end) {
		return false 
	}

	i := rc.findInsertIndex(start)
	if i > 0 && !rc.intervals[i -1].end.Before(start) {
		return false 
	}
	if i < len(rc.intervals) && !rc.intervals[i].start.After(end) {
		return false 
	}
	return true 
}


// Reserve inserts a new HOLD interval for reservationId at [start, end).
func (rc *RoomCalendar) Reserve(reservationId string, start, end time.Time ) error {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	if !start.Before(end) {
		return errors.New("invalid interval")
	}

	// check if reservationId already present in the calendar
	for _, inv := range rc.intervals {
		if inv.reservationId == reservationId {
			return errors.New("reservationId already present in calendar")
		}
	}

	// get the insert index
	i := rc.findInsertIndex(start)
	// check if it overlaps with previous interval
	if i > 0 && !rc.intervals[i -1].end.Before(start) {
		return errors.New("overlap with the previous interval") 
	}
	// check if it overlaps with next interval
	if i < len(rc.intervals) && !rc.intervals[i].start.After(end) {
		return errors.New("overlap with the next interval") 
	}
	// create a new reservation interval 
	newIntv := &ReservationInterval{
		start: start, 
		end: end,
		reservationId: reservationId,
		status: ReservationStatus_OnHold,
	}
	// increase the size of intervals by 1 
	rc.intervals = append(rc.intervals, &ReservationInterval{})
	// move all intervals from i index by 1 
	copy(rc.intervals[i + 1:], rc.intervals[i:])
	// set ith interval to be new interval 
	rc.intervals[i] = newIntv

	return nil 
}

// Confirm marks interval with reservationID as CONFIRMED.
func(rc *RoomCalendar) Confirm(reservationId string) error {
	// get the interval corresponding to this reservationID
	rc.mu.Lock()
	defer rc.mu.Unlock()
	for _, intv  := range rc.intervals {
		if intv.reservationId == reservationId {
			intv.status = ReservationStatus_Confirmed
			return nil 
		}
	}
	return errors.New("reservation interval not found")
}

// Release removes interval for reservationID (used for cancel).
func(rc *RoomCalendar) Release(reservationId string) error {
	// get the interval corresponding to this reservationID
	rc.mu.Lock()
	defer rc.mu.Unlock()
	for i, intv  := range rc.intervals {
		if intv.reservationId == reservationId {
			rc.intervals = append(rc.intervals[:i], rc.intervals[i + 1:]...)
			return nil; 
		}
	}
	return errors.New("reservation interval not found")
}