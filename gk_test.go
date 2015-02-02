/*
* @Author: Xier
* @Date:   2015-02-02 12:35:19
* @Last Modified by:   Xier
* @Last Modified time: 2015-02-02 16:22:50
 */

package gika

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadDocument(t *testing.T) {
	// Create io.Reader
	f, err := os.Open("./doc-example/file.pdf")
	defer func() {
		f.Close()
	}()
	if err != nil {
		panic(err)
	}

	// Create io.Writer
	ws := &bytes.Buffer{}

	err = DocToText(f, ws)
	assert.NoError(t, err)
	assert.True(t, ws.Len() > 0)
}

func TestIsSupport(t *testing.T) {
	sup := IsSupport("./example/file.doc")
	assert.True(t, sup)
}

func TestIsSupportFalse(t *testing.T) {
	sup := IsSupport("./example/file.docc")
	assert.False(t, sup)
}

func TestIsSupportNoPath(t *testing.T) {
	sup := IsSupport("file.docx")
	assert.True(t, sup)
}
