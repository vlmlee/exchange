package lsm

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
