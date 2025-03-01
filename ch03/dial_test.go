package ch03

import (
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		defer func() {
			// 1:
			done <- struct{}{}
		}()

		for {
			conn, err := listener.Accept()
			if err != nil {
				// 리스너가 종료되면 오류가 발생한다.
				// 리스너의 종료로 인한 오류는 로깅하고 넘어가도 문제 없다.
				t.Log(err)
				return
			}
			go func(c net.Conn) {
				defer func() {
					c.Close()
					// 2:
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							// 커넥션이 정상적으로 종료하면 로깅한다.
							t.Error(err)
						}
						return
					}
					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	conn.Close()
	// 2: 을 기다린다.
	<-done
	// 1: 을 기다린다.
	listener.Close()
	<-done
}
