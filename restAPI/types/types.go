package types

type User struct {
	ID    int64
	Email string
}

type UserCreate struct {
	Email string
}
