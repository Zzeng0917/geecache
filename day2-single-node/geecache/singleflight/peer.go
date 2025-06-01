package geecache

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

//根据传入的key选择相应的节点PeerGetter

type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}

//从对应group查找缓存值
