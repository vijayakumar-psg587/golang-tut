package models

var (
	addresses []*Address
)

type Address struct {
	aptName string
	street  string
	city    string
	state   string
}
