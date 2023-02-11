package static

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"` //binding валидирует наличие полей в теле запроса (это gin)
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
