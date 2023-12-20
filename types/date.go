package types

import "time"

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("\"2006-01-02\"", string(b))
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(d).Format("\"2006-01-02\"")), nil
}
