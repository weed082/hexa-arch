package port

//! Chat
type ChatRepo interface {
	CreateRoom() (int, error)
	UploadMsg() error
}

//! User
type UserRepo interface {
}
