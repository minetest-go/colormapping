package colormapping_test

import (
	"testing"

	"github.com/minetest-go/colormapping"
	"github.com/stretchr/testify/assert"
)

func TestNewMapping(t *testing.T) {
	m := colormapping.NewColorMapping()
	err := m.LoadDefaults()
	assert.NoError(t, err)

	assert.NotNil(t, m.GetColors())

	c := m.GetColor("scifi_nodes:blacktile2", 0)
	assert.NotNil(t, c)
	assert.Equal(t, uint8(20), c.R)
	assert.Equal(t, uint8(20), c.G)
	assert.Equal(t, uint8(20), c.B)
	assert.Equal(t, uint8(255), c.A)

	c = m.GetColor("homedecor:piano", 0)
	assert.NotNil(t, c)
	assert.Equal(t, uint8(61), c.R)
	assert.Equal(t, uint8(61), c.G)
	assert.Equal(t, uint8(61), c.B)
	assert.Equal(t, uint8(255), c.A)

	c = m.GetColor("default:river_water_flowing", 0)
	assert.Equal(t, uint8(39), c.R)
	assert.Equal(t, uint8(90), c.G)
	assert.Equal(t, uint8(106), c.B)
	assert.Equal(t, uint8(128), c.A)
	assert.NotNil(t, c)

	c = m.GetColor("unifiedbricks:brickblock_multicolor_dark", 100)
	assert.NotNil(t, c)

	c = m.GetColor("mymod:my_red_node", 0)
	assert.NotNil(t, c)
	assert.Equal(t, uint8(107), c.R)
	assert.Equal(t, uint8(54), c.G)
	assert.Equal(t, uint8(53), c.B)

	c = m.GetColor("mymod:my_dark_red_node", 0)
	assert.NotNil(t, c)
	assert.Equal(t, uint8(107-30), c.R)
	assert.Equal(t, uint8(54-30), c.G)
	assert.Equal(t, uint8(53-30), c.B)

}

func TestNewMappingErrors(t *testing.T) {
	m := colormapping.NewColorMapping()
	count, err := m.LoadBytes([]byte("stuff"))
	assert.Error(t, err)
	assert.Equal(t, 0, count)

	count, err = m.LoadBytes([]byte("my:node invalid_r 0 0"))
	assert.Error(t, err)
	assert.Equal(t, 0, count)

	count, err = m.LoadBytes([]byte("my:node 0 invalid_g 0"))
	assert.Error(t, err)
	assert.Equal(t, 0, count)

	count, err = m.LoadBytes([]byte("my:node 0 0 invalid_b"))
	assert.Error(t, err)
	assert.Equal(t, 0, count)
}

func TestLoadErrors(t *testing.T) {
	m := colormapping.NewColorMapping()
	_, err := m.LoadVFSColors("bogus.txt")
	assert.Error(t, err)
}
