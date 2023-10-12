package main

import (
	"net"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	serverAddr := "localhost:12345"
	listener, err := startTestServer(serverAddr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	go func() {
		os.Exit(m.Run())
	}()

	time.Sleep(1 * time.Second) // Небольшая задержка для запуска сервера
	os.Exit(0)
}

func TestTelnetClient(t *testing.T) {
	input := "Test Input\n"
	expected := input

	actual := runTelnetClient("localhost:12345", input)

	if actual != expected {
		t.Errorf("Ожидалось: %s, получено: %s", expected, actual)
	}
}

func runTelnetClient(address, input string) string {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err.Error()
	}
	defer conn.Close()

	_, err = conn.Write([]byte(input))
	if err != nil {
		return err.Error()
	}

	buf := make([]byte, 1024) // буфер для чтения клиентских данных
	n, err := conn.Read(buf)
	if err != nil {
		return err.Error()
	}
	return string(buf[:n])
}

func startTestServer(serverAddr string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		return nil, err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}

	go func() {
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go handleTestConn(conn)
		}
	}()
	return listener, nil
}

func handleTestConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024) // буфер для чтения клиентских данных
	for {
		n, err := conn.Read(buf) // читаем из сокета в buf
		if err != nil {
			return
		}
		conn.Write(buf[:n])
	}
}
