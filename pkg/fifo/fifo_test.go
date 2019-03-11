package fifo_test

import (
	"fmt"
	"github.com/go-chassis/huawei-apm/pkg/fifo"
	"github.com/go-mesh/openlogging"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWriter(t *testing.T) {
	w, err := fifo.NewWriter("app1", "service1")
	assert.NoError(t, err)
	t.Log("write")
	go func() {
		i := 0
		for i < 100 {
			fmt.Println("write string to named pipe file.")
			_, err = w.WriteString(fmt.Sprintf("test write times:%d\n", i))
			assert.NoError(t, err)
			w.Flush()
			i++
		}
	}()

	r, err := fifo.NewReader("app1", "service1")
	assert.NoError(t, err)
	i := 0
	for {
		openlogging.Info("reading")
		_, err := r.ReadBytes('\n')
		assert.NoError(t, err)
		i++
		if i == 100 {
			break
		}
	}
}
