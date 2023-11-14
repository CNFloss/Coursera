package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct {
	id int
	mu sync.Mutex
}

type Philosopher struct {
	id                        int
	leftChopstick, rightChopstick *Chopstick
	eatCount                  int
}

type Host struct {
	permissionChannel chan [2]*Philosopher
	doneEatingChannel chan *Philosopher
}

func NewHost() *Host {
	return &Host{
		permissionChannel: make(chan [2]*Philosopher),
		doneEatingChannel: make(chan *Philosopher),
	}
}

func (host *Host) manage(philosophers []*Philosopher) {
    eatCount := make(map[int]int)
    for {
        allDone := true
        for _, philosopher := range philosophers {
            if eatCount[philosopher.id] < 3 {
                allDone = false
                break
            }
        }
        if allDone {
            break
        }

        select {
        case pair := <-host.permissionChannel:
            if pair[0] != nil {
                pair[0].eatCount++
            }
            if pair[1] != nil {
                pair[1].eatCount++
            }
        case philosopher := <-host.doneEatingChannel:
            eatCount[philosopher.id]++
        }
    }
    close(host.permissionChannel)
}

func (philosopher *Philosopher) dine(host *Host) {
	for philosopher.eatCount < 3 {
		host.permissionChannel <- [2]*Philosopher{philosopher, nil}

		philosopher.leftChopstick.mu.Lock()
		philosopher.rightChopstick.mu.Lock()

		fmt.Printf("starting to eat %d\n", philosopher.id)
		time.Sleep(time.Second)
		fmt.Printf("finishing eating %d\n", philosopher.id)

		philosopher.leftChopstick.mu.Unlock()
		philosopher.rightChopstick.mu.Unlock()

		host.doneEatingChannel <- philosopher
	}
}

func main() {
	var wg sync.WaitGroup
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = &Chopstick{id: i + 1}
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%5], 0}
	}

	host := NewHost()
	go host.manage(philosophers)

	wg.Add(5)
	for _, philosopher := range philosophers {
		go func(p *Philosopher) {
			defer wg.Done()
			p.dine(host)
		}(philosopher)
	}
	wg.Wait()
}
