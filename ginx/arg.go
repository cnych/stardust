package ginx

import (
    "github.com/gin-gonic/gin"
    "strconv"
)

const RC_Key = "ctx"

func StrArg(gc *gin.Context, k, def string) string {
    s := gc.Query(k)
    if s != "" {
        return s
    }
    s = gc.Param(k)
    if s != "" {
        return s
    }
    return def
}

func Int64Arg(gc *gin.Context, k string, def int64) int64 {
    s := StrArg(gc, k, "")
    if s == "" {
        return def
    }
    v, err := strconv.ParseInt(s, 10, 64)
    if err != nil {
        return def
    }
    return v
}

func IntArg(gc *gin.Context, k string, def int) int {
    s := StrArg(gc, k, "")
    if s == "" {
        return def
    }
    v, err := strconv.Atoi(s)
    if err != nil {
        return def
    }
    return v
}

func BoolArg(gc *gin.Context, k string, def bool) bool {
    s := StrArg(gc, k, "")
    if s == "" {
        return def
    }
    b, err := strconv.ParseBool(s)
    if err != nil {
        return def
    }
    return b
}
