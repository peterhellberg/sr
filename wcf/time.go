package wcf

import (
	"regexp"
	"strconv"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(in []byte) (err error) {
	r := regexp.MustCompile("[0-9]{10}")

	ts, err := strconv.ParseInt(r.FindString(string(in)), 10, 64)
	if err != nil {
		return err
	}

	t.Time = time.Unix(ts, 0)

	return nil
}
