package model

type Parser interface {
	Parse([]byte) ParseResult
	Serialize() (string, []interface{})
}
