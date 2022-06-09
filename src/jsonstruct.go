package src

type File struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Linux       string `json:"linux"`
	Windows     string `json:"windows"`
	Darwin      string `json:"darwin"`
}

type Group struct {
	Name  string `json:"name"`
	Files File   `json:"files"`
}

type Config struct {
	Files  File  `json:"files"`
	Groups Group `json:"groups"`
}
