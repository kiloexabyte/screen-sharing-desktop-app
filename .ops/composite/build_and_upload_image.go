package composite

import "fmt"

func (o Ops) BuildAndUploadImage() error {
	if err := o.BuildImage(); err != nil {
		return fmt.Errorf("build_image: %w", err)
	}

	if err := o.UploadImage(); err != nil {
		return fmt.Errorf("upload_image: %w", err)
	}

	return nil
}
