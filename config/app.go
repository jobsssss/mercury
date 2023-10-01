package config

import "mercury/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// Application name
			"name": config.Env("APP_NAME", "mercury"),

			// Enviroment, support: local, stage, production, test
			"env": config.Env("APP_ENV", "production"),

			// Whether to debugging mode
			"debug": config.Env("APP_DEBUG", false),

			// The http server port
			"port": config.Env("APP_PORT", "3000"),

			// Seceret for session,JWT
			"key": config.Env("APP_KEY", "44556a9bcf9ea160a0a6532b277da32f304af0ed"),

			// Return web applition url
			"url": config.Env("APP_URL", "http://localhost:3000"),

			// Timezonf,application context with JWT, logs
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
