package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/widnyana/nvltr/conf"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

// LogReq is http request log
type LogReq struct {
	URI         string `json:"uri"`
	Method      string `json:"method"`
	IP          string `json:"ip"`
	ContentType string `json:"content_type"`
	Agent       string `json:"agent"`
}

// InitLog use for initial log module
func initLog() error {

	var err error

	// init logger
	LogAccess = logrus.New()
	LogError = logrus.New()

	LogAccess.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
		ForceColors:     true,
		FullTimestamp:   true,
	}

	LogError.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
		ForceColors:     true,
		FullTimestamp:   true,
	}

	// set logger
	if err := SetLogLevel(LogAccess, conf.Config.Log.AccessLevel); err != nil {
		return errors.New("Set access log level error: " + err.Error())
	}

	if err := SetLogLevel(LogError, conf.Config.Log.ErrorLevel); err != nil {
		return errors.New("Set error log level error: " + err.Error())
	}

	if err = SetLogOut(LogAccess, conf.Config.Log.AccessLog); err != nil {
		return errors.New("Set access log path error: " + err.Error())
	}

	if err = SetLogOut(LogError, conf.Config.Log.ErrorLog); err != nil {
		return errors.New("Set error log path error: " + err.Error())
	}

	return nil
}

// SetLogOut provide log stdout and stderr output
func SetLogOut(log *logrus.Logger, outString string) error {
	switch outString {
	case "stdout":
		log.Out = os.Stdout
	case "stderr":
		log.Out = os.Stderr
	default:
		f, err := os.OpenFile(outString, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

		if err != nil {
			return err
		}

		log.Out = f
	}

	return nil
}

// SetLogLevel is define log level what you want
// log level: panic, fatal, error, warn, info and debug
func SetLogLevel(log *logrus.Logger, levelString string) error {
	level, err := logrus.ParseLevel(levelString)

	if err != nil {
		return err
	}

	log.Level = level

	return nil
}

// LogRequest record http request
func LogRequest(uri string, method string, ip string, contentType string, agent string) {
	var output string
	log := &LogReq{
		URI:         uri,
		Method:      method,
		IP:          ip,
		ContentType: contentType,
		Agent:       agent,
	}

	if conf.Config.Log.Format == "json" {
		logJSON, _ := json.Marshal(log)

		output = string(logJSON)
	} else {
		// format is string
		output = fmt.Sprintf("|%s header %s| %s %s %s %s %s",
			magenta, reset,
			log.Method,
			log.URI,
			log.IP,
			log.ContentType,
			log.Agent,
		)
	}

	LogAccess.Info(output)
}

func hideToken(token string, markLen int) string {
	if len(token) == 0 {
		return ""
	}

	if len(token) < markLen*2 {
		return strings.Repeat("*", len(token))
	}

	start := token[len(token)-markLen:]
	end := token[0:markLen]

	result := strings.Replace(token, start, strings.Repeat("*", markLen), -1)
	result = strings.Replace(result, end, strings.Repeat("*", markLen), -1)

	return result
}

// LogMiddleware provide gin router handler.
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		LogRequest(c.Request.URL.Path, c.Request.Method, c.ClientIP(), c.ContentType(), c.Request.Header.Get("User-Agent"))
		c.Next()
	}
}
