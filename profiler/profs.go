package profiler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

type executor func(args ...string)

func initProfs()map[string]executor{
	profs := make(map[string]executor)
	profs["reminit"] = func(args ...string) {

		err := Init(args[2])
		if err != nil{
			log.Fatal("Something wrong happened.", err)
		}
		path, err := os.Getwd()
		if err != nil{
			log.Fatal("Something wrong happened.", err)
		}
		log.Println("Successfully inited git repo in ",path)


		err = Add(args[0])
		if err != nil{
			log.Fatal("Something wrong happened.", err)
		}

		err = Commit("init commit")
		if err != nil{
			log.Fatal("Something wrong happened.", err)
		}
			switch args[1] {
			case "":
				err = exec.Command("cmd","/C", "git", "push", "origin").Run()
			default:
				err = exec.Command("cmd","/C", "git", "push", "origin", args[1]).Run()
			}
			if err != nil {
				log.Fatal(err)
			}
			log.Print("Pushed to: ", args[1])
			if args[1] == "" {
				log.Println("default branch")

		}
		log.Println("Successfully committed and pushed changes for", args[0], "in branch", args[1])
	}
	profs["committer"] = func(args ...string) {
		req, err := http.NewRequest("GET", "http://whatthecommit.com/index.txt", nil)
		if err != nil {
			log.Fatal(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		msg, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Your commit message: ", string(msg))
		switch runtime.GOOS {
		case "windows":
			err := exec.Command("cmd", "/C", "git", "add", args[0]).Run()
			if err != nil {
				log.Fatal("Something wrong happened. Check your files list", err)
			}
			log.Println("Added files: ", args[0])
			err = exec.Command("cmd", "/C", "git", "commit", "-m", string(msg)).Run()
			if err != nil {
				log.Fatal(err)
			}
			switch args[1] {
			case "":
				err = exec.Command("cmd","/C", "git", "push", "origin").Run()
			default:
				err = exec.Command("cmd","/C", "git", "push", "origin", args[1]).Run()
			}
			if err != nil {
				log.Fatal(err)
			}
			log.Print("Pushed to: ", args[1])
			if args[1] == "" {
				log.Println("default branch")
			}
		case "linux":
			err := exec.Command("bash", "git", "add", args[0]).Run()
			if err != nil {
				log.Fatal("Something wrong happened. Check your files list", err)
			}
			log.Println("Added files: ", args[0])
			err = exec.Command("bash", "git", "commit", "-m", string(msg)).Run()
			if err != nil {
				log.Fatal(err)
			}
			switch args[1] {
			case "":
				err = exec.Command("bash", "git", "push", "origin").Run()
			default:
				err = exec.Command("bash", "git", "push", "origin", args[1]).Run()
			}
			if err != nil {
				log.Fatal(err)
			}
			log.Print("Pushed to: ", args[1])
			if args[1] == "" {
				log.Println("default branch")
			}
		}
		log.Println("Successfully committed and pushed changes for", args[0], "in branch", args[1])
	}
	return profs
}
