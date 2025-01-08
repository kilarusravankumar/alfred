package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
    "os"
)


var logger *zap.Logger


func InitLogger() *zap.Logger {
    logger , _ =  zap.NewProduction()
    defer logger.Sync()

    atmoicLevel := zap.NewAtomicLevel()
    atmoicLevel.SetLevel(zapcore.InfoLevel)
    
    logger =  zap.New(zapcore.NewCore(
        zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
        zapcore.Lock(os.Stdout),
        atmoicLevel,

    ))

    return logger
}
