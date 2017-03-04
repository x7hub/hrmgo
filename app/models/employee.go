package models

import (
    "time"
)

type User struct {
    id int
    name string
    regularizeDate  time.Time
    birthday time.Time
}