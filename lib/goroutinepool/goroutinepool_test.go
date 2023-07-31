package goroutinepool

import (
	"errors"
	"sync"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGorourinepool_without_error(t *testing.T) {
	pool, err := NewGoroutinePool("TEST-POOL", 2, 10)
	assert.Nil(t, err)
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		idx := i
		pool.Put("TASK", func() error {
			defer wg.Done()
			log.Infof("TASK ID:%d", idx+1)
			return nil
		})
	}
	wg.Wait()
	pool.Close()
}

func TestGorourinepool_with_error(t *testing.T) {
	pool, err := NewGoroutinePool("TEST-POOL", 1, 10)
	assert.Nil(t, err)
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		index := i
		pool.Put("TASK", func() error {
			defer wg.Done()
			log.Infof("TASK ID:%d", index+1)
			if index == 0 {
				return errors.New("task execute failed")
			}
			return nil
		})
	}
	wg.Wait()
	pool.Close()
}

func TestGorourinepool_with_timeout(t *testing.T) {
	pool, err := NewGoroutinePool("TEST-POOL", 3, 10)
	assert.Nil(t, err)
	for i := 0; i < 3; i++ {
		index := i
		pool.PutWithTimeout("TASK", func() error {
			time.Sleep(time.Millisecond * 10)
			log.Infof("TASK ID:%d", index+1)
			return nil
		}, time.Millisecond)
	}

	t1 := time.Now()
	pool.Close()
	dur := time.Now().Sub(t1)
	log.Infof("pool exited, duration:%+v", dur)
	assert.True(t, dur < time.Millisecond*5)
}

func TestGorourinepool_without_timeout(t *testing.T) {
	pool, err := NewGoroutinePool("TEST-POOL-WITHOUT-TIMEOUT", 3, 10)
	assert.Nil(t, err)
	for i := 0; i < 3; i++ {
		index := i
		pool.Put("TASK", func() error {
			time.Sleep(time.Millisecond * 10)
			log.Infof("TASK ID:%d", index+1)
			return nil
		})
	}

	t1 := time.Now()

	time.Sleep(time.Millisecond)

	pool.Close()
	dur := time.Now().Sub(t1)
	log.Infof("pool exited, duration:%+v", dur)
	assert.True(t, dur >= time.Millisecond*10)
}
