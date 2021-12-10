package todo

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int `json:"id"`
	ListId int `json:"list_id"`
	UserId int `json:"user_id"`
}

type ListsItem struct {
	Id     int `json:"id"`
	ItemId int `json:"item_id"`
	ListId int `json:"list_id"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Done        bool   `json:"done"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
