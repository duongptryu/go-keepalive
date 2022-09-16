package curl

import (
	"context"
	"log"
	"net/http/httptrace"
	"sync"
	"testing"
)

func TestClient_GetAPI(t *testing.T) {
	NewClient()
	wg := sync.WaitGroup{}
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) { log.Printf("conn: %#v -- Reuse %v", info.Conn.LocalAddr(), info.Reused) },
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	for i := 0; i <= 100; i ++ {
		wg.Add(1)
		//time.Sleep(5 *time.Second)
		go func() {
			defer wg.Done()
			if _, err := GetAPI(traceCtx, "http://localhost:8080"); err != nil {
				log.Println(err)
			}
		}()
	}
	wg.Wait()
}

func TestClient_PostAPI(t *testing.T) {
	NewClient()
	wg := sync.WaitGroup{}
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) { log.Printf("conn: %#v -- Reuse %v", info.Conn.LocalAddr(), info.Reused) },
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	for i := 0; i <= 100; i ++ {
		wg.Add(1)
		//time.Sleep(5 *time.Second)
		go func() {
			defer wg.Done()
			if _, err := PostAPI(traceCtx, "http://localhost:8080"); err != nil {
				log.Println(err)
			}
		}()
	}
	wg.Wait()
}