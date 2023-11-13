package main

import (
	"os"

	"github.com/merliot/dean"
	"github.com/merliot/hub/models/hp2430n"
	"github.com/merliot/hub/models/hub"
	"github.com/merliot/hub/models/ps30m"
)

func main() {
	h := hub.New("hub01", "hub", "hub01").(*hub.Hub)

	h.ParseWifiAuth()

	gitRemote := os.Getenv("GIT_REMOTE")
	gitKey := os.Getenv("GIT_KEY")
	gitAuthor := os.Getenv("GIT_AUTHOR")
	h.SetGit(gitRemote, gitKey, gitAuthor)

	_, h.BackupHub = os.LookupEnv("BACKUP_HUB")

	server := dean.NewServer(h)
	h.SetServer(server)

	server.Addr = ":8000"
	if port, ok := os.LookupEnv("PORT"); ok {
		server.Addr = ":" + port
	}

	if user, ok := os.LookupEnv("USER"); ok {
		if passwd, ok := os.LookupEnv("PASSWD"); ok {
			server.BasicAuth(user, passwd)
		}
	}

	server.RegisterModel("ps30m", ps30m.New)
	server.RegisterModel("hp2430n", hp2430n.New)

	go server.ListenAndServe()
	server.Run()
}

