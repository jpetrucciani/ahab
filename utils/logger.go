package utils

import (
	"fmt"
	"log"

	"hash/fnv"

	"github.com/fatih/color"
)

// 91,92,93,94,95,96,97
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
	log.SetOutput(new(baseLogWriter))
}

type baseLogWriter struct {
}

func (baseLogWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
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
	log.Print(w.format(b))
	return len(b), nil
}

func (w *LogWriter) format(b []byte) string {
	return fmt.Sprintf("%s: %s", w.color(w.name), string(b))
}
