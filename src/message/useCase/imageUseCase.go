package useCase

import (
	"bufio"
	"encoding/base64"
	"io"
	"math/rand"
	"os"

	multipart "mime/multipart"

	"github.com/spf13/viper"
)

type ImageUseCase interface {
	RecoverImage(string) string
	UpLoadImage(*multipart.FileHeader) string
}

type imageUseCase struct {
}

func (imageUseCase) UpLoadImage(file *multipart.FileHeader) string {
	randomString := generateRandomString(6)
	filename := viper.GetString("imagesFilePath") + randomString

	src, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer src.Close()
	dst, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer dst.Close()
	io.Copy(dst, src)
	return filename
}
func (imageUseCase) RecoverImage(name string) string {
	filename := viper.GetString("imagesFilePath") + name
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fInfo, _ := file.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(file)
	fReader.Read(buf)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	defer file.Close()

	return "data:image/png;base64," + imgBase64Str
}

func NewImageUseCase() ImageUseCase {

	return imageUseCase{}

}
func generateRandomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)

}
