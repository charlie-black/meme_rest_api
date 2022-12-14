package models

type Meme struct {
	ID      int    `json:"id"`
	Creator string `json:"creator"`
	Url     string `json:"url"`
}

type DeleteMeme struct {
	ID int `db:"id"`
}

type CreateMeme struct {
	Creator string `json:"creator"`
	Url     string `json:"url"`
}

type GetMemeByID struct {
	ID      int    `db:"id"`
	Creator string `json:"creator"`
	Url     string `json:"url"`
}

type UpdateMeme struct {
	ID      int    `db:"id"`
	Creator string `json:"creator"`
	Url     string `json:"url"`
}

type UserClaims struct {

Username string  `json:"user_name"`
}

type LoginParams struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type SignUpParams struct{
	Email string `json:"email"`
	Password string `json:"password"`
}