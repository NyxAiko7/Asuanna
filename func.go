package main

import (
	"encoding/json"
	"strings"
	"path/filepath"
	"os"
	"os/exec"
)

type User struct {
	Path string `json:"Path"`
	Wallpaperdaemon string `json:"wallpaper-daemon"`
	Wal  string `json:"Wal"`
	Pkill string `json:"pkill"`
	Astart string `json"astart"`
	Swaync string `json"swaync"`
	WalTelegram string `json"WalTelegram"`
}

func loadConfig(customPath string) User {
	path := customPath
	if path == "" {
		home := os.Getenv("HOME")
		path = filepath.Join(home, ".config", "asuanna", "config.json")
	}

	file, err := os.Open(path)
	if err != nil {
		return User{}
	}
	defer file.Close()

	var user User
	decoder := json.NewDecoder(file)
	_ = decoder.Decode(&user)

	return user
}

func restartags(user User){
	cmd := exec.Command("sh", "-c", "pkill waybar && waybar")
	defer cmd.Start()
}

func getpywal(user User){
	if user.Wal == "yes" || user.Wal == "Yes"{
		cmd := exec.Command("wal", "-i", user.Path)
		defer cmd.Start()
	}
}

func getpywalTelegram(user User){
	cmd := exec.Command("sh", "-c", user.WalTelegram)
	defer cmd.Start()
}

func installwallppepar(user User) error{
	args := append(strings.Fields(user.Wallpaperdaemon), user.Path)
	return exec.Command(args[0], args[1:]...).Start()
}

func restarswaync(user User){
	cmd := exec.Command("sh", "-c", user.Swaync)
	defer cmd.Start()
}
