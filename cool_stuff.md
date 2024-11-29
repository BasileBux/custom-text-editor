# Cool stuff

This file contains all the stuff I have tried at one point and that could be cool to have at least as an option

It could also be some nice starting points for features ideas. Or some complete garbage (but fun).

## Transparent background

Transparent window backgorund. Maybe something could be made with custom effects like blur idk. 

```go
// Put this flag before window creation
rl.SetConfigFlags(rl.FlagWindowTransparent)

// Put this as reset in the main loop
transparent := rl.Color{
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		}
rl.ClearBackground(transparent)
```

## Shaders 

Needs 1 or 2 .fs glsl shader

```go
bloom := rl.LoadShader("", "shaders/bloom.fs")
rl.BeginShaderMode(bloom)
rl.EndShaderMode()
```
