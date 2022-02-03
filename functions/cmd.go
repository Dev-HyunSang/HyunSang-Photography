package functions

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type UploadRequest struct {
}

/* Function
Photos UPLOAD, Photos Download, Photo Info
*/

// Welecome Message
func HelloHandler(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "안녕하세요. 뷰파인더로 세상을 담는 박현상입니다.",
		"time":    time.Now(),
	})
}

func AdminUploadHandler(c *fiber.Ctx) error {
	return c.Status(200).JSON("TEST")
}
