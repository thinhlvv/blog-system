package model

import "time"

// Date is a custom type.
type Date struct {
	time.Time
}

// ToString methods print datetime with customized format.
func (d *Date) ToString() string {
	return d.Time.String()
}
