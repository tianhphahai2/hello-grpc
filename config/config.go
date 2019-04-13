package config

type Config struct {
	 Datatable struct{
	 	Host string `json:"host"`
	 	Post int `json:"port"`
	 	Name string `json:"name"`
	 	User string `json:"user"`
	 	Password string `json:"password"`
	 } `json:"database"`
}
