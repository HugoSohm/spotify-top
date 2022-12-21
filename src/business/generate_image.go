package business

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func generateImage(w http.ResponseWriter, r *http.Request, data []string, dataType string) {
	// Load background image
	image, err := gg.LoadJPG("assets/background.jpg")
	if err != nil {
		NewError(w, http.StatusInternalServerError, "Failed to load image")
		return
	}

	// Create image context and set font color
	imageCtx := gg.NewContextForImage(image)
	imageCtx.SetRGB(1, 1, 1)

	// Load title font
	if err = imageCtx.LoadFontFace("assets/unbounded-bold.ttf", 30); err != nil {
		NewError(w, http.StatusInternalServerError, "Failed to load font")
		return
	}

	// Draw Title
	imageCtx.DrawStringAnchored(fmt.Sprintf("Mes 10 %s les plus écoutés en %s", dataType, MonthNameToFrench(time.Now().Month().String())), 1500/2, 150, 0.5, 0.5)

	// Load body font
	if err = imageCtx.LoadFontFace("assets/unbounded-regular.ttf", 30); err != nil {
		NewError(w, http.StatusInternalServerError, "Failed to load font")
		return
	}

	// Write body
	for i, s := range data[:10] {
		imageCtx.DrawStringAnchored(s, 1500/2, 300+60*float64(i), 0.5, 0.5)
	}

	// Save image
	imageName := fmt.Sprintf("%s.png", uuid.New().String())
	if err = imageCtx.SavePNG(fmt.Sprintf("images/%s", imageName)); err != nil {
		NewError(w, http.StatusInternalServerError, "Failed to save image")
	}

	http.Redirect(w, r, fmt.Sprintf("/%s", imageName), http.StatusSeeOther)
}
