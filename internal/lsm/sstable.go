package lsm

import (
	"encoding/csv"
	"os"
	"sort"
)

type SSTable struct {
	data map[string]string
}

func (s *SSTable) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		writer.Write([]string{k, s.data[k]})
	}

	return nil
}

func CompactTables(files []string, outputFile string) error {
	mergedData := make(map[string]string)

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()

		if err != nil {
			file.Close()
			return err
		}

		file.Close()

		for _, record := range records {
			if len(record) == 2 {
				key := record[0]
				value := record[1]
				mergedData[key] = value
			}
		}
	}

	sstable := SSTable{data: mergedData}
	return sstable.SaveToFile(outputFile)
}
