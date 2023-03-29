package inmem

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/jonnaylang101/image-resizer/pkg/core/ports"
)

func TestNewStore(t *testing.T) {
	cn := []byte("this is the file contents")
	sk := "somefile.jpg"
	t.Run("when we pass an invalid path to NewStore, we should get an error", func(t *testing.T) {
		_, err := NewStore("")
		expectError(true, err, t)
	})

	t.Run("when we pass a valid path to NewStore, it should return a store with valid memory path", func(t *testing.T) {
		td := t.TempDir()
		st, err := NewStore(td)
		expectError(false, err, t)
		expectMemoryPath(td, st, t)
	})

	t.Run("when we pass a valid path to NewStore, we want to be able to use it's methods", func(t *testing.T) {
		td := t.TempDir()
		st, err := NewStore(td)
		expectError(false, err, t)
		err = st.Add(sk, bytes.NewBuffer(cn))
		expectError(false, err, t)
		expectFileAtPathContents(cn, td+"/"+sk, t)
	})
}

func Test_storage_Add(t *testing.T) {
	st := new(storage)
	cn := []byte("this is the file contents")
	sk := "somefile.jpg"

	t.Run("when the storage path is missing we should receive an error", func(t *testing.T) {
		err := st.Add("", bytes.NewBuffer(cn))
		expectError(true, err, t)
		expectErrorMsg(ErrInvalidStoragePath, err, t)
	})

	t.Run("when the image file has nil value, we should receive an error", func(t *testing.T) {
		err := st.Add(sk, nil)
		expectError(true, err, t)
		expectErrorMsg(ErrInvalidImageFile, err, t)
	})

	t.Run("when the file has been added to the store, we should be able to retrieve it", func(t *testing.T) {
		st.memPath = t.TempDir()
		err := st.Add(sk, bytes.NewReader(cn))
		expectError(false, err, t)
		expectFileAtPathContents(cn, st.memPath+"/"+sk, t)
	})

	t.Run("when the file has already been added to the store, we should receive a duplication error", func(t *testing.T) {
		st.memPath = t.TempDir()
		err := st.Add(sk, bytes.NewReader(cn))
		expectError(false, err, t)

		err = st.Add(sk, bytes.NewReader(cn))
		expectError(true, err, t)
	})

	t.Run("when the storage path has parts, we should not receive an error", func(t *testing.T) {
		st.memPath = t.TempDir()
		sk := "this/has/parts/somefile.jpg"
		err := st.Add(sk, bytes.NewReader(cn))
		expectError(false, err, t)
	})
}

func Test_storage_GetByStoragePath(t *testing.T) {
	st := new(storage)
	sk := "somefile.jpg"
	cn := []byte("this is the file contents")

	t.Run("when the storage path is missing we should receive an error", func(t *testing.T) {
		_, err := st.GetByStoragePath("")
		expectError(true, err, t)
		expectErrorMsg(ErrInvalidStoragePath, err, t)
	})

	t.Run("when the file can't be found, we should receive an error", func(t *testing.T) {
		st.memPath = t.TempDir()
		_, err := st.GetByStoragePath("missingfile.jpg")
		expectError(true, err, t)
		expectErrorMsg(ErrFileNotFound, err, t)
	})

	t.Run("when the file is found, we should expect it back with no error", func(t *testing.T) {
		st.memPath = t.TempDir()
		err := st.Add(sk, bytes.NewBuffer(cn))
		expectError(false, err, t)
		f, err := st.GetByStoragePath(sk)
		expectError(false, err, t)
		expectFileContents(cn, f, t)
	})

	t.Run("when the filepath has parts in it, we should still expect it back with no error", func(t *testing.T) {
		st.memPath = t.TempDir()
		sk := "this/has/parts/somefile.jpg"
		err := st.Add(sk, bytes.NewBuffer(cn))
		expectError(false, err, t)
		f, err := st.GetByStoragePath(sk)
		expectError(false, err, t)
		expectFileContents(cn, f, t)
	})
}

func expectError(wantErr bool, err error, t *testing.T) {
	if (err != nil) != wantErr {
		t.Errorf("expected error to be %v, but got %v", wantErr, err)
	}
}

func expectErrorMsg(wantErrMsg string, err error, t *testing.T) {
	if err != nil && err.Error() != wantErrMsg {
		t.Errorf("wanted error msg to be `%s` but got `%s`", wantErrMsg, err)
	}
}

func expectFileAtPathContents(contents []byte, filepath string, t *testing.T) {
	b, err := os.ReadFile(filepath)
	expectError(false, err, t)
	if !reflect.DeepEqual(contents, b) {
		t.Errorf("expected contents to be '%s' but got '%s'", contents, b)
	}
}

func expectFileContents(contents []byte, file io.Reader, t *testing.T) {
	b, err := io.ReadAll(file)
	expectError(false, err, t)
	if !reflect.DeepEqual(contents, b) {
		t.Errorf("expected contents to be '%s' but got '%s'", contents, b)
	}
}

func expectMemoryPath(wantMemPath string, st ports.Storage, t *testing.T) {
	_st, ok := st.(*storage)
	if !ok {
		t.Fatal("can't assert to storage type")
	}
	if _st.memPath != wantMemPath {
		t.Errorf("expected memory path to be %s but got %s", wantMemPath, _st.memPath)
	}
}
