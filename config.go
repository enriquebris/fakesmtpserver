package main

type Config struct {
	Address           string `json:"address"`
	Domain            string `json:"domain"`
	AllowInsecureAuth bool   `json:"allowInsecureAuth"`
	Output            string `json:"output"`
}
