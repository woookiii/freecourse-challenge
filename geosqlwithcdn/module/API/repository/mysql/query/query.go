package query

const (
	InsertIgnoreUser = "" +
		"INSERT IGNORE user (username, description, hobby) VALUES (?, ?, ?);"
	//if IGNORE, insert run when there is no existing row for unique column, kinda IF NOT exists

	InsertIgnoreUserLocation = "" +
		"INSERT IGNORE user_location (username, latitude, hardness, location) VALUES (?, ?, ?, POINT(?, ?));"
)
