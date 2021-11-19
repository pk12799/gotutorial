package Models

type TODO struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (b *TODO) Tablename() string {
	return "todo"
}
