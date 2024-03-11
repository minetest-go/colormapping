package colormapping

import (
	"bufio"
	"bytes"
	"embed"
	"errors"
	"image/color"
	"strconv"
	"strings"
	"sync"
)

//go:embed colors/*
var Files embed.FS

type ColorMapping struct {
	colors               map[string]*color.RGBA
	extendedpaletteblock map[string]bool
	extendedpalette      *Palette
	mutex                *sync.RWMutex
}

var simpleNamedColors = map[string]*color.RGBA{
	"black":   {R: 32, G: 32, B: 32, A: 255},
	"blue":    {R: 59, G: 68, B: 93, A: 255},
	"grey":    {R: 150, G: 148, B: 149, A: 255},
	"gray":    {R: 150, G: 148, B: 149, A: 255},
	"green":   {R: 93, G: 112, B: 70, A: 255},
	"orange":  {R: 154, G: 116, B: 75, A: 255},
	"pink":    {R: 183, G: 150, B: 160, A: 255},
	"brown":   {R: 48, G: 39, B: 31, A: 255},
	"cyan":    {R: 60, G: 96, B: 100, A: 255},
	"magenta": {R: 118, G: 93, B: 119, A: 255},
	"red":     {R: 107, G: 54, B: 53, A: 255},
	"violet":  {R: 70, G: 52, B: 83, A: 255},
	"white":   {R: 231, G: 223, B: 225, A: 255},
	"yellow":  {R: 155, G: 136, B: 75, A: 255},
}

func maxUint8(i1, i2 uint8) uint8 {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

func (m *ColorMapping) GetColor(name string, param2 int) *color.RGBA {
	//TODO: list of node->palette
	if m.extendedpaletteblock[name] {
		// param2 coloring
		return m.extendedpalette.GetColor(param2)
	}

	m.mutex.RLock()
	c := m.colors[name]
	if c != nil {
		// perfect match found
		m.mutex.RUnlock()
		return c
	}
	m.mutex.RUnlock()

	// fall back to simple color-name matching
	for scname, sc := range simpleNamedColors {
		if strings.Contains(name, "dark_"+scname) || strings.Contains(name, scname+"_dark") {
			// dark color
			c = &color.RGBA{
				R: maxUint8(sc.R-30, 0),
				G: maxUint8(sc.G-30, 0),
				B: maxUint8(sc.B-30, 0),
			}
		} else if strings.Contains(name, scname+"_") || strings.Contains(name, "_"+scname) {
			// normal color
			c = sc
		}
	}

	if c != nil {
		// add name-color pair to map
		m.mutex.Lock()
		m.colors[name] = c
		m.mutex.Unlock()
	}

	return c
}

func (m *ColorMapping) GetColors() map[string]*color.RGBA {
	return m.colors
}

func (m *ColorMapping) LoadBytes(buffer []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(buffer))
	count := 0
	line := 0

	for scanner.Scan() {
		line++

		txt := strings.Trim(scanner.Text(), " ")

		if len(txt) == 0 {
			//empty
			continue
		}

		if strings.HasPrefix(txt, "#") {
			//comment
			continue
		}

		parts := strings.Fields(txt)

		if len(parts) < 4 {
			return 0, errors.New("invalid line: #" + strconv.Itoa(line))
		}

		if len(parts) >= 4 {
			r, err := strconv.ParseInt(parts[1], 10, 32)
			if err != nil {
				return 0, err
			}

			g, err := strconv.ParseInt(parts[2], 10, 32)
			if err != nil {
				return 0, err
			}

			b, err := strconv.ParseInt(parts[3], 10, 32)
			if err != nil {
				return 0, err
			}

			a := int64(255)

			c := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			m.colors[parts[0]] = &c
			count++
		}
	}

	return count, nil
}

func (m *ColorMapping) LoadVFSColors(filename string) (int, error) {
	buffer, err := Files.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	return m.LoadBytes(buffer)
}

func (m *ColorMapping) LoadDefaults() error {
	list := []string{
		"advtrains.txt",
		"custom.txt",
		"everness.txt",
		"mc2.txt",
		"miles.txt",
		"naturalbiomes.txt",
		"mtg.txt",
		"nodecore.txt",
		"scifi_nodes.txt",
		"vanessa.txt",
	}
	for _, name := range list {
		_, err := m.LoadVFSColors("colors/" + name)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewColorMapping() *ColorMapping {
	data, err := Files.ReadFile("colors/unifieddyes_palette_extended.png")
	if err != nil {
		panic(err)
	}

	extendedpalette, err := NewPalette(data)
	if err != nil {
		panic(err)
	}

	data, err = Files.ReadFile("colors/extended_palette.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	extendedpaletteblock := make(map[string]bool)

	for scanner.Scan() {
		txt := strings.Trim(scanner.Text(), " ")

		if len(txt) == 0 {
			//empty
			continue
		}

		extendedpaletteblock[txt] = true
	}

	return &ColorMapping{
		colors:               make(map[string]*color.RGBA),
		extendedpaletteblock: extendedpaletteblock,
		extendedpalette:      extendedpalette,
		mutex:                &sync.RWMutex{},
	}
}
