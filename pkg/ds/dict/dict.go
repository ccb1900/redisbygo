package dict

import "sync"

type Dict struct {
	storage sync.Map
}

func Store(d *Dict, key *interface{}, value *interface{}) {
	d.storage.Store(key, value)
}

func Get(d *Dict, key *interface{}) interface{} {
	value, _ := d.storage.Load(key)
	return value
}
