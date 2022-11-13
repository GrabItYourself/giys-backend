package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/GrabItYourself/giys-backend/lib/logger"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
	"github.com/pkg/errors"
)

func CompressImage(buffer string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(buffer)
	if err != nil {
		return "", errors.Wrap(err, "can't decode base64 image")
	}
	logger.Debug(fmt.Sprintf("(image) (pre) len: %d", len(decoded)))

	img, _, err := image.Decode(bytes.NewReader(decoded))
	if err != nil {
		return "", errors.Wrap(err, "can't create image from bytes")
	}
	logger.Debug(fmt.Sprintf("(image) (pre) width: %d, height: %d", img.Bounds().Dx(), img.Bounds().Dy()))

	img, err = cutter.Crop(img, cutter.Config{
		Width:   1,
		Height:  1,
		Mode:    cutter.Centered,
		Options: cutter.Ratio,
	})
	if err != nil {
		return "", errors.Wrap(err, "can't crop image")
	}

	img = resize.Thumbnail(500, 500, img, resize.Lanczos3)
	logger.Debug(fmt.Sprintf("(image) (post) width: %d, height: %d", img.Bounds().Dx(), img.Bounds().Dy()))

	var buf bytes.Buffer

	err = jpeg.Encode(&buf, img, nil)
	if err != nil {
		return "", errors.Wrap(err, "can't encode image as jpg")
	}
	logger.Debug(fmt.Sprintf("(image) (post) len: %d", len(buf.Bytes())))

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
