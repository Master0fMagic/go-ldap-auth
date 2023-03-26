package common

type Role int16

const (
	None     Role = -1
	Admin    Role = 1
	Operator Role = 2
)
