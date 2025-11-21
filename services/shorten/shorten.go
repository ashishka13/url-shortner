package shorten

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
)

func ShortenLink(link string) (shortLink string, err error) {
	letters := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	randomNum := rand.Int63n(128 * 128)

	lettersArr := strings.Split(letters, "")
	lettersReverseArr := lettersArr
	slices.Reverse(lettersReverseArr)

	// pick letter from these indexes
	letterPlaces := rand.Int63n(52 * 2)
	arr := strings.Split(fmt.Sprint(letterPlaces), "")

	newLetter := ""
	for _, val := range arr {
		index, _ := strconv.Atoi(val)
		newLetter = newLetter + lettersArr[index]
	}

	revLetter := ""
	for _, val := range arr {
		index, _ := strconv.Atoi(val)
		revLetter = revLetter + lettersReverseArr[index]
	}

	url := fmt.Sprintf("%s-%d-%s", newLetter, randomNum, revLetter)

	return url, nil
}
