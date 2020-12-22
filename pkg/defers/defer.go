package defers

var (
	globalDefers = NewStack()
)

// Register 注册一个defer函数
func Register(fns ...func() error) {
	globalDefers.Push(fns...)
}

// Run 运行
func Run() {
	globalDefers.Run()
}
