package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

type Marshaler interface {
	json.Marshaler
	json.Unmarshaler
}

const dateLayout = time.DateOnly

var (
	_ Marshaler     = (*Date)(nil)
	_ query.Encoder = (*Date)(nil)
)

type Date time.Time

func Today() Date {
	return Date(time.Now())
}

func NewDate(year, month, day int) Date {
	if year == 0 && month == 0 && day == 0 {
		return Date(time.Time{})
	}

	return Date(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
}

func (d *Date) UnmarshalJSON(data []byte) error {
	data = bytes.ReplaceAll(bytes.Trim(data, `"`), []byte("null"), nil)

	if len(data) == 0 {
		return nil
	}

	t, err := time.Parse(dateLayout, string(data))
	if err != nil {
		return err
	}

	*d = Date(t)

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	s := time.Time(d).Format(dateLayout)

	return json.Marshal(s)
}

func (d Date) EncodeValues(key string, v *url.Values) error {
	s := time.Time(d).Format(dateLayout)

	v.Set(key, s)

	return nil
}

func (d Date) IsZero() bool {
	return d.AsTime().IsZero()
}

func (d Date) AsTime() time.Time {
	return time.Time(d)
}

type Bool bool

var (
	_ Marshaler     = (*Bool)(nil)
	_ query.Encoder = (*Bool)(nil)
)

func (b Bool) MarshalJSON() ([]byte, error) {
	if b == true {
		return json.Marshal("1")
	}

	return json.Marshal("0")
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	var a any

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	switch v := a.(type) {
	case string:
		s := strings.Trim(v, ` "`)
		if s != "1" && s != "0" && s != "true" && s != "false" {
			return fmt.Errorf("unexpected value: %q", s)
		}

		*b = s == "1" || s == "true"
	case int:
		*b = v != 0
	case bool:
		*b = Bool(v)
	}

	return nil
}

func (b Bool) EncodeValues(key string, v *url.Values) error {
	s := "0"

	if b {
		s = "1"
	}

	v.Set(key, s)

	return nil
}

type Int int64

var _ Marshaler = (*Int)(nil)

func (i *Int) UnmarshalJSON(data []byte) error {
	var a any

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	switch v := a.(type) {
	case string:
		s := strings.Trim(v, ` "`)
		if s == "" {
			return nil
		}

		c, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("atoi: %w", err)
		}

		*i = Int(c)
	case int:
		*i = Int(v)
	case float64:
		*i = Int(int(v))
	}

	return nil
}

func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(i))
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	data = bytes.ReplaceAll(bytes.Trim(data, `"`), []byte("null"), nil)

	if len(data) == 0 {
		return nil
	}

	v, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return fmt.Errorf("parseInt: %w", err)
	}

	*t = Timestamp(time.Unix(v, 0))

	return nil
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Unix())
}
