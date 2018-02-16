package layout

import "testing"

import "reflect"
import "image/color"

import "github.com/fyne-io/fyne/ui"

func TestMaxLayout(t *testing.T) {
	size := ui.NewSize(100, 100)

	obj := ui.NewRectangle(color.RGBA{0, 0, 0, 0})
	container := &ui.Container{
		Size:    size,
		Objects: []ui.CanvasObject{obj},
	}

	MaxLayout(container, size)

	if !reflect.DeepEqual(obj.Size, size) {
		t.Fatalf("Expected %s but got %s", size, obj.Size)
	}
}
