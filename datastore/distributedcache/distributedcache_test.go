package distributedcache

import (
	"flag"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestDistributedCache(t *testing.T) {
	var (
		listenAddr = flag.String("listenaddr", ":3000", "listening address of the server")
		leaderAddr = flag.String("leaderaddr", ":", "listening address of the leader")
	)
	flag.Parse()
	opts := ServerOpts{
		ListenAddr: *listenAddr,
		IsLeader:   true,
		LeaderAddr: *leaderAddr,
	}
	go func() {
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", opts.ListenAddr)
		if err != nil {
			t.Fatalf("Failed to connect to server: %v", err)
		}
		conn.Write([]byte("SET Foo Bar 250000"))

		time.Sleep(time.Second * 2)
		conn.Write([]byte("GET Foo Bar 250000"))

		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[:n]))
	}()

	server := NewRaftServer(opts, NewCache())
	server.Start()

}
