package jdvop

import (
	"net/url"
)

type Request interface {
	Method() string
	Values() url.Values
}
