package options

type Types int

const (
	INDEX Types = iota
	UNIQUE
)

var INDEXESVALUE = []string{"INDEX", "UNIQUE INDEX"}
