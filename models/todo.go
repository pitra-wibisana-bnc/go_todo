package models

import "time"

type Todo struct {
	ID          int64     `xorm:"bigint not null unique 'id' autoincr pk" json:"id"`
	Title       string    `xorm:"varchar(255) not null 'title'" json:"title"`
	Description string    `xorm:"text not null 'description'" json:"description"`
	CreatedAt   time.Time `xorm:"'created_at'" json:"created_at"`
	UpdatedAt   time.Time `xorm:"'updated_at'" json:"-"`
}
