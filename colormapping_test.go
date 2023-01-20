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

	c := m.GetColor("scifi_nodes:blacktile2", 0)
	assert.NotNil(t, c)
	assert.Equal(t, uint8(20), c.R)
	assert.Equal(t, uint8(20), c.G)
	assert.Equal(t, uint8(20), c.B)

	c = m.GetColor("default:river_water_flowing", 0)
	assert.NotNil(t, c)

	c = m.GetColor("unifiedbricks:brickblock_multicolor_dark", 100)
	assert.NotNil(t, c)

}
