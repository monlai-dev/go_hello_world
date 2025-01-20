package services

import "github.com/robfig/cron/v3"

type CronJobService struct {
	Cron *cron.Cron
}

func NewCronJobService() *CronJobService {
	return &CronJobService{
		Cron: cron.New(),
	}
}

func (c *CronJobService) StartCronJob() {
	c.Cron.Start()
}

func (c *CronJobService) StopCronJob() {
	c.Cron.Stop()
}

func (c *CronJobService) AddJob(spec string, job cron.Job) (cron.EntryID, error) {
	return c.Cron.AddJob(spec, job)
}

func (c *CronJobService) RemoveJob(id cron.EntryID) {
	c.Cron.Remove(id)
}

func (c *CronJobService) GetJobs() []cron.Entry {
	return c.Cron.Entries()
}

func (c *CronJobService) RunJob(id cron.EntryID) {
	c.Cron.Entry(id).Job.Run()
}

func (c *CronJobService) AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	return c.Cron.AddFunc(spec, cmd)
}
