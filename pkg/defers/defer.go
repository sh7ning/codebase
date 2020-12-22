package defers

var (
	globalDefers = newStack()
)

// Register 注册一个defer函数
func Register(fns ...func() error) {
	globalDefers.push(fns...)
}

// Run 运行
func Run() {
	globalDefers.run()
}
