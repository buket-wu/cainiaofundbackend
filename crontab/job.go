package crontab

import "cainiaofundbackend/crontab/job"

func GetSyncJob() job.SyncFund {
	return job.SyncFund{
		Name: "syncFund",
	}
}
