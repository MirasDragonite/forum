package handlers

import (
	"log"
	"os"
)

func (h *Handler) infoLog(message string) {
	f, err := os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	infoLogPrint := log.New(os.Stdout, "[INFO]: ", log.Ldate|log.Ltime)
	infoLogWrite := log.New(f, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogPrint.Println(message)
	infoLogWrite.Println(message)
}

func (h *Handler) errorLog(message string) {
	f, err := os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	infoLogPrint := log.New(os.Stdout, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogWrite := log.New(f, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogPrint.Println(message)
	infoLogWrite.Println(message)
}
