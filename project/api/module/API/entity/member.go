package entity

type Member struct {
	Id          string `db:"id" binding:"required"`
	Name        string `db:"name" binding:"required"`
	Email       string `db:"email" binding:"required"`
	Password    string `db:"password" binding:"required"`
	Role        string `db:"role" binding:"required"`
	CreatedTime string `db:"created_time" binding:"required"`
	UpdatedTime string `db:"updated_time"`
	DeletedTime string `db:"deleted_time"`
}
