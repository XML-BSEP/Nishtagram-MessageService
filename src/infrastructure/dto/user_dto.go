package dto

type UserDto struct {
	Id		 string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Private bool `json:"private"`
	Image	string `json:"image"`
}

type UserIdsDto struct{
	Ids []string `json:"ids"`
}
