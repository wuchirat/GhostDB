package base

import (
	"testing"

	"github.com/ghostdb/ghostdb-cache-node/config"
	"github.com/ghostdb/ghostdb-cache-node/utils"
	"github.com/ghostdb/ghostdb-cache-node/store/request"
)

func TestStore(t *testing.T) {

	conf := config.InitializeConfiguration()

	var store *Store
	store = NewStore("LRU")
	store.BuildStore(conf)
	store.RunStore()

	x := store.Execute("get", request.NewRequestFromValues("Key1", "NewValue1", -1))
	utils.AssertEqual(t, x.Message, "CACHE_MISS", "")

	x = store.Execute("put", request.NewRequestFromValues("Key1", "NewValue1", -1))
	utils.AssertEqual(t, x.Status, int32(1), "")
}