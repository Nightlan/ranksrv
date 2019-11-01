package rank

var mgr *Manager

func InitRankMgr(cacheDir, httpPort string) (err error) {
	mgr = NewManager(cacheDir)
	if err = mgr.Start(httpPort); err != nil {
		return
	}
	return
}

func GetMgrInstance() *Manager {
	return mgr
}

func CloseRankMgr() (err error) {
	return mgr.Stop()
}
