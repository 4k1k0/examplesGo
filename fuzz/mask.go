package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var (
	validMask = regexp.MustCompile(`^\d{4}[X]{8}\d{4}$`)
)

func Mask(card string) (string, error) {
	lenCard := len(card)
	log.Println(card, lenCard)

	if lenCard != 16 {
		return "", errors.New("len should be 16")
	}

	start := card[0:4]
	between := strings.Repeat("X", 8)
	end := card[12:]
	res := fmt.Sprintf("%s%s%s", start, between, end)

	if !validMask.MatchString(res) {
		return "", errors.New("is not a valid card")
	}

	return res, nil
}
