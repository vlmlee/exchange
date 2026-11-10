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
