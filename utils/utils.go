package utils

func MakeLogKeys() []string {
  logKeys := []string{
    "name", 
    "msg", 
    "args", 
    "levelname", 
    "levelno", 
    "pathname", 
    "filename", 
    "module", 
    "exc_info", 
    "exc_text", 
    "stack_info", 
    "lineno", 
    "funcName", 
    "created", 
    "msecs", 
    "relativeCreated", 
    "thread", 
    "threadName", 
    "processName", 
    "process",
  }
  return logKeys
}
