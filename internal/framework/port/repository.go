package port

type Repository interface {
	Disconnect()
}

type UserSql interface {
}

type UserNoSql interface {
}
