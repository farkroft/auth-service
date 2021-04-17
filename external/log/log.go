package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
	"gitlab.com/farkroft/auth-service/external/constants"
	"gitlab.com/farkroft/auth-service/external/util"
)

// WIBFormatter struct
type WIBFormatter struct {
	logrus.Formatter
}

// NewLogger format logrus
func NewLogger() {
	logrus.SetLevel(logrus.ErrorLevel)

	switch logrus.GetLevel() {
	case logrus.ErrorLevel:
		logrus.SetFormatter(WIBFormatter{&logrus.JSONFormatter{}})
	case logrus.InfoLevel:
		logrus.SetFormatter(WIBFormatter{&logrus.TextFormatter{FullTimestamp: true}})
	}

	f, err := os.OpenFile(fmt.Sprintf("%s%s", constants.EnvConfigFile, constants.LogFile), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("cannot open file log: %s", err.Error())
	}

	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	logrus.SetOutput(mw)
}

// Format return time formatted with timezone
func (w WIBFormatter) Format(e *logrus.Entry) ([]byte, error) {
	value := util.WIBTimezone(e.Time.UTC())
	e.Time = value
	return w.Formatter.Format(e)
}

// Errorf error with format string
func Errorf(str string, err error) {
	str = str + ": %s"
	_, file, line, _ := runtime.Caller(1) // skip caller to one frame
	message := fmt.Sprintf("%s, %s:%d", err, file, line)
	logrus.Errorf(str, message)
}

// Infoln log info message
func Infoln(str string) {
	logrus.Infoln(str)
}
