package constants

import (
	"os"
	"runtime"
)

var (
	LOG_PATH     = "devmind.log"
	CONFIG_PATHS = []string{"."}
	SETUP_PATH   = ""
)

func Init() {
	if runtime.GOOS == "windows" {
		LOG_PATH = os.Getenv("APPDATA") + "\\devmind.log"
		CONFIG_PATHS = append(CONFIG_PATHS,
			os.Getenv("APPDATA")+"\\devmind",
			os.Getenv("LOCALAPPDATA")+"\\devmind",
			os.Getenv("PROGRAMDATA")+"\\devmind")

		SETUP_PATH = os.Getenv("APPDATA") + "\\devmind"

	} else if runtime.GOOS == "linux" {
		LOG_PATH = "/var/logs/devmind.log"
		CONFIG_PATHS = append(CONFIG_PATHS,
			os.Getenv("HOME")+"/.devmind",
			"/etc/devmind",
			os.Getenv("XDG_CONFIG_HOME")+"/devmind")
		SETUP_PATH = os.Getenv("HOME") + "/.devmind"

	} else if runtime.GOOS == "darwin" {
		LOG_PATH = "/Library/Logs/devmind.log"
		CONFIG_PATHS = append(CONFIG_PATHS,
			"$HOME/Library/Application Support/devmind")
		SETUP_PATH = os.Getenv("HOME") + "/Library/Application Support/devmind"
	}
}
