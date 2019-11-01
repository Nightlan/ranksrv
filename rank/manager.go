package rank

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"os"
	"ranksrv/proto/pb"
	"ranksrv/rank/core"
	"ranksrv/rank/persistence"
	"sync"
)

const (
	manageFileName = "rank_info"
)

var (
	TypeNotExist     = errors.New("rank type not exist")
	NameExist        = errors.New("rank name exist")
	NameRankNotExist = errors.New("rank not exist")
)

type Manager struct {
	running  bool // 运行标志位,多线程竞争不敏感
	ranks    sync.Map
	pstCache *persistence.SequenceCache // 持久化缓存
	cacheDir string
	wLock    sync.Mutex
}

func NewManager(cacheDir string) *Manager {
	return &Manager{
		running:  false,
		cacheDir: cacheDir,
	}
}

func (m *Manager) Start(httpPort string) (err error) {
	if m.running {
		return
	}
	_, ret := os.Stat(m.cacheDir)
	if ret != nil {
		if err = os.Mkdir(m.cacheDir, os.ModePerm); err != nil {
			return
		}
	}
	cachePath := m.cacheDir + manageFileName
	m.pstCache = persistence.NewSequenceCache(cachePath)
	if err = m.pstCache.Init(); err != nil {
		return
	}
	data, err := m.pstCache.Read()
	if err != nil {
		return
	}
	for _, v := range data {
		rankInfo := &pb.RankInfo{}
		if err = rankInfo.XXX_Unmarshal(v); err != nil {
			return
		}
		if err = m.doRankInit(rankInfo); err != nil {
			return
		}
	}
	go httpListen(httpPort)
	m.running = true
	return
}

func (m *Manager) RankInfo() (rankInfo []*pb.RankInfo) {
	m.ranks.Range(func(key, value interface{}) bool {
		rank := value.(core.IRank)
		rankInfo = append(rankInfo, rank.RankInfo())
		return true
	})
	return
}

func (m *Manager) RankInit(rankInfo *pb.RankInfo) (err error) {
	if err = m.doRankInit(rankInfo); err != nil {
		return
	}
	return m.saveRankInfo()
}

func (m *Manager) doRankInit(rankInfo *pb.RankInfo) (err error) {
	if _, ok := m.ranks.Load(rankInfo.Name); ok {
		err = NameExist
		return
	}
	var rank core.IRank
	switch rankInfo.Type {
	case pb.RANK_TYPE_WHOLE:
		rank = core.NewWholeRank(rankInfo)
	case pb.RANK_TYPE_PART:
		rank = core.NewPartRank(rankInfo)
	default:
		err = TypeNotExist
		return
	}
	if err = rank.Start(m.cacheDir); err != nil {
		return
	}
	m.ranks.Store(rankInfo.Name, rank)
	return
}

func (m *Manager) saveRankInfo() (err error) {
	ranksInfo := make([]*pb.RankInfo, 0)
	m.ranks.Range(func(key, value interface{}) bool {
		rank := value.(core.IRank)
		ranksInfo = append(ranksInfo, rank.RankInfo())
		return true
	})
	data := make([][]byte, 0, len(ranksInfo))
	var buf []byte
	for _, v := range ranksInfo {
		buf, err = proto.Marshal(v)
		if err != nil {
			return
		}
		data = append(data, buf)
	}
	m.wLock.Lock()
	if err = m.pstCache.ReWrite(data); err != nil {
		m.wLock.Unlock()
		return
	}
	m.wLock.Unlock()
	return
}

func (m *Manager) RankUpdate(name string, data *pb.RankUnitData) (err error) {
	value, ok := m.ranks.Load(name)
	if !ok {
		err = NameRankNotExist
		return
	}
	rank := value.(core.IRank)
	rank.Update(data)
	return
}

func (m *Manager) RankDelete(name, uniqueID string) (err error) {
	value, ok := m.ranks.Load(name)
	if !ok {
		err = NameRankNotExist
		return
	}
	rank := value.(core.IRank)
	rank.Delete(uniqueID)
	return
}

func (m *Manager) RankQuery(name, uniqueID string) (rankInfo *pb.UnitRankInfo, err error) {
	value, ok := m.ranks.Load(name)
	if !ok {
		err = NameRankNotExist
		return
	}
	rank := value.(core.IRank)
	rankInfo, err = rank.QueryRank(uniqueID)
	if err != nil {
		return
	}
	return
}

func (m *Manager) RankQueryTop(name string, reqNum uint32) (
	topList []*pb.RankUnitData, err error) {
	value, ok := m.ranks.Load(name)
	if !ok {
		err = NameRankNotExist
		return
	}
	rank := value.(core.IRank)
	return rank.QueryTop(reqNum)
}

func (m *Manager) RankQueryByScore(name string, score []uint64) (rank uint32, err error) {
	value, ok := m.ranks.Load(name)
	if !ok {
		err = NameRankNotExist
		return
	}
	iRank := value.(core.IRank)
	return iRank.QueryByScore(score)
}

func (m *Manager) RankClose(name string) (err error) {
	value, ok := m.ranks.Load(name)
	if !ok {
		err = NameRankNotExist
		return
	}
	m.ranks.Delete(name)
	if err = m.saveRankInfo(); err != nil {
		return
	}
	rank := value.(core.IRank)
	rank.Close()
	return
}

func (m *Manager) Stop() (err error) {
	if !m.running {
		return
	}
	m.running = false
	m.ranks.Range(func(key, value interface{}) bool {
		rank := value.(core.IRank)
		rank.Stop()
		return true
	})
	return
}
