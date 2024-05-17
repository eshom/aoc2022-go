package assert

import "log"

func NoError(e error) {
    if e != nil {
        panic(e)
    }
}

func Assert(exp bool, format string, v ...any) {
    if !exp {
        log.Panicf(format, v...)
    }
}
