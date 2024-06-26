package entity

var ExtBlackList = map[string]struct{}{
	".download": {},
	".part":     {},
}

var DirsBlackList = map[string]struct{}{
	"#recycle": {},
	"@eaDir":   {},
}
