package response

type OnlineMenu struct {
	Title     string       `json:"title"`
	Key       string       `json:"key"`
	Name      string       `json:"name"`
	Component string       `json:"component"`
	Redirect  string       `json:"redirect"`
	Children  []OnlineMenu `json:"children"`
}
