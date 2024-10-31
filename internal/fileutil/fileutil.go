package fileutil

import (
	"fmt"
	"mime/multipart"
)

func CheckValidProductImages(imgHeaders []*multipart.FileHeader) error {
	if len(imgHeaders) == 0 {
		return fmt.Errorf("no files uploaded")
	}

	if len(imgHeaders) > 3 {
		return fmt.Errorf("files uploaded must be 3 files at max")
	}

	for _, imgHeader := range imgHeaders {
		contentType := imgHeader.Header.Get("Content-Type")
		if contentType != "image/png" && contentType != "image/jpeg" {
			return fmt.Errorf("invalid image type (accepted .png or .jpeg)")
		}
	}

	return nil
}
