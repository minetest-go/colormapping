
Golang minetest colormappings

![](https://github.com/minetest-go/colormapping/workflows/test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/minetest-go/colormapping/badge.svg)](https://coveralls.io/github/minetest-go/colormapping)

# Example

```golang
m := colormapping.NewColorMapping()
err := m.LoadDefaults()
assert.NoError(t, err)

c := m.GetColor("scifi_nodes:blacktile2", 0)
assert.NotNil(t, c)
assert.Equal(t, uint8(20), c.R)
assert.Equal(t, uint8(20), c.G)
assert.Equal(t, uint8(20), c.B)
```

# License

Code: **MIT**