# 有关 goruntine 与 互斥锁   
**sync.Mutex**是 Go 语言标准库提供的一个互斥锁，当一个协程(goroutine)获得了这个锁的拥有权后，其它
请求锁的协程(goroutine) 就会阻塞在 Lock() 方法的调用上，直到调用 Unlock() 锁被释放。

**问题本质**：多个 goroutine 同时读写共享的map变量set，但 Go 的map不是并发安全的。当两个或多个goroutine
同时检测到!exist（即num不在map中），它们可能都会执行fmt.Println并尝试向map写入相同的键。

**可能的执行路径**：
多次输出：如果多个 goroutine 在map写入完成前都读取到!exist，它们都会打印100。例如，若 4 个 goroutine
同时检测到!exist并打印，最终map中只会有一个键值对，但输出了 4 次。
单次输出：若某个 goroutine 先完成写入，其他 goroutine 后续读取时会发现exist为true，从而不再打印。



# crc32.ChecksumIEEE算法

**CRC（循环冗余校验）** 是一种常用的错误检测算法，通过生成校验和（Checksum）来验证数据传输或存储的完整性。CRC32.ChecksumIEEE是 CRC32 算法的一种具体实现，基于 IEEE 标准（通常指 IEEE 802.3 标准），广泛应用于网络协议（如以太网、ZIP 压缩、PNG 图像等）中
