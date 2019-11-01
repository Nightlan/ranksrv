package persistence

import (
	"bufio"
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
	"os"
	"ranksrv/proto/pb"
)

const (
	BlockSize = 64 //格式文件块大小
	DataLen   = 4  //数据长度定义大小
)

var (
	FormatError = errors.New("format error")
	SizeError   = errors.New("struct is too large")
)

type FormatCache struct {
	filePath    string
	handle      *os.File
	offsetIndex map[string]int64
}

func NewFormatCache(filePath string) *FormatCache {
	return &FormatCache{
		filePath:    filePath,
		offsetIndex: make(map[string]int64),
	}
}

func (f *FormatCache) Init() (err error) {
	f.handle, err = os.OpenFile(f.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	return
}

func (f *FormatCache) Read() (unitData []*pb.RankUnitData, err error) {
	// 移动文件指针到最开始
	f.handle.Seek(0, 0)
	fileInfo, err := f.handle.Stat()
	if err != nil {
		return
	}
	if fileInfo.Size()%BlockSize != 0 {
		err = FormatError
		return
	}
	buf := make([]byte, fileInfo.Size())
	_, err = f.handle.Read(buf)
	if err != nil {
		return
	}
	dataLen := fileInfo.Size() / BlockSize
	unitData = make([]*pb.RankUnitData, 0, dataLen)
	for i := 0; i < len(buf); i += BlockSize {
		lenEnd := i + DataLen
		dataLen := int(binary.BigEndian.Uint32(buf[i:lenEnd]))
		dataEnd := lenEnd + dataLen
		data := &pb.RankUnitData{}
		if err = proto.Unmarshal(buf[lenEnd:dataEnd], data); err != nil {
			return
		}
		// 初始化索引
		f.offsetIndex[data.UniqueID] = int64(i)
		unitData = append(unitData, data)
	}
	return
}

// 写数据
func (f *FormatCache) Write(unitData []*pb.RankUnitData) (err error) {
	notExist := make([]*pb.RankUnitData, 0)
	for _, v := range unitData {
		offset, ok := f.offsetIndex[v.UniqueID]
		if ok { //存在直接更新
			if err = f.update(offset, v); err != nil {
				return
			}
		} else { //不存在保存起来一次性写入
			notExist = append(notExist, v)
		}
	}
	// 在文件后追加
	if err = f.append(notExist); err != nil {
		return
	}
	return
}

func (f *FormatCache) append(unitData []*pb.RankUnitData) (err error) {
	// 移动文件指针到文件末尾
	offsetStart, err := f.handle.Seek(0, 2)
	if err != nil {
		return
	}
	// 重写文件
	fileWrite := bufio.NewWriter(f.handle)
	for _, v := range unitData {
		buf, err1 := f.blockEncode(v)
		err = err1
		if err != nil {
			return
		}
		dataLen, err1 := fileWrite.Write(buf)
		err = err1
		if err != nil {
			return
		}
		f.offsetIndex[v.UniqueID] = offsetStart
		offsetStart += int64(dataLen)
	}
	if err = fileWrite.Flush(); err != nil {
		return
	}
	return
}

func (f *FormatCache) update(offset int64, unitData *pb.RankUnitData) (err error) {
	buf, err := f.blockEncode(unitData)
	if err != nil {
		return
	}
	// 写文件
	_, err = f.handle.WriteAt(buf, offset)
	if err != nil {
		return
	}
	return
}

func (f *FormatCache) blockEncode(unitData *pb.RankUnitData) (buf []byte, err error) {
	dataBuf, err := proto.Marshal(unitData)
	if err != nil {
		return
	}
	if BlockSize-DataLen < len(dataBuf) {
		err = SizeError
		return
	}
	// 格式化单条数据
	buf = make([]byte, 0, BlockSize)
	dataLen := make([]byte, DataLen)
	binary.BigEndian.PutUint32(dataLen, uint32(len(dataBuf)))
	buf = append(buf, dataLen...)
	buf = append(buf, dataBuf...)
	//填充切片后续内存
	buf = buf[:BlockSize]
	return
}

func (f *FormatCache) Close() (err error) {
	err = f.handle.Close()
	return
}
