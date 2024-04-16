package global

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type CsvRecord struct {
	file_path string
	file      *os.File
	writer    *csv.Writer
}

func GenCsvRecord(file_path string) *CsvRecord {
	cr := &CsvRecord{}
	cr.file_path = file_path
	var err error
	cr.file, err = os.Create(file_path)
	if err != nil {
		log.Fatal(err)
	}
	cr.writer = csv.NewWriter(cr.file)

	return cr
}

func TargetCsvRecord(file_path string) *CsvRecord {
	cr := &CsvRecord{}
	cr.file_path = file_path
	cr.writer = csv.NewWriter(cr.file)

	return cr
}

func ReadCsvData(fileName string) [][]string {

	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	ans := make([][]string, 0)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		ans = append(ans, row)
	}

	return ans
}

func (C *CsvRecord) WriteLine(info []string) {
	C.writer.Write(info)
	C.writer.Flush()
}

// 在现有文件中追加数据
func (C *CsvRecord) UpdateLine(info []string) {

	//OpenFile读取文件，不存在时则创建，使用追加模式
	File, err := os.OpenFile(C.file_path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("文件打开失败！")
	}
	defer File.Close()

	//创建写入接口
	WriterCsv := csv.NewWriter(File)

	//写入一条数据，传入数据为切片(追加模式)
	err1 := WriterCsv.Write(info)
	if err1 != nil {
		panic("WriterCsv写入文件失败")
	}
	WriterCsv.Flush() //刷新，不刷新是无法写入的
}

func (C *CsvRecord) Close() {
	C.file.Close()
}

func ReadCsv(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	reader := csv.NewReader(f)

	// 可以一次性读完
	result, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return result
}
