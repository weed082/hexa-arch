package port

//! Chat
type ChatRepo interface {
	CreateRoom() (int, error)
	SendMessage() error
}

type ChatSql interface {
}

type FileRepo interface {
}

type FileSql interface {
}

//! User
type UserRepo interface {
}

type UserSql interface {
}

type UserNoSql interface {
}
