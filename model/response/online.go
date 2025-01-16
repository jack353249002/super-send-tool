package response

type OnlineMenu struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Key       string       `json:"key"`
	Name      string       `json:"name"`
	Component string       `json:"component"`
	Redirect  string       `json:"redirect"`
	Token     string       `json:"token"`
	Children  []OnlineMenu `json:"children"`
}
