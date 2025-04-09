package cronjobfx

import (
	"go.uber.org/fx"
	"webapp/internal/services"
)

var Module = fx.Provide(provideCronJob)

func provideCronJob() *services.CronJobService {
	return services.NewCronJobService()
}
