package main

import "go.uber.org/zap"
import "time"

func main() {
	logger, _ := zap.NewDevelopment()
	logger.Info("Hello zap", zap.String("key", "value"), zap.Time("now", time.Now()))
  
	logger.Debug("msg", zap.String("key", "String"), zap.Ints("ints", []int{10, 20}))
 
}
