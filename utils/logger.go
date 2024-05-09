package utils

import (
	"fmt"
	"log"
	"time"

	"hash/fnv"

	"github.com/fatih/color"
)

var COLORS = []color.Attribute{
	color.FgHiRed,
	color.FgHiGreen,
	color.FgHiYellow,
	color.FgHiBlue,
	color.FgHiMagenta,
	color.FgHiCyan,
	color.FgHiWhite,
}

func init() {
	log.SetFlags(0)
	log.SetOutput(new(timestampedLogWriter))
}

type timestampedLogWriter struct {
}

func (timestampedLogWriter) Write(bytes []byte) (int, error) {
	return fmt.Printf("%s %s", formatTime(time.Now().UTC()), string(bytes))
}

func formatTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05.000")
}

type LogWriter struct {
	log.Logger
	name  string
	color func(a ...interface{}) string
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func NewLogWriter(name string) *LogWriter {
	hashInt := hash(name)
	return &LogWriter{
		log.Logger{},
		name,
		color.New(COLORS[int(hashInt)%len(COLORS)]).SprintFunc(),
	}
}

func (w *LogWriter) Write(b []byte) (int, error) {
	log.Printf(w.format(b))
	return len(b), nil
}

func (w *LogWriter) format(b []byte) string {
	return fmt.Sprintf("%s: %s", w.color(w.name), string(b))
}
