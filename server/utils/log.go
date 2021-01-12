package utils

import (
  "go-element-admin/configs"
  "fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"time"
)

func NewLogger() *logrus.Logger  {
	logr := logrus.New()
	return logr
}

func DefaultLogger(debug bool) *logrus.Logger {
	lgr := NewLogger()

	writerSto := os.Stdout

	date := time.Now().Format("2006-01-02")
	name := path.Join(configs.ApplicationConfig.LogPath,fmt.Sprintf("%s.%s",date,"log"))
	writerFile, err := os.OpenFile(name,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil {
		log.Fatalf("create file log failed: %v",err)
	}

	if debug {
		lgr.SetLevel(logrus.DebugLevel)
	} else {
		lgr.SetLevel(logrus.InfoLevel)
	}

	lgr.SetOutput(io.MultiWriter(writerSto,writerFile))

	lgr.SetFormatter(&logrus.TextFormatter{})

	return lgr
}
