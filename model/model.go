package model

type Userprofile struct {
	ID          uint32 `json:"id"`
	Username    string `json:"username"`
	Dob         string `json:"dob"`
	Age         int32  `json:"age"`
	Email       string `json:"email"`
	Phonenumber uint32 `json:"phonenumber"`
}
