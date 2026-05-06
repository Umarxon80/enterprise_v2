package dto





type OutputNode struct {
	Id int `json:"id"`
	Title string `json:"title" `
	Role  string `json:"role"`
	PrevNodeID     int `json:"prev_node_id" `
}
