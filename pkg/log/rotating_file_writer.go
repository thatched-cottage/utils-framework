package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// RotatingFileWriter 管理日志滚动的结构体
type RotatingFileWriter struct {
	mu      sync.Mutex
	out     io.Writer
	path    string
	name    string
	current *os.File
}

// NewRotatingFileWriter 创建一个新的 RotatingFileWriter 实例
func NewRotatingFileWriter(filePath, fileName string) *RotatingFileWriter {
	rfw := &RotatingFileWriter{
		path: filePath,
		name: fileName,
	}
	rfw.rotateFile()
	go rfw.scheduleRotation()
	return rfw
}

// Write 实现 io.Writer 接口，写入数据
func (r *RotatingFileWriter) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.current.Write(p)
}

// rotateFile 轮换当前日志文件
func (r *RotatingFileWriter) rotateFile() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.current != nil {
		r.current.Close()
	}
	now := time.Now()
	filename := fmt.Sprintf("%s-%d-%02d-%02d-%02d.log", r.name, now.Year(), now.Month(), now.Day(), now.Hour())
	file, err := os.OpenFile(filepath.Join(r.path, filename), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	r.current = file
}

// scheduleRotation 设置每小时轮换一次
func (r *RotatingFileWriter) scheduleRotation() {
	for {
		now := time.Now()
		// 等待到下一个整点
		next := now.Truncate(time.Hour).Add(time.Hour)
		time.Sleep(time.Until(next))
		r.rotateFile()
	}
}
