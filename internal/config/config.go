package config

import "os"

func NewConfig() (config Config) {

	config = Config{
		GitPath:       os.Getenv("GITHUB_PATH"),
		GitToken:      os.Getenv("GITHUB_TOKEN"),
		GitHookSecret: os.Getenv("GITHUB_HOOK_SECRET"),
		GitUsername:   os.Getenv("GITHUB_USERNAME"),
		SideServer:    os.Getenv("SIDE_SERVER"),
	}
	return config
}

type Config struct {
	GitPath string //Local system path to the github repository

	GitToken string //TToken to authenticate against github

	GitHookSecret string //Secret to authenticate github hooks

	GitUsername string

	SideServer string
}
