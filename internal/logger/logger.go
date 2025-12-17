package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


var Log *zap.Logger;

// Initializes the logger, receives `environment` - basically "production" or "development"
// Only initialise it once for better performance
func Init(environment string) {
	var cfg zap.Config;

	if environment == "production" {
		// builds a production level config
		cfg = zap.NewProductionConfig();
		// timestamp
		cfg.EncoderConfig.TimeKey = "timestamp";
		//
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		// builds a development level config.
		cfg = zap.NewDevelopmentConfig();
	}

	// AddCallerSkip to know the actual caller of the logger.
	logger, err := cfg.Build(zap.AddCaller(), zap.AddCallerSkip(1));

	if err != nil {
		panic(err);
	}

	Log = logger;
}