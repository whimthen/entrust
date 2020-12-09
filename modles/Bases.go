package modles

import "time"

type BaseEntity struct {
	ID       int64     `sql:"id"`
	CreateAt time.Time `sql:"create_at"`
	UpdateAt time.Time `sql:"update_at"`
}
