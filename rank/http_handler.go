package rank

import (
	"encoding/json"
	"net/http"
	"ranksrv/proto/pb"
	"strconv"
)

func httpListen(httpPort string) {
	addr := ":" + httpPort
	http.HandleFunc("/rank/help", rankHelp)
	http.HandleFunc("/rank/info", rankInfo)
	http.HandleFunc("/rank/init", rankInit)
	http.HandleFunc("/rank/close", rankClose)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func rankHelp(w http.ResponseWriter, r *http.Request) {
	helpContext := `	***排行榜模块管理***
	1.排行榜信息接口url:http://ip:port/rank/info
	2.排行榜初始化接口url:http://ip:port/rank/init
		请求格式：
			name=star_rank&type=1&top_num=100&flush_interval=30&score_min_limit=25&score_num=1
		参数说明：
			name：为初始化排行榜名字，每个排行榜唯一，后续对该排行榜进行操作的依据
			type：为排行榜类型，1表示限量排行榜，2表示全量排行榜
			top_num：为设置限量排行榜大小，全量排行榜无意义
			flush_interval：全量排行榜为排行榜刷新频率，限量排行榜为数据落地频率
			score_min_limit：限制进入排行榜的分数的第一个维度下限
			score_num：分数的维度数量
	3.排行榜关闭接口url:http://ip:port/rank/close
		请求格式：
			name=star_rank
		参数说明: 
			name为要删除关闭排行榜名字`
	w.Write([]byte(helpContext + "\n"))
}

func rankInfo(w http.ResponseWriter, r *http.Request) {
	rankInfo := mgr.RankInfo()
	buf, err := json.Marshal(rankInfo)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(buf)
}

func rankInit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := "name"
	name := r.Form.Get(key)
	key = "type"
	rankType, err := strconv.Atoi(r.Form.Get(key))
	if err != nil {
		w.Write([]byte(key + err.Error()))
		return
	}
	key = "top_num"
	topNum, err := strconv.Atoi(r.Form.Get(key))
	if err != nil {
		w.Write([]byte(key + err.Error()))
		return
	}
	key = "flush_interval"
	flushInterval, err := strconv.Atoi(r.Form.Get(key))
	if err != nil {
		w.Write([]byte(key + err.Error()))
		return
	}
	key = "score_min_limit"
	scoreMinLimit, err := strconv.Atoi(r.Form.Get(key))
	if err != nil {
		w.Write([]byte(key + err.Error()))
		return
	}
	key = "score_num"
	scoreNum, err := strconv.Atoi(r.Form.Get(key))
	if err != nil {
		w.Write([]byte(key + err.Error()))
		return
	}
	rankInfo := &pb.RankInfo{
		Name:          name,
		Type:          pb.RANK_TYPE(rankType),
		TopNum:        uint32(topNum),
		FlushInterval: uint32(flushInterval),
		ScoreMinLimit: uint64(scoreMinLimit),
		ScoreNum:      uint32(scoreNum),
	}
	if err = mgr.RankInit(rankInfo); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("init success\n"))
	return
}

func rankClose(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	if err := mgr.RankClose(name); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("close success\n"))
}
