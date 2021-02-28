package beer

import (
	"errors"
	"fmt"
)

const (
	verseGeneral = "%s of beer on the wall, %s of beer.\nTake one down and pass it around, %s of beer on the wall.\n"
	verseLast     = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
	verseNoBeer   = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
)

// Song returns the whole song
func Song() string {
	song,_ := Verses(99,0)
	return song
}

func Verses(from, to int) (string, error) {
	if from < 0 || to < 0 || from > 99 || to > 99  || from < to {
		return "", errors.New("invalid number of beers")
	}

	var song string
	for i := from; i>=to; i-- {
		verse,err := Verse(i)
		if err != nil {
			return "", err
		}
		song += verse + "\n"
	}

	return song, nil
}

func Verse(count int) (string,error)  {

	if count < 0 || count > 99 {
		return "",errors.New("invalid number of beers")
	}

	switch count {
	case 1:
		return verseLast,nil
	case 0:
		return verseNoBeer,nil
	default:
		return fmt.Sprintf(verseGeneral,bottlesUp(count),bottlesUp(count),bottlesUp(count-1)),nil
	}
	
}

func bottlesUp(count int) (string) {
	switch count {
	case 1:
		return "1 bottle"
	case 0:
		return "no more bottles:"
	default:
		return fmt.Sprintf("%d bottles", count)
	}
}