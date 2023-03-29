package service

import (
	"reflect"
	"testing"

	"github.com/jonnaylang101/image-resizer/pkg/core/domain"
)

func Test_service_Resize_ValidateInputs(t *testing.T) {
	sv := service{}
	t.Run("when no storage paths are provided, we should receive an error and empty response", func(t *testing.T) {
		res, err := sv.Resize(200, 200, "thumbnail")
		expectError(true, err, t)
		expectErrorMsg(ErrNoProvidedStoragePaths, err, t)
		expectResponse(domain.ResizeResponse{}, res, t)
	})

	t.Run("when the width is zero, we should receive an error and empty response", func(t *testing.T) {
		res, err := sv.Resize(0, 200, "thumbnail", "some/storage/path.jpg")
		expectError(true, err, t)
		expectErrorMsg(ErrInvalidWidth, err, t)
		expectResponse(domain.ResizeResponse{}, res, t)
	})

	t.Run("when the height is zero, we should receive an error and empty response", func(t *testing.T) {
		res, err := sv.Resize(200, 0, "thumbnail", "some/storage/path.jpg")
		expectError(true, err, t)
		expectErrorMsg(ErrInvalidHeight, err, t)
		expectResponse(domain.ResizeResponse{}, res, t)
	})

	t.Run("when the filenameSuffix is empty, we should default it to --resized-<width>-<height>", func(t *testing.T) {
		wantRes := domain.ResizeResponse{
			ResizedImagesStoragePaths: []string{"some/storage/path--resized-200-300.jpg", "another/files/path--resized-200-300.jpg"},
		}
		res, err := sv.Resize(200, 300, "", "some/storage/path.jpg", "another/files/path.jpg")
		expectError(false, err, t)
		expectResponse(wantRes, res, t)
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

func expectResponse(wantResponse, gotResponse domain.ResizeResponse, t *testing.T) {
	if !reflect.DeepEqual(wantResponse, gotResponse) {
		t.Errorf("wanted response to be... \n%+v\n but got \n%+v\n", wantResponse, gotResponse)
	}
}

func Test_addSuffix(t *testing.T) {
	type args struct {
		origPath string
		suffix   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "when the suffix is empty we get back the original file path",
			args: args{
				origPath: "some/path/file.jpg",
				suffix:   "",
			},
			want: "some/path/file.jpg",
		},
		{
			name: "when the suffix is `-hello` we get should append it just before the extension",
			args: args{
				origPath: "some/path/file.jpg",
				suffix:   "-hello",
			},
			want: "some/path/file-hello.jpg",
		},
		{
			name: "when the filepath has no extension, we should still append the suffix",
			args: args{
				origPath: "some/path/file",
				suffix:   "-hello",
			},
			want: "some/path/file-hello",
		},
		{
			name: "when the storage path is empty, we should get an empty string back",
			args: args{
				origPath: "",
				suffix:   "-hello",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addSuffix(tt.args.origPath, tt.args.suffix); got != tt.want {
				t.Errorf("addSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_addSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = addSuffix("some/path/file.jpg", "-meow")
	}
}
