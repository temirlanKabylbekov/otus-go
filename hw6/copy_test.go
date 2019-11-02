package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type TempFile struct {
	prefixName string
	content    []byte
	file       *os.File
}

func (tf *TempFile) create() error {
	tmpFile, err := ioutil.TempFile(os.TempDir(), tf.prefixName)
	tf.file = tmpFile

	if err != nil {
		return err
	}
	return nil
}

func (tf *TempFile) fill() error {
	if _, err := tf.file.Write(tf.content); err != nil {
		return err
	}
	return nil
}

func (tf *TempFile) clean() error {
	defer os.Remove(tf.file.Name())

	if err := tf.file.Close(); err != nil {
		return err
	}
	return nil
}

func newTempFile(prefixName string, content []byte) (*TempFile, error) {
	tf := new(TempFile)
	tf.content = content
	tf.prefixName = prefixName

	if err := tf.create(); err != nil {
		return tf, err
	}
	if err := tf.fill(); err != nil {
		return tf, err
	}
	return tf, nil
}

func TestSrcFileNotExist(t *testing.T) {
	dest := os.TempDir() + t.Name()
	err := Copy("not/existing/file.txt", dest, READ_ALL_THE__REST_OF_THE_FILE, 0)

	require.EqualError(t, err, "open not/existing/file.txt: no such file or directory")
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		t.Errorf("file is not copied but destination file was created")
	}
}

func TestDestFileAlreadyExists(t *testing.T) {
	tfFrom, err := newTempFile(t.Name(), []byte("This content of copying file"))
	defer tfFrom.clean()
	require.Nil(t, err)

	tfTo, err := newTempFile(t.Name(), []byte("Previous content for overriding by copy in dest file"))
	defer tfTo.clean()
	require.Nil(t, err)

	err = Copy(tfFrom.file.Name(), tfTo.file.Name(), READ_ALL_THE__REST_OF_THE_FILE, 0)
	require.Nil(t, err)

	copied, err := ioutil.ReadFile(tfTo.file.Name())
	require.Nil(t, err)
	require.Equal(t, []byte("This content of copying file"), copied)
}

func TestCopyOffsetIsTooLarge(t *testing.T) {
	dest := os.TempDir() + t.Name()
	defer os.Remove(dest)

	tf, err := newTempFile(t.Name(), []byte("This content of copying file"))
	defer tf.clean()
	require.Nil(t, err)

	err = Copy(tf.file.Name(), dest, READ_ALL_THE__REST_OF_THE_FILE, 100500)

	require.EqualError(t, err, "offset more or equal than file size")
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		t.Errorf("file is not copied but destination file was created")
	}
}

func TestPassOffsetEqualFileSize(t *testing.T) {
	dest := os.TempDir() + t.Name()
	defer os.Remove(dest)

	tf, err := newTempFile(t.Name(), []byte("This content of copying file"))
	defer tf.clean()
	require.Nil(t, err)

	err = Copy(tf.file.Name(), dest, READ_ALL_THE__REST_OF_THE_FILE, 100500)

	require.EqualError(t, err, "offset more or equal than file size")
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		t.Errorf("file is not copied but destination file was created")
	}
}

func TestTryToCopyInstantlyGrowingFile(t *testing.T) {
	dest := os.TempDir() + t.Name()
	defer os.Remove(dest)

	err := Copy("/dev/urandom", os.TempDir()+t.Name(), READ_ALL_THE__REST_OF_THE_FILE, 0)

	require.EqualError(t, err, "could not determine file size")
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		t.Errorf("file is not copied but destination file was created")
	}
}

func TestCopyFileSlice(t *testing.T) {
	dest := os.TempDir() + t.Name()
	defer os.Remove(dest)

	tf, err := newTempFile(t.Name(), []byte("This content of copying file"))
	defer tf.clean()
	require.Nil(t, err)

	err = Copy(tf.file.Name(), os.TempDir()+t.Name(), 7, 5)

	copied, err := ioutil.ReadFile(dest)
	require.Nil(t, err)
	require.Equal(t, []byte("content"), copied)
}
