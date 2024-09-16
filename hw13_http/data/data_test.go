package data

type T2 struct {
	Users []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"users"`
}
