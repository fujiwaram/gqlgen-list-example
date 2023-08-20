package model

import (
	"errors"
	"fmt"
	"io"
	"time"
)

type Time time.Time

const (
	timeFormat = time.DateTime
)

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t *Time) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf(": %w", errors.New("invalid time value"))
	}

	tm, err := time.Parse(timeFormat, str)
	if err != nil {
		return fmt.Errorf("invalid time format : %w", err)
	}
	*t = Time(tm)

	return nil
}

func (t Time) MarshalGQL(w io.Writer) {
	w.Write([]byte("\"" + t.String() + "\""))
}
