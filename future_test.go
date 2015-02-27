package future

import "testing"
import "sync"

type Data struct {
	x int
}

func TestCompleteAwait(t *testing.T) {
	producer, consumer := Pair()
	var data int = 0
	data_handle := &data

	*data_handle = 8
	producer.Complete()

	consumer.Await()
	if *data_handle != 8 {
		t.Fatal("data_handle was set to", *data_handle, "after await")
	}
}

func TestAwaitComplete(t *testing.T) {
	producer, consumer := Pair()

	data := 0
	data_handle := &data

	mutex := &sync.Mutex{}

	mutex.Lock()
	go func() {
		mutex.Unlock()
		consumer.Await()

		if *data_handle != 8 {
			t.Fatal("data_handle was set to", *data_handle, "after await")
		}
	}()

	mutex.Lock()
	*data_handle = 8
	producer.Complete()
}
