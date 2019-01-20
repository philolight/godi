package trace

import (
	"fmt"
	"runtime"
	"strings"
)

type trace struct {
	m map[string][]string // 여기에 (이벤트 : 이벤트 구독자) 쌍을 저장함
}

var t = &trace{make(map[string][]string)}

func MyName() string {
	return PathFileWithPC(1)
}

func PathFileWithPC(counter int) string {
	pc, file, _, _ := runtime.Caller(counter + 1)
	path := strings.Split(runtime.FuncForPC(pc).Name(), ".")[0]
	path = path[strings.Index(path, "/")+1:]
	path = strings.Replace(path, "/", ".", -1)

	fileSlice := strings.Split(file, "/")
	file = strings.Split(fileSlice[len(fileSlice)-1], ".")[0]

	return path + ":" + file
}

func Trace() {
	// 이벤트 타입
	eventTypePC, _, _, _ := runtime.Caller(1)
	eventTypePath := runtime.FuncForPC(eventTypePC).Name()
	eventTypePath = strings.Replace(eventTypePath, "/", ".", -1)

	listener := PathFileWithPC(2)

	var slice []string
	var ok bool

	if slice, ok = t.m[eventTypePath]; !ok {
		slice = make([]string, 0)
	}

	slice = append(slice, listener)
	t.m[eventTypePath] = slice
}

func Dump() {
	for eventType, v := range t.m {
		fmt.Println(eventType)
		for _, listenerName := range v {
			fmt.Println(" |-", listenerName)
		}

		fmt.Println()
	}
}
