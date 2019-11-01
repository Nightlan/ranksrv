package persistence

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

type SequenceCache struct {
	filePath string
	handle   *os.File
}

func NewSequenceCache(filePath string) *SequenceCache {
	return &SequenceCache{
		filePath: filePath,
	}
}

func (f *SequenceCache) Init() (err error) {
	f.handle, err = os.OpenFile(f.filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	return
}

func (f *SequenceCache) Read() (data [][]byte, err error) {
	f.handle.Seek(0, 0)
	fileInfo, err := f.handle.Stat()
	if err != nil {
		return
	}
	buf := make([]byte, fileInfo.Size())
	_, err = f.handle.Read(buf)
	if err != nil {
		return
	}
	for i := 0; i < len(buf); {
		dataLen := int(binary.BigEndian.Uint32(buf[i : i+4]))
		i += 4
		data = append(data, buf[i:i+dataLen])
		i += dataLen
	}
	return
}

func (f *SequenceCache) ReWrite(data [][]byte) (err error) {
	// 移动文件指针
	_, err = f.handle.Seek(0, 0)

	// 重写文件
	fileWrite := bufio.NewWriter(f.handle)
	for _, v := range data {
		dataLen := make([]byte, 4)
		binary.BigEndian.PutUint32(dataLen, uint32(len(v)))
		_, err = fileWrite.Write(dataLen)
		if err != nil {
			return
		}
		_, err = fileWrite.Write(v)
		if err != nil {
			return
		}
	}
	if err = fileWrite.Flush(); err != nil {
		return
	}
	// 根据新文件大小改变文件大小并同步到磁盘
	size, err := f.handle.Seek(0, 1)
	if err != nil {
		return
	}
	if err = f.handle.Truncate(size); err != nil {
		return
	}
	err = f.handle.Sync()
	return
}

func (f *SequenceCache) Close() (err error) {
	err = f.handle.Close()
	if err != nil {
		fmt.Println("Close file err", err)
		return
	}
	return
}
