<!-- PROJECT LOGO -->
<h1 align="center">PL2</h1>
<p align="center">
  Package for transcoding PL2 image files.
  <br />
  <br />
  <a href="https://github.com/gravestench/dcc/issues">Report Bug</a>
  ·
  <a href="https://github.com/gravestench/dcc/issues">Request Feature</a>
</p>

# About
This code is a derivation of Lectem's work, [found here](https://github.com/Lectem/Worldstone).

This repo contains a codec for PL2 palette transformation files, as well as some command line tools and shell scripts.

## PL2 - A Palette Transformation data structure
The PL2 file format was used in Blizzard's Diablo 2, and is a relic of the 8-bit gaming industry.

The PL2 data structure represents a color palette, 
and many pre-computed transformations into that palette. 
The color palette contains 256 colors, each of which has 8-bit RGB 
components and no alpha support. After the palette, there are 
several categories of "transformations" into the palette 
(eg. gamma, contrast, blend modes, hue shifts).  

After the palette and transforms, **there is another palette and set of transformations**. 
This second palette and set of transformations is intended to be used for bitmap fonts.

## Command-line tools

- `pl2-to-gpl` exports the base palette as a GIMP palette.
- `pl2-from-gpl` generates a PL2 and its transformations from a GIMP palette.
- `pl2-to-png` renders PL2 palette and transformation data as PNG.

Install all tools with `go install ./cmd/...`. See [`cmd/README.md`](cmd/README.md) for flags and examples.

### More about the data structure
In the following table, **a transform is 256 bytes**, each byte is an index that points to a color in the palette of the PL2.

| Name        | Length         |  Notes |
| ----------- | -------------- |  ----- |
| Palette     | `256 * 4` bytes | 0's between each R,G,B byte sequence |
| Dark-to-mid Light Levels   | 32 transforms | a gradient of transforms, each being lighter |
| Mid-to-bright Light Levels   | 16 transforms | a gradient of transforms, each being lighter |
| Alpha-blend | `3 * 256` transforms | 25%, 50%, 75% opacity blends against each other color in the palette |
| Add blend-mode | 256 transforms | the "add" blend against each color in the palette |
| Add multiply-mode | 256 transforms | the "multiply" blend against each color in the palette |
| Color variations | 111 transforms | various hue-shifting transformations |
| Red Tones | 1 transform | a transform of only red tones |
| Green Tones | 1 transform | same, but only green tones |
| Blue Tones | 1 transform | same, but only blue tones |
| Unknown(?) | 14 transforms |  |
| Max Component blend | 256 transforms |  |
| Darkened color shift | 1 transform |  |
| Text color palette | `13 * 3` bytes | this is another color palette, but for text |
| Text color shifts | 13 transforms | color-shifts used for text |
