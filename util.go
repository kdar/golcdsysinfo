package golcdsysinfo

import (
  "strings"
)

func Color24To16(color int) BgColor {
  r := uint16((color & 0xFF0000) >> 19) // 5 bits
  g := uint16((color & 0x00FF00) >> 10) // 6 bits
  b := uint16((color & 0x0000FF) >> 3)  // 5 bits
  return BgColor((r << 11) | (g << 5) | (b))
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

func ternaryInt(truth bool, a, b int) int {
  if truth {
    return a
  }
  return b
}

func alignText(text string, align TextAlign, screenPx, stringLengthPx int) string {
  spaces := (screenPx - stringLengthPx) / 17
  pixels := (screenPx - stringLengthPx) % 17
  sliceSize := max(spaces, pixels)
  before := make([]byte, 0, sliceSize)
  after := make([]byte, 0, sliceSize)

  switch align {
  case TEXT_ALIGN_CENTER:
    for x := 0; x < pixels; x++ {
      if x%2 == 0 {
        before = append(before, '{')
      } else {
        after = append(after, '{')
      }
    }
    for x := 0; x < spaces; x++ {
      if x%2 == 0 {
        after = append(after, ' ')
      } else {
        before = append(before, ' ')
      }
    }
  case TEXT_ALIGN_LEFT:
    for x := 0; x < spaces; x++ {
      after = append(after, ' ')
    }
    for x := 0; x < pixels; x++ {
      after = append(after, '{')
    }
  case TEXT_ALIGN_RIGHT:
    for x := 0; x < spaces; x++ {
      before = append(before, ' ')
    }
    for x := 0; x < pixels; x++ {
      before = append(before, '{')
    }
  }

  return string(before) + text + string(after) + "\x00"
}

func convertText(text string, padForIcon bool, align TextAlign) string {
  screenPx := 0
  if padForIcon {
    screenPx = (40 * 7) + 1
  } else {
    screenPx = (40 * 8) - 1
  }
  text = strings.TrimRight(text, " ")
  text = strings.Replace(text, " ", "_", -1)
  stringLengthPx := 0

  for i, v := range text {
    asciiValue := int(v)
    charLengthPx := 0
    if asciiValue >= 32 && asciiValue <= 125 {
      charLengthPx = int(fontLengthTable[asciiValue-32])
    }
    if stringLengthPx+charLengthPx > screenPx {
      text = text[0:i]
      break
    }
    stringLengthPx += charLengthPx
  }

  return alignText(text, align, screenPx, stringLengthPx)
}

func TextLength(text string) int {
  px := 0
  for _, v := range text {
    asciiValue := int(v)
    if asciiValue >= 32 && asciiValue <= 125 {
      px += int(fontLengthTable[asciiValue-32])
    }
  }

  return px
}
