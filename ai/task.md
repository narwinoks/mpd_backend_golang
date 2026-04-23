role
senior golang development

context
make logging middleware to handler error 

make with logrus

tambahkan package
gopkg.in/natefinch/lumberjack.v2 untuk memecah log
setting an

Filename:   "./storage/mpd.log",
MaxSize:    50, 
MaxBackups: 30, 
MaxAge:     28, 
Compress:   true, 


i hope format error
{
"level": "error",
"time": "2026-04-23T23:17:28Z",
"message": "User failed to authenticate",
"username": "winarno",
"ip_address": "192.168.1.5",
"module": "auth",
"action": "login"
}