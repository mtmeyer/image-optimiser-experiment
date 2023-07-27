package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/h2non/bimg"
)

func main() {
	app := fiber.New()

	app.Get("/img/*", func(c *fiber.Ctx) error {
    url := c.Params("*")
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		defer res.Body.Close()

		decodedImage, _, err := image.Decode(res.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		buffer := new(bytes.Buffer)
		encodeErr := jpeg.Encode(buffer, decodedImage, nil)
		if encodeErr != nil {
			fmt.Fprintln(os.Stderr, encodeErr)
		}

		imageSize, err := bimg.NewImage(buffer.Bytes()).Size()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		heightResizeRatio := 600 / imageSize.Width

		newImage, err := bimg.NewImage(buffer.Bytes()).Resize(600, (600 * heightResizeRatio))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		convert, err := bimg.NewImage(newImage).Convert(bimg.WEBP)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		reader := bytes.NewReader(convert)

		c.Set(fiber.HeaderContentType, "image/webp")
		return c.SendStream(reader)
	})

  port := GetEnv("PORT", "3000")

	app.Listen(port)

}
