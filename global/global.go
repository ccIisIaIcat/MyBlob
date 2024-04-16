package global

type IndexStruct struct {
	Index     []string
	Indexhtml []string
}

// 从data中读取全部目录和目录对应html路径，
func GetIndex(index_path string, data_base_path string) [](map[string]string) {
	data := ReadCsvData(index_path)
	ans := make([](map[string]string), 0)
	for _, d := range data {
		temp := make(map[string]string, 0)
		temp["name"] = d[0]
		temp["href"] = data_base_path + d[1]
		ans = append(ans, temp)
	}

	return ans
}
