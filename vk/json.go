package vk
import (
	"time"
	"strconv"
)

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	*t = Timestamp(time.Unix(int64(ts), 0))

	return nil
}

type Count struct {
	Count int
}