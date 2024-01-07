package model 

type User struct {
	Name string 
	Family string 
	ID int
}

type Body struct{
	Name string `json:owner`
	ID int `json : id`
}