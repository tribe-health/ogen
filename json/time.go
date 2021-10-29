package json

import (
	"time"
)

const (
	dateLayout = "2006-01-02"
	timeLayout = "15:04:05"
)

func ReadDate(i *Reader) (v time.Time, err error) {
	s, err := i.Str()
	if err != nil {
		return v, err
	}
	return time.Parse(dateLayout, s)
}

func WriteDate(s *Writer, v time.Time) {
	s.Str(v.Format(dateLayout))
}

func ReadTime(i *Reader) (v time.Time, err error) {
	s, err := i.Str()
	if err != nil {
		return v, err
	}
	return time.Parse(timeLayout, s)
}

func WriteTime(s *Writer, v time.Time) {
	s.Str(v.Format(timeLayout))
}

func ReadDateTime(i *Reader) (v time.Time, err error) {
	s, err := i.Str()
	if err != nil {
		return v, err
	}
	return time.Parse(time.RFC3339, s)
}

func WriteDateTime(s *Writer, v time.Time) {
	s.Str(v.Format(time.RFC3339))
}

func ReadDuration(i *Reader) (v time.Duration, err error) {
	s, err := i.Str()
	if err != nil {
		return v, err
	}
	return time.ParseDuration(s)
}

func WriteDuration(s *Writer, v time.Duration) {
	s.Str(v.String())
}
