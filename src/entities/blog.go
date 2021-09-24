package entities

import "fmt"

type Tag struct {
	Id         int64
	Title      string
	Body       string
	Created_at string
	Updated_at string
}

func (tag Tag) ToString() string {
	return fmt.Sprintf("id: %d\ntitle: %s\nbody: %s\ncreated_at: %s\nupdated_at: %s",
		tag.Id, tag.Title, tag.Body, tag.Created_at, tag.Updated_at)
}
