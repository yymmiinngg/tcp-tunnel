package wan

import (
	"fmt"
	"net"
	"os"

	"github.com/yymmiinngg/goargs"
)

/*


公网端：

clientConn -> client_connection_port <==> lan_connection_port -> lanConn

*/

var lanConns = make(chan net.Conn, 1024)

func Start() {

	var argsArr = os.Args
	template := `
    Usage: {{COMMAND}} WAN {{OPTION}}

	+ -a, --application-port  # 开放的应用程序端口（ip:port, 默认：127.0.0.1:80）
	* -s, --server-port       # 开放的服务端口（ip:port）
	+ -i, --io-timeout        # 读写超时时长（单位：秒，默认：120)

    ? -h, --help              # 显示帮助后退出
    ? -v, --version           # 显示版本后退出
`

	// 定义变量
	var applicationAddress string
	var serverAddress string
	var ioTimeout int

	// 编译模板
	args, err := goargs.Compile(template)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 绑定变量
	args.StringOption("-a", &applicationAddress, "127.0.0.1:80")
	args.StringOption("-s", &serverAddress, "")
	args.IntOption("-i", &ioTimeout, 120)

	// 处理参数
	err = args.Parse(argsArr)

	// 显示帮助
	if args.HasItem("-h", "--help") {
		fmt.Println(args.Usage())
		return
	}

	// 显示版本
	if args.HasItem("-v", "--version") {
		fmt.Println("v0.0.1")
		return
	}

	// 错误输出
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("--------------------------------------------------")
		fmt.Println(args.Usage())
		return
	}

	serverInfo := MakeServerInfo(serverAddress, applicationAddress, ioTimeout)
	serverInfo.StartServer()
}