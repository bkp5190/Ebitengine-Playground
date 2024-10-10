// Specify the 'pixel' mode.
//kage:unit pixels

package main

// Uniform variables.
var Time float
var Cursor vec2

// Fragment is the entry point of the fragment shader.
// Fragment returns the color value for the current position.
func Fragment(dstPos vec4, srcPos vec2, color vec4) vec4 {
	// You can define variables with a short variable declaration like Go.
	pos := dstPos.xy - imageDstOrigin()

	lightpos := vec3(Cursor, 50)
	lightdir := normalize(lightpos - vec3(pos, 0))
	normal := normalize(imageSrc1UnsafeAt(srcPos) - 0.5)
	const ambient = 0.25
	diffuse := 0.75 * max(0.0, dot(normal.xyz, lightdir))

	// You can treat multiple source images by
	// imageSrc[N]At or imageSrc[N]UnsafeAt.
	return imageSrc0UnsafeAt(srcPos) * (ambient + diffuse)
}
