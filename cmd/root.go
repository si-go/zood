package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	// Use: ,
	Long: `this is a simple tool for share youre files in local network over TCP/IP
for more info usege : zood -h, --help`,
}

var (
	path string
	port string
)

func Exexute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//var (
	//	err    error
	//	//path   string
	//	router *mux.Router
	//	//port  string
	//	ipv4 string
	//	conn net.Conn
	//	//defaultPort = "4242"
	//)
	//
	//
	//
	//router = mux.NewRouter()
	////fmt.Print("Enter the path you want share: ")
	////_, err = fmt.Scanln(&path)
	////if err != nil {
	////	fmt.Println(err)
	////	os.Exit(1)
	////}
	////
	////fmt.Printf("Enter the port (Default: %s just enter y/Y ;) ", defaultPort)
	////_, err = fmt.Scan(&port)
	////if err != nil {
	////	fmt.Println(err)
	////	os.Exit(1)
	////}
	////
	////if port == "y" || port == "Y" {
	////	port = defaultPort
	////}
	////
	//
	//conn, err = net.Dial("udp", "8.8.8.8:80")
	//// handle err...
	//
	//defer conn.Close()
	//ipv4 = conn.LocalAddr().(*net.UDPAddr).IP.String()
	//
	//router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(path))))
	//log.Printf("Local host: http://localhost:%s\n", port)
	//log.Printf("Local network: http://%s:%s\n", ipv4, port)
	//
	//err = http.ListenAndServe(":" + port, router)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(path)
	//fmt.Println(port)
}

func init() {
	rootCmd.SetHelpTemplate(`https://github.com/si-go/zood
author: Sina Meshkini (https://sinameshkini.ir - https://sottabyte.ir)
this is a simple tool for share youre files in local network over TCP/IP
for more info usege : zood -h, --help
	--path -d  path of what you want share, file or directory/ (default .)
	--port -p  port (default 4242)
	
	Example:
		zood -d $HOME/Public -p 3666
		zood -d C:\Users

`)
	rootCmd.PersistentFlags().StringVarP(&path, "path", "d", ".", "path of what you want share, file or directory/")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "4242", "port")
}
