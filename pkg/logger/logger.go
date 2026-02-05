package logger

import (
    "encoding/json"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// Write 支持多个参数 msg（字符串、结构体、切片、map 等）
// 会把多个参数合并成一条日志记录写入
func Write(module string, msgs ...interface{}) {
    dateStr := time.Now().Format("2006-01-02")

    // 确保日志写到项目根目录
    dir, _ := filepath.Abs(filepath.Join("log", module))
    _ = os.MkdirAll(dir, os.ModePerm)

    logFile := filepath.Join(dir, dateStr+".log")
    log.Println("写日志到:", logFile)

    f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Println("打开日志文件失败:", err)
        return
    }
    defer f.Close()

    logger := log.New(f, "", log.LstdFlags)

    var parts []string
    for _, msg := range msgs {
        switch v := msg.(type) {
        case string:
            parts = append(parts, v)
        default:
            b, err := json.Marshal(v)
            if err != nil {
                parts = append(parts, "日志序列化失败")
            } else {
                parts = append(parts, string(b))
            }
        }
    }

    // 把所有参数合并成一行
    line := strings.Join(parts, " ")
    logger.Println(line)
}
