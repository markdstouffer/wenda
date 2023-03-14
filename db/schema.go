package db

import "time"

type TaskStatus int8

const (
	ToDo TaskStatus = iota
	Completed
	Archived
)

type User struct {
	UID         string `json:"uid"`
	Token       string `json:"token"`
	DiscordID   string `json:"discordID"`
	DiscordName string `json:"discordName"`
}

type Task struct {
	ID           int8      `json:"id"`
	DiscordID    string    `json:"discordID"`
	TimeCreated  time.Time `json:"timeCreated"`
	LastModified time.Time `json:"lastModified"`
	Content      string    `json:"content"`
	Status       int8      `json:"status"`
	TaskDate     time.Time `json:"taskDate"`
}
