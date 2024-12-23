package dlog

import (
	"log/slog"
	"os"
	"runtime"
	"time"
)

func init() {
	folder := os.Getenv("LOG_FOLDER")

	if folder == "" {
		switch runtime.GOOS {
		case "windows":
			folder = `C:\Logs`
		default:
			folder = `/var/log`
		}
	}

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			panic("Unable to create log folder: " + err.Error())
		}
	}

	today := time.Now().Format("2006-01-02")
	logFile := folder + "/" + today + "-json.log"

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("Unable to open log file: " + err.Error())
	}

	jsonHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		AddSource: true,
	})

	logger := slog.New(jsonHandler)
	slog.SetDefault(logger)
}
