package activeLists

import (
	"github.com/google/btree"
	"time"
)

type expirationTree struct {
	tree *btree.BTree
}

func (r *expirationTree) touch(key string, expiresAt int64) {
	if expiresAt > 0 {
		r.tree.ReplaceOrInsert(&expirationTreeItem{key: key, expiresAt: expiresAt})
	}
}

func (r *expirationTree) del(key string) {
	r.tree.Delete(&expirationTreeItem{key: key})
}

func (r *expirationTree) getExpiredKeys() (keys []string) {
	pivot := &expirationTreeItem{expiresAt: time.Now().UnixNano()}
	r.tree.AscendLessThan(pivot, func(i btree.Item) bool {
		keys = append(keys, i.(*expirationTreeItem).key)
		return true
	})
	return
}

type expirationTreeItem struct {
	key       string
	expiresAt int64
}

func (r *expirationTreeItem) Less(than btree.Item) bool {
	inputItem := than.(*expirationTreeItem)
	if r.key == inputItem.key {
		return false
	}
	return r.expiresAt < inputItem.expiresAt
}

func createExpirationTree() expirationTree {
	return expirationTree{tree: btree.New(64)}
}
