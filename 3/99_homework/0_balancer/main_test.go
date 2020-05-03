package main

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

type RoundRobinBalancer struct {
	stats []int
	serversCount int
	nextServerIndex int
	mux sync.Mutex
}

func (balancer *RoundRobinBalancer) Init(serversCount int) {
	balancer.serversCount = serversCount
	balancer.nextServerIndex = 0
	balancer.stats = make([]int, serversCount)
}

func (balancer *RoundRobinBalancer) GiveStat() []int {
	balancer.mux.Lock()
	defer balancer.mux.Unlock()

	return balancer.stats
}

func (balancer *RoundRobinBalancer) GiveNode() int {
	balancer.mux.Lock()
	defer balancer.mux.Unlock()

	balancer.stats[balancer.nextServerIndex]++
	serverToSendIndex := balancer.nextServerIndex
	balancer.nextServerIndex = (balancer.nextServerIndex + 1) % balancer.serversCount

	return serverToSendIndex
}

func TestRoundRobinBalancer(t *testing.T) {

	balancer := new(RoundRobinBalancer)
	balancer.Init(3)
	balancer.GiveStat()

	expected := []int{0, 0, 0}
	result := balancer.GiveStat()
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "have", result)
	}

	n := balancer.GiveNode()

	expected = []int{1, 0, 0}
	result = balancer.GiveStat()
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "have", result)
	}
	fmt.Println(n, expected)

	n = balancer.GiveNode()

	expected = []int{1, 1, 0}
	result = balancer.GiveStat()
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "have", result)
	}
	fmt.Println(n, expected)

}

func TestRoundRobinBalancerMany(t *testing.T) {

	tests := []struct {
		servers  int
		clients  int
		requests int
		expected []int
	}{
		{2, 1, 100, []int{50, 50}},
		{2, 100, 100, []int{5000, 5000}},
		{5, 100, 100, []int{2000, 2000, 2000, 2000, 2000}},
	}

	for _, test := range tests {
		wg := new(sync.WaitGroup)
		balancer := new(RoundRobinBalancer)
		balancer.Init(test.servers)
		for i := 0; i < test.clients; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for req := 0; req < test.requests; req++ {
					balancer.GiveNode()
				}
			}()
		}
		wg.Wait()

		expected := test.expected
		result := balancer.GiveStat()
		if !reflect.DeepEqual(result, expected) {
			t.Error("expected", expected, "have", result)
		}
	}

}