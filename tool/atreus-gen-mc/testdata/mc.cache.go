// Code generated by kratos tool mcgen. DO NOT EDIT.

/*
  Package testdata is a generated mc cache package.
  It is generated from:
  type _mc interface {
		// mc: -key=demoKey
		CacheDemos(c context.Context, keys []int64) (map[int64]*Demo, error)
		// mc: -key=demoKey
		CacheDemo(c context.Context, key int64) (*Demo, error)
		// mc: -key=keyMid
		CacheDemo1(c context.Context, key int64, mid int64) (*Demo, error)
		// mc: -key=noneKey
		CacheNone(c context.Context) (*Demo, error)
		// mc: -key=demoKey
		CacheString(c context.Context, key int64) (string, error)

		// mc: -key=demoKey -expire=d.demoExpire -encode=json
		AddCacheDemos(c context.Context, values map[int64]*Demo) error
		// mc: -key=demo2Key -expire=d.demoExpire -encode=json
		AddCacheDemos2(c context.Context, values map[int64]*Demo, tp int64) error
		// 这里也支持自定义注释 会替换默认的注释
		// mc: -key=demoKey -expire=d.demoExpire -encode=json|gzip
		AddCacheDemo(c context.Context, key int64, value *Demo) error
		// mc: -key=keyMid -expire=d.demoExpire -encode=gob
		AddCacheDemo1(c context.Context, key int64, value *Demo, mid int64) error
		// mc: -key=noneKey
		AddCacheNone(c context.Context, value *Demo) error
		// mc: -key=demoKey -expire=d.demoExpire
		AddCacheString(c context.Context, key int64, value string) error

		// mc: -key=demoKey
		DelCacheDemos(c context.Context, keys []int64) error
		// mc: -key=demoKey
		DelCacheDemo(c context.Context, key int64) error
		// mc: -key=keyMid
		DelCacheDemo1(c context.Context, key int64, mid int64) error
		// mc: -key=noneKey
		DelCacheNone(c context.Context) error
	}
*/

package testdata

import (
	"context"
	"fmt"

	"github.com/mapgoo-lab/atreus/pkg/cache/memcache"
	"github.com/mapgoo-lab/atreus/pkg/log"
)

var (
	_ _mc
)

// CacheDemos get data from mc
func (d *Dao) CacheDemos(c context.Context, ids []int64) (res map[int64]*Demo, err error) {
	l := len(ids)
	if l == 0 {
		return
	}
	keysMap := make(map[string]int64, l)
	keys := make([]string, 0, l)
	for _, id := range ids {
		key := demoKey(id)
		keysMap[key] = id
		keys = append(keys, key)
	}
	replies, err := d.mc.GetMulti(c, keys)
	if err != nil {
		log.Errorv(c, log.KV("CacheDemos", fmt.Sprintf("%+v", err)), log.KV("keys", keys))
		return
	}
	for _, key := range replies.Keys() {
		v := &Demo{}
		err = replies.Scan(key, v)
		if err != nil {
			log.Errorv(c, log.KV("CacheDemos", fmt.Sprintf("%+v", err)), log.KV("key", key))
			return
		}
		if res == nil {
			res = make(map[int64]*Demo, len(keys))
		}
		res[keysMap[key]] = v
	}
	return
}

// CacheDemo get data from mc
func (d *Dao) CacheDemo(c context.Context, id int64) (res *Demo, err error) {
	key := demoKey(id)
	res = &Demo{}
	if err = d.mc.Get(c, key).Scan(res); err != nil {
		res = nil
		if err == memcache.ErrNotFound {
			err = nil
		}
	}
	if err != nil {
		log.Errorv(c, log.KV("CacheDemo", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// CacheDemo1 get data from mc
func (d *Dao) CacheDemo1(c context.Context, id int64, mid int64) (res *Demo, err error) {
	key := keyMid(id, mid)
	res = &Demo{}
	if err = d.mc.Get(c, key).Scan(res); err != nil {
		res = nil
		if err == memcache.ErrNotFound {
			err = nil
		}
	}
	if err != nil {
		log.Errorv(c, log.KV("CacheDemo1", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// CacheNone get data from mc
func (d *Dao) CacheNone(c context.Context) (res *Demo, err error) {
	key := noneKey()
	res = &Demo{}
	if err = d.mc.Get(c, key).Scan(res); err != nil {
		res = nil
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
	}
	if err != nil {
		log.Errorv(c, log.KV("CacheNone", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// CacheString get data from mc
func (d *Dao) CacheString(c context.Context, id int64) (res string, err error) {
	key := demoKey(id)
	err = d.mc.Get(c, key).Scan(&res)
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.Errorv(c, log.KV("CacheString", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheDemos Set data to mc
func (d *Dao) AddCacheDemos(c context.Context, values map[int64]*Demo) (err error) {
	if len(values) == 0 {
		return
	}
	for id, val := range values {
		key := demoKey(id)
		item := &memcache.Item{Key: key, Object: val, Expiration: d.demoExpire, Flags: memcache.FlagJSON}
		if err = d.mc.Set(c, item); err != nil {
			log.Errorv(c, log.KV("AddCacheDemos", fmt.Sprintf("%+v", err)), log.KV("key", key))
			return
		}
	}
	return
}

// AddCacheDemos2 Set data to mc
func (d *Dao) AddCacheDemos2(c context.Context, values map[int64]*Demo, tp int64) (err error) {
	if len(values) == 0 {
		return
	}
	for id, val := range values {
		key := demo2Key(id, tp)
		item := &memcache.Item{Key: key, Object: val, Expiration: d.demoExpire, Flags: memcache.FlagJSON}
		if err = d.mc.Set(c, item); err != nil {
			log.Errorv(c, log.KV("AddCacheDemos2", fmt.Sprintf("%+v", err)), log.KV("key", key))
			return
		}
	}
	return
}

// AddCacheDemo 这里也支持自定义注释 会替换默认的注释
func (d *Dao) AddCacheDemo(c context.Context, id int64, val *Demo) (err error) {
	if val == nil {
		return
	}
	key := demoKey(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.demoExpire, Flags: memcache.FlagJSON | memcache.FlagGzip}
	if err = d.mc.Set(c, item); err != nil {
		log.Errorv(c, log.KV("AddCacheDemo", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheDemo1 Set data to mc
func (d *Dao) AddCacheDemo1(c context.Context, id int64, val *Demo, mid int64) (err error) {
	if val == nil {
		return
	}
	key := keyMid(id, mid)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.demoExpire, Flags: memcache.FlagGOB}
	if err = d.mc.Set(c, item); err != nil {
		log.Errorv(c, log.KV("AddCacheDemo1", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheNone Set data to mc
func (d *Dao) AddCacheNone(c context.Context, val *Demo) (err error) {
	if val == nil {
		return
	}
	key := noneKey()
	item := &memcache.Item{Key: key, Object: val, Expiration: d.demoExpire, Flags: memcache.FlagJSON}
	if err = d.mc.Set(c, item); err != nil {
		log.Errorv(c, log.KV("AddCacheNone", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheString Set data to mc
func (d *Dao) AddCacheString(c context.Context, id int64, val string) (err error) {
	if len(val) == 0 {
		return
	}
	key := demoKey(id)
	bs := []byte(val)
	item := &memcache.Item{Key: key, Value: bs, Expiration: d.demoExpire, Flags: memcache.FlagRAW}
	if err = d.mc.Set(c, item); err != nil {
		log.Errorv(c, log.KV("AddCacheString", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DelCacheDemos delete data from mc
func (d *Dao) DelCacheDemos(c context.Context, ids []int64) (err error) {
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		key := demoKey(id)
		if err = d.mc.Delete(c, key); err != nil {
			if err == memcache.ErrNotFound {
				err = nil
				continue
			}
			log.Errorv(c, log.KV("DelCacheDemos", fmt.Sprintf("%+v", err)), log.KV("key", key))
			return
		}
	}
	return
}

// DelCacheDemo delete data from mc
func (d *Dao) DelCacheDemo(c context.Context, id int64) (err error) {
	key := demoKey(id)
	if err = d.mc.Delete(c, key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.Errorv(c, log.KV("DelCacheDemo", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DelCacheDemo1 delete data from mc
func (d *Dao) DelCacheDemo1(c context.Context, id int64, mid int64) (err error) {
	key := keyMid(id, mid)
	if err = d.mc.Delete(c, key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.Errorv(c, log.KV("DelCacheDemo1", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DelCacheNone delete data from mc
func (d *Dao) DelCacheNone(c context.Context) (err error) {
	key := noneKey()
	if err = d.mc.Delete(c, key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		log.Errorv(c, log.KV("DelCacheNone", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}
