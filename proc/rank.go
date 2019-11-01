package proc

import (
	frw "aladinfun.com/ZonePet/ZonePetServer/framework"
	"aladinfun.com/ZonePet/ZonePetServer/proto/pb"
	"aladinfun.com/ZonePet/ZonePetServer/srv/ranksrv/rank"
	"aladinfun.com/oss/logger"
	"time"
)

type RankInit struct {
	frw.ProcBase
}

func (p *RankInit) New() frw.IProcBase {
	return &RankInit{}
}

func (p *RankInit) Handle() (rspProto frw.IRspProto, err error) {
	req := &pb.RankInitReq{}
	if err = p.ReceiveRequest(req); err != nil {
		return
	}
	rsp := &pb.RankInitRsp{
		Head: &pb.RspHead{},
	}
	rspProto = rsp
	rankMgr := rank.GetMgrInstance()
	err = rankMgr.RankInit(req.Info)
	if err != nil {
		p.Log(logger.WARN, "Init rank failed", err.Error())
		return
	}
	return
}

type DataUpdate struct {
	frw.ProcBase
}

func (p *DataUpdate) New() frw.IProcBase {
	return &DataUpdate{}
}

func (p *DataUpdate) Handle() (rspProto frw.IRspProto, err error) {
	req := &pb.RankDataUpdateReq{}
	if err = p.ReceiveRequest(req); err != nil {
		return
	}
	rsp := &pb.RankDataUpdateRsp{
		Head: &pb.RspHead{},
	}
	rspProto = rsp
	rankMgr := rank.GetMgrInstance()
	data := &pb.RankUnitData{
		UniqueID:   req.UniqueID,
		Score:      req.Score,
		UpdateTime: uint32(time.Now().Unix()),
	}
	err = rankMgr.RankUpdate(req.Name, data)
	if err != nil {
		p.Log(logger.WARN, "Update rank failed", err.Error())
		return
	}
	return
}

type RankClose struct {
	frw.ProcBase
}

func (p *RankClose) New() frw.IProcBase {
	return &RankClose{}
}

func (p *RankClose) Handle() (rspProto frw.IRspProto, err error) {
	req := &pb.RankCloseReq{}
	if err = p.ReceiveRequest(req); err != nil {
		return
	}
	rsp := &pb.RankCloseRsp{
		Head: &pb.RspHead{},
	}
	rspProto = rsp

	rankMgr := rank.GetMgrInstance()
	err = rankMgr.RankClose(req.Name)
	if err != nil {
		p.Log(logger.WARN, "Close rank failed", err.Error())
		return
	}
	return
}

type DataDelete struct {
	frw.ProcBase
}

func (p *DataDelete) New() frw.IProcBase {
	return &DataDelete{}
}

func (p *DataDelete) Handle() (rspProto frw.IRspProto, err error) {
	req := &pb.RankDataDeleteReq{}
	if err = p.ReceiveRequest(req); err != nil {
		return
	}
	rsp := &pb.RankDataDeleteRsp{
		Head: &pb.RspHead{},
	}
	rspProto = rsp
	rankMgr := rank.GetMgrInstance()
	if err = rankMgr.RankDelete(req.Name, req.UniqueID); err != nil {
		p.Log(logger.WARN, "Delete rank data failed", err.Error())
		return
	}
	return
}

type QueryRank struct {
	frw.ProcBase
}

func (p *QueryRank) New() frw.IProcBase {
	return &QueryRank{}
}

func (p *QueryRank) Handle() (rspProto frw.IRspProto, err error) {
	req := &pb.RankQueryReq{}
	if err = p.ReceiveRequest(req); err != nil {
		return
	}
	rsp := &pb.RankQueryRsp{
		Head: &pb.RspHead{},
	}
	rspProto = rsp

	rankMgr := rank.GetMgrInstance()
	rankInfo, err := rankMgr.RankQuery(req.Name, req.UniqueID)
	if err != nil {
		p.Log(logger.WARN, "Query rank data failed", err.Error())
		return
	}
	rsp.RankInfo = rankInfo
	return
}

type QueryTop struct {
	frw.ProcBase
}

func (p *QueryTop) New() frw.IProcBase {
	return &QueryTop{}
}

func (p *QueryTop) Handle() (rspProto frw.IRspProto, err error) {
	req := &pb.RankTopQueryReq{}
	if err = p.ReceiveRequest(req); err != nil {
		return
	}
	rsp := &pb.RankTopQueryRsp{
		Head: &pb.RspHead{},
	}
	rspProto = rsp
	rankMgr := rank.GetMgrInstance()
	topList, err := rankMgr.RankQueryTop(req.Name, req.ReqNum)
	if err != nil {
		p.Log(logger.WARN, "Query rank top failed", err.Error())
		return
	}
	rsp.TopRank = topList
	return
}

type QueryByScore struct {
	frw.ProcBase
}

func (p *QueryByScore) New() frw.IProcBase {
	return &QueryByScore{}
}

func (p *QueryByScore) Handle() (rspProto frw.IRspProto, err error) {
	req := &pb.RankQueryByScoreReq{}
	if err = p.ReceiveRequest(req); err != nil {
		return
	}
	rsp := &pb.RankQueryByScoreRsp{
		Head: &pb.RspHead{},
	}
	rspProto = rsp
	rankMgr := rank.GetMgrInstance()
	rank, err := rankMgr.RankQueryByScore(req.Name, req.Score)
	if err != nil {
		p.Log(logger.WARN, "Query rank by score failed", err.Error())
		return
	}
	rsp.Ranking = rank
	return
}
