package logger

import (
	"log"
	"os"
)

var AppLogger *log.Logger

func InitLogger() {

	file, _ :=
		os.OpenFile(
			"logs/app.log",
			os.O_APPEND|
				os.O_CREATE|
				os.O_WRONLY,
			0666,
		)

	AppLogger =
		log.New(
			file,
			"APP:",
			log.LstdFlags,
		)
}

func Info(
	msg string,
) {

	AppLogger.Println(
		"[INFO]",
		msg,
	)
}

func Error(
	msg string,
) {

	AppLogger.Println(
		"[ERROR]",
		msg,
	)
}
