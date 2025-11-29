package model

import "time"

type Task struct {
    ID        int       `json:"id"`      
    Task      string    `json:"task"`       
    Status    bool      `json:"status"`       
    CreatedAt time.Time `json:"created_at"` 
}
