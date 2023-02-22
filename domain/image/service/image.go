package service

type Image struct {
	Destination string
}

func NewImage(destination string) *Image {
	return &Image{
		Destination: destination,
	}
}
