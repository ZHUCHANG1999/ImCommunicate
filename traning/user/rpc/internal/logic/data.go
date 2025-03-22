package logic

type User struct {
	Id    string
	Name  string
	Phone string
}

var users = map[string]*User{
	"1": {
		Id:    "1",
		Name:  "木兮",
		Phone: "1111",
	},
	"2": {
		Id:    "2",
		Name:  "xx",
		Phone: "123123",
	},
}
