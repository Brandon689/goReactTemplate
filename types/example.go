package types

type Profile struct {
	ID     int64 `bun:",pk,autoincrement"`
	Lang   string
	Active bool
	UserID int64
}

// User has many profiles.
type User struct {
	ID       int64 `bun:",pk,autoincrement"`
	Name     string
	Profiles []*Profile `bun:"rel:has-many,join:id=user_id"`
}
