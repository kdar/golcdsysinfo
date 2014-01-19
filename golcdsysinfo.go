package golcdsysinfo

import (
  "fmt"
  "github.com/kylelemons/gousb/usb"
)

const (
  VENDOR_ID  = 0x16c0
  PRODUCT_ID = 0x05dc
)

// BgColor represents a 16 bit color. 5 bits for red,
// 6 bits for green, and 5 bits for blue.
type BgColor uint16

const (
  BG_COLOR_YELLOW BgColor = 0xFFE0
  BG_COLOR_WHITE          = 0xFFFF
  BG_COLOR_BLACK          = 0x0000
  BG_COLOR_BLUE           = 0x001F
  BG_COLOR_RED            = 0xF800
  BG_COLOR_GREEN          = 0x07E0
  BG_COLOR_ORANGE         = 0xFD20
  BG_COLOR_PURPLE         = 0xF81F
)

// TextAlign determines how a drawn text is aligned.
type TextAlign int

const (
  TEXT_ALIGN_NONE   TextAlign = -1
  TEXT_ALIGN_CENTER           = 0
  TEXT_ALIGN_LEFT             = 1
  TEXT_ALIGN_RIGHT            = 2
)

// FgColor represents colors built into the LCDSysInfo
// device. They usually use it to represent foreground
// colors.
type FgColor uint16

const (
  FG_COLOR_GREEN       FgColor = 1
  FG_COLOR_YELLOW              = 2
  FG_COLOR_RED                 = 3
  FG_COLOR_WHITE               = 5
  FG_COLOR_CYAN                = 6
  FG_COLOR_GREY                = 7
  FG_COLOR_BLACK               = 13
  FG_COLOR_BROWN               = 15
  FG_COLOR_BRICK_RED           = 16
  FG_COLOR_DARK_BLUE           = 17
  FG_COLOR_LIGHT_BLUE          = 18
  FG_COLOR_ORANGE              = 21
  FG_COLOR_PURPLE              = 22
  FG_COLOR_PINK                = 23
  FG_COLOR_PEACH               = 24
  FG_COLOR_GOLD                = 25
  FG_COLOR_LAVENDER            = 26
  FG_COLOR_ORANGE_RED          = 27
  FG_COLOR_MAGENTA             = 28
  FG_COLOR_NAVY                = 30
  FG_COLOR_LIGHT_GREEN         = 31
  FG_COLOR_YELLOWGREEN         = 31
)

// TextLine represents 1-6 lines on the device. This
// is only used for l.Clear().
type TextLine int

const (
  LINE_1 TextLine = 1
  LINE_2          = 1 << iota
  LINE_3
  LINE_4
  LINE_5
  LINE_6
  LINE_ALL = LINE_1 | LINE_2 | LINE_3 | LINE_4 | LINE_5 | LINE_6
)

var (
  fontLengthTable = []byte{
    0x11, 0x06, 0x08, 0x15, 0x0E, 0x19, 0x15, 0x03, 0x08, 0x08, 0x0F, 0x0D,
    0x05, 0x08, 0x06, 0x0B, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11,
    0x11, 0x11, 0x06, 0x06, 0x13, 0x10, 0x13, 0x0C, 0x1A, 0x14, 0x10, 0x12,
    0x13, 0x0F, 0x0D, 0x13, 0x11, 0x04, 0x07, 0x11, 0x0E, 0x14, 0x11, 0x15,
    0x0F, 0x15, 0x12, 0x10, 0x13, 0x11, 0x14, 0x1C, 0x13, 0x13, 0x12, 0x07,
    0x0B, 0x07, 0x0B, 0x02, 0x08, 0x0E, 0x0F, 0x0E, 0x0F, 0x10, 0x0B, 0x0F,
    0x0E, 0x04, 0x07, 0x0F, 0x04, 0x18, 0x0E, 0x10, 0x0F, 0x0F, 0x0A, 0x0D,
    0x0B, 0x0E, 0x10, 0x16, 0x10, 0x10, 0x0E, 0x01, 0x11, 0x02,
  }
)

// LCDSysInfo holds information about a LCDSysInfo device and allows you
// to draw various things onto the device.
type LCDSysInfo struct {
  ctx *usb.Context
  dev *usb.Device
}

// Open an LCDSysInfo device at index. If you have one device connected,
// then choose index 0. If you have more than one device connected, you will
// have to play around with index until you get the device you want.
func (l *LCDSysInfo) Open(index int) error {
  l.ctx = usb.NewContext()
  l.ctx.Debug(0)

  // Find our LCDSysInfo device.
  devs, err := l.ctx.ListDevices(func(desc *usb.Descriptor) bool {
    if desc.Vendor == VENDOR_ID && desc.Product == PRODUCT_ID {
      return true
    }

    return false
  })
  // I don't test for an error because I've noticed it returns a
  // device and throws a NOT_FOUND error. Ok?...
  if len(devs) == 0 {
    return fmt.Errorf("Could not find LCDSysInfo device: %v", err)
  }

  // Pick the nth index of the device. If we only have one device
  // connected, this will always be 0.
  if len(devs) > index {
    l.dev = devs[index]
  }

  // gousb will open all the devices it returns. So we must close
  // devices we are not using.
  for i, v := range devs {
    if i == index {
      continue
    }

    v.Close()
  }

  return nil
}

// Close the device. This MUST be called.
func (l *LCDSysInfo) Close() {
  if l.ctx != nil {
    l.ctx.Close()
  }
  if l.dev != nil {
    l.dev.Close()
  }
}

// Return the usb device.
func (l *LCDSysInfo) GetDevice() *usb.Device {
  return l.dev
}

// Set the brightness level. 0-255
func (l *LCDSysInfo) SetBrightness(value int) error {
  if value > 255 {
    value = 255
  }

  _, err := l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 13, uint16(value), uint16(value), []byte{})
  return err
}

// Tell the device whether to dim when it's idle or not.
func (l *LCDSysInfo) DimOnIdle(enable bool) (err error) {
  if enable {
    _, err = l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 17, 0, uint16(266), []byte{})
  } else {
    _, err = l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 17, uint16(1), 0, []byte{})
  }
  return
}

// Clear lines on the device. There are a total of 6 lines on the device.
// Example:
//   l.Clear(LINE_2|LINE_3, BG_COLOR_BLACK)
func (l *LCDSysInfo) Clear(lines TextLine, color BgColor) error {
  _, err := l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 26, uint16(lines), uint16(color), []byte{})
  return err
}

// Draw a builtin icon on the grid. There are 48 total positions on the grid. Given that
// there are 6 lines, that gives us 8 positions per line. So if you wanted to display
// an icon in the first spot of row 5, you would do:
//   l.DrawIconOnGrid(4*8+0, 15)
// If you wanted to draw an icon on the last spot of row 3, you would do:
//   l.DrawIconOnGrid(2*8+7, 15)
func (l *LCDSysInfo) DrawIconOnGrid(position, icon int) error {
  value := uint16(position*512 + icon)
  _, err := l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 27, value, uint16(25600), []byte{})
  return err
}

// Draw a builtin icon anywhere.
func (l *LCDSysInfo) DrawIconAnywhere(x, y, icon int) error {
  value := uint16(icon*256 + icon)
  data := []byte{
    byte(y / 256),
    byte(y % 256),
    byte(x / 256),
    byte(x % 256),
  }
  _, err := l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 29, value, value, data)
  return err
}

// Set the background color of any text that is drawn after this call.
func (l *LCDSysInfo) SetTextBackgroundColor(color BgColor) error {
  _, err := l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 30, uint16(color), 0, []byte{})
  return err
}

// Rect represents a rectangle to be drawn.
type Rect struct {
  X1, Y1, X2, Y2 int
  Fill           bool
  Color          BgColor
  LineWidth      uint16
}

// Draw a rectangle.
func (l *LCDSysInfo) DrawRect(r Rect) error {
  var data []byte
  filln := 2

  if r.Fill {
    data = []byte{
      byte(r.Y1 / 256),
      byte(r.Y1 % 256),
      byte(r.X1 / 256),
      byte(r.X1 % 256),
      byte(r.Y2 / 256),
      byte(r.Y2 % 256),
      byte(r.X2 / 256),
      byte(r.X2 % 256),
      byte(0 / 256),
      byte(r.LineWidth % 256),
    }
    filln = 1
  } else {
    data = []byte{
      byte(r.X1 / 256),
      byte(r.X1 % 256),
      byte(r.Y1 / 256),
      byte(r.Y1 % 256),
      byte(r.X2 / 256),
      byte(r.X2 % 256),
      byte(r.Y2 / 256),
      byte(r.Y2 % 256),
      byte(0 / 256),
      byte(r.LineWidth % 256),
    }
  }

  n, err := l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 95, uint16(filln), uint16(r.Color), data)
  if err == nil && n != len(data) {
    return fmt.Errorf("FillRectangle: tried to send %d bytes but only %d transferred", len(data), n)
  }

  return err
}

// Line represents a line to be drawn.
type Line struct {
  X1, Y1, X2, Y2 int
  Color          BgColor
}

// Draw a line.
func (l *LCDSysInfo) DrawLine(line Line) error {
  data := []byte{
    byte(line.Y1 / 256),
    byte(line.Y1 % 256),
    byte(line.X1 / 256),
    byte(line.X1 % 256),
    byte(line.Y2 / 256),
    byte(line.Y2 % 256),
    byte(line.X2 / 256),
    byte(line.X2 % 256),
    byte(0 / 256),
    byte(0 % 256),
  }

  n, err := l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 95, uint16(3), uint16(line.Color), data)
  if err == nil && n != len(data) {
    return fmt.Errorf("DrawLine: tried to send %d bytes but only %d transferred", len(data), n)
  }

  return err
}

// ProgressBar represents a progress bar to be drawn.
type ProgressBar struct {
  X1, Y1, X2, Y2      int
  Percent             int
  GradientTopColor    BgColor
  BarColor            BgColor
  GradientBottomColor BgColor
}

// Draw a progress bar.
func (l *LCDSysInfo) DrawProgressBar(bar ProgressBar) error {
  data := []byte{
    byte(bar.Y1 / 256),
    byte(bar.Y1 % 256),
    byte(bar.X1 / 256),
    byte(bar.X1 % 256),
    byte(bar.Y2 / 256),
    byte(bar.Y2 % 256),
    byte(bar.X2 / 256),
    byte(bar.X2 % 256),
    byte(0 / 256),
    byte(bar.Percent % 256),
    byte(bar.GradientTopColor / 256),
    byte(bar.GradientTopColor % 256),
    byte(bar.BarColor / 256),
    byte(bar.BarColor % 256),
  }

  n, err := l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 95, uint16(4), uint16(bar.GradientBottomColor), data)
  if err == nil && n != len(data) {
    return fmt.Errorf("DrawProgressBar: tried to send %d bytes but only %d transferred", len(data), n)
  }

  return err
}

// Draw text on a particular line (1-6).
func (l *LCDSysInfo) DrawTextOnLine(line int, text string, padForIcon bool, align TextAlign, color FgColor) error {
  data := convertText(text, padForIcon, align)
  dataLen := len(data)
  if !padForIcon {
    dataLen += 256
  }
  if color > 32 {
    color = 0
  }
  if line < 1 || line > 6 {
    line = 1
  }

  idx := uint16(((line - 1) * 256) + int(color))

  n, err := l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 24, uint16(dataLen), idx, []byte(data))
  if err == nil && n != len(data) {
    return fmt.Errorf("DrawTextOnLine: tried to send %d bytes but only %d transferred", len(data), n)
  }

  return err
}

// Draw text anywhere
func (l *LCDSysInfo) DrawTextAnywhere(x, y int, text string, color FgColor) error {
  data := []byte{
    byte(x / 256),
    byte(x % 256),
    byte(y / 256),
    byte(y % 256),
    byte(319 / 256),
    byte(319 % 256),
    byte((y + 40) / 256),
    byte((y + 40) % 256),
  }

  if color > 32 {
    color = 0
  }

  data = append(data, text...)
  n, err := l.dev.Control(uint8(usb.ENDPOINT_DIR_OUT|usb.REQUEST_TYPE_VENDOR), 25, uint16(len(text)), uint16(color), data)
  if err == nil && n != len(data) {
    return fmt.Errorf("DrawTextAnywhere: tried to send %d bytes but only %d transferred", len(data), n)
  }

  return err
}
