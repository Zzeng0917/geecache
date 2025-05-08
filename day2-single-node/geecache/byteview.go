package geecache

type ByteView struct {
	b []byte
}

func (v ByteView) String() string {
	return string(v.b)
	//注意：与原文不同，这里直接转换为字符串（使用原文在测试时出现问题）
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
