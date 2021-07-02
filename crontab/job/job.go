package job

func GetSyncJob() SyncFund {
	j := SyncFund{
		Name: "syncFund",
	}

	j.SetCtxId()

	return j
}
