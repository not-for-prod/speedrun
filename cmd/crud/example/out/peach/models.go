package peach_repository

import (
	peach "github.com/not-for-prod/speedrun/cmd/crud/example/in"
)

type dbPeach struct {
	Id    int `db:"id"`
	Size  int `db:"size"`
	Juice int `db:"juice"`
}

func fromEntity(e peach.Peach) dbPeach {
	return dbPeach{
		Id:    e.Id,
		Size:  e.Size,
		Juice: e.Juice,
	}
}

func (m dbPeach) toEntity() peach.Peach {
	return peach.Peach{
		Id:    m.Id,
		Size:  m.Size,
		Juice: m.Juice,
	}
}
