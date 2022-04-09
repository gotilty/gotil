package gotil

import (
	"github.com/gotilty/gotil/internal"
)

func Shuffle(a interface{}) (interface{}, error) {
	return internal.Shuffle(a)
}

func ShuffleSeed(a interface{}, seed int64) (interface{}, error) {
	return internal.ShuffleSeed(a, seed)

}
