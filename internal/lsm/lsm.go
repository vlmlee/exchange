package lsm

import (
	"encoding/csv"
	"fmt"
	"os"
)

type LSMTree struct {
	memtable  *Memtable
	sstables  []string
	threshold int
	fileCount int
}

func NewLSMTree(threshold int) *LSMTree {
	return &LSMTree{
		memtable:  NewMemtable(),
		threshold: threshold,
		fileCount: 0,
	}
}

func (lsm *LSMTree) Set(key, value string) {
	lsm.memtable.Set(key, value)
	if len(lsm.memtable.data) >= lsm.threshold {
		lsm.FlushToSSTable()
	}
}

func (lsm *LSMTree) Get(key string) (string, bool) {
	if val, exists := lsm.memtable.Get(key); exists {
		return val, true
	}

	for i := len(lsm.sstables) - 1; i >= 0; i-- {
		if val, exists := searchSSTable(lsm.sstables[i], key); exists {
			return val, true
		}
	}

	return "", false
}

func (lsm *LSMTree) FlushToSSTable() {
	filename := fmt.Sprintf("sstable_%d.csb", lsm.fileCount)
	sstable := SSTable{data: lsm.memtable.data}

	if err := sstable.SaveToFile(filename); err == nil {
		lsm.sstables = append(lsm.sstables, filename)
		lsm.memtable = NewMemtable()
		lsm.fileCount++
	} else {
		fmt.Println("Error writing SSTable:", err)
	}
}

func searchSSTable(filename, key string) (string, bool) {
	file, err := os.Open(filename)
	if err != nil {
		return "", false
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return "", false
	}

	for _, record := range records {
		if record[0] == key {
			return record[1], true
		}
	}

	return "", false
}
