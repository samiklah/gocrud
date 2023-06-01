package main

import (
	"context"
	"fmt"
	"time"
)

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*100)
	defer cancel()
	respChan := make(chan Response)

	go func() {
		val, err := fetcheThearedParty()
		respChan <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data took too long")
		case resp := <-respChan:
			return resp.value, resp.err
		}
	}
}

func fetcheThearedParty() (int, error) {
	time.Sleep(time.Second * 500)
	return 1, nil
}

func main1() {
	start := time.Now()
	ctx := context.Background()
	val, err := fetchUserData(ctx, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("val", val)
	fmt.Println("took", time.Since(start))

}
