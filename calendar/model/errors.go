package model

import "errors"

var BusyError = errors.New("time is busy")
var DuplicateIdError = errors.New("event with id already presented")
var NotFoundError = errors.New("event with id not found")
