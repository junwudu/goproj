package auth


import (
	"regexp"
)

type AuthProcessor struct {
	Provider regexp.Regexp
	Do func(*Auth, *SignParameter) string
	Do2Url func(*Auth, *SignParameter) string
}


var AuthProcessors = []AuthProcessor {

}
