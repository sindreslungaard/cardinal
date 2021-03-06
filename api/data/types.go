package data

type Server struct {
	Uid  string `json:"uid"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type Container struct {
	Uid  string `json:"uid"`
	Host string `json:"host"`

	// docker container settings
	Image   string            `json:"image"`
	Restart string            `json:"restart"`
	Env     []string          `json:"env"`
	Ports   map[string]string `json:"ports"`
}

type User struct {
	Uid  string `json:"uid"`
	Hash string `json:"hash"`
}

type Data struct {
	Servers    map[string]Server    `json:"servers"`
	Containers map[string]Container `json:"containers"`
	Users      map[string]User      `json:"users"`
}
