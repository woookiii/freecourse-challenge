package query

const (
	InsertIgnoreUser = "" +
		"INSERT IGNORE user (username, description, hobby) VALUES (?, ?, ?);"
	//if IGNORE, insert run when there is no existing row for unique column, kinda IF NOT exists

	InsertIgnoreUserLocation = "" +
		"INSERT IGNORE user_location (username, latitude, longitude, location) VALUES (?, ?, ?, POINT(?, ?));"
)

const (
	GetUserByNone = "" +
		"SELECT " +
		"u.username, u.image, u.description, u.hobby, " +
		"ul.latitude, ul.longitude " +
		"FROM user AS u JOIN user_location AS ul " +
		"ON u.username = ul.username WHERE u.username = ?;"

	GetAroundUsers = "" +
		"SELECT " +
		"u.username, u.image, u.description, u.hobby, " +
		"ul.latitude, ul.longitude " +
		"FROM user AS u JOIN user_location AS ul ON u.username = ul.username " +
		"WHERE u.username != ? AND ST_Distance_Sphere(POINT(?, ?), POINT(ul.longitude, ul.latitude)) <= ? " +
		"ORDER BY ST_Distance_Sphere(POINT(?, ?), POINT(ul.longitude, ul.latitude)) LIMIT ?;"
)
