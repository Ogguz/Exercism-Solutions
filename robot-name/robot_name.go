package robotname

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// Robot
type Robot struct {
	name string
}

var usedNames []string

// Name creates an unique name with 2 upper characters followed by 3 numbers
func (r *Robot) Name() (string,error) {

	if r.name == "" {
		candidate := stringValueOf(getRandomNumber(26)) + stringValueOf(getRandomNumber(26)) + strconv.Itoa(getRandomNumber(9)) + strconv.Itoa(getRandomNumber(9)) + strconv.Itoa(getRandomNumber(9))
		for _,n := range usedNames{
			if candidate == n {
				return "", errors.New("duplicate robot name")
			}
		}
		usedNames = append(usedNames, candidate)
		r.name = candidate
	}

	return r.name,nil
}

// Reset Wipes robot name
func (r *Robot) Reset()  {
	r.name = ""
}

//
func stringValueOf(i int) string {
	var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return string(alphabet[i])
}

//
func getRandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
