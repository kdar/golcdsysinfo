package main

import (
  "fmt"
  "os"
  "time"

  lsi "golcdsysinfo"
)

func rectDemo(l *lsi.LCDSysInfo) {
  l.Clear(lsi.LINE_ALL, lsi.BG_COLOR_BLACK)
  time.Sleep(500 * time.Millisecond)

  l.DrawRect(lsi.Rect{X1: 10, Y1: 10, X2: 200, Y2: 40, Fill: true, Color: lsi.BG_COLOR_YELLOW})
  l.DrawRect(lsi.Rect{X1: 10, Y1: 40, X2: 80, Y2: 70, Fill: true, Color: lsi.BG_COLOR_RED})
  l.DrawRect(lsi.Rect{X1: 100, Y1: 100, X2: 200, Y2: 130, Fill: true, Color: lsi.BG_COLOR_GREEN})
  l.DrawRect(lsi.Rect{X1: 200, Y1: 210, X2: 250, Y2: 220, Fill: true, Color: lsi.BG_COLOR_BLUE})

  l.DrawRect(lsi.Rect{X1: 20, Y1: 200, X2: 50, Y2: 230, LineWidth: 3, Color: lsi.BG_COLOR_YELLOW})
  l.DrawRect(lsi.Rect{X1: 150, Y1: 200, X2: 300, Y2: 239, LineWidth: 10, Color: lsi.BG_COLOR_PURPLE})
  l.DrawRect(lsi.Rect{X1: 12, Y1: 12, X2: 198, Y2: 38, LineWidth: 2, Color: lsi.BG_COLOR_ORANGE})
}

func textDemo(l *lsi.LCDSysInfo) {
  l.Clear(lsi.LINE_ALL, lsi.BG_COLOR_BLACK)
  time.Sleep(500 * time.Millisecond)

  l.Clear(lsi.LINE_2, lsi.BG_COLOR_RED)
  time.Sleep(50 * time.Millisecond)

  l.DrawIconOnGrid(0, 14)
  l.SetTextBackgroundColor(lsi.BG_COLOR_BLACK)
  l.DrawTextOnLine(1, "|LCDsysinfo", true, lsi.TEXT_ALIGN_LEFT, lsi.FG_COLOR_YELLOW)
  l.DrawIconOnGrid(6, 8)
  l.SetTextBackgroundColor(lsi.BG_COLOR_RED)
  l.DrawTextOnLine(2, "Enter your message", false, lsi.TEXT_ALIGN_LEFT, lsi.FG_COLOR_LIGHT_GREEN)
  l.SetTextBackgroundColor(lsi.BG_COLOR_BLACK)
}

func iconsTextAnywhereDemo(l *lsi.LCDSysInfo) {
  l.Clear(lsi.LINE_ALL, lsi.BG_COLOR_BLACK)
  time.Sleep(500 * time.Millisecond)

  l.DrawIconAnywhere(0, 0, 15)
  time.Sleep(60 * time.Millisecond)
  l.DrawIconAnywhere(12, 20, 27)
  time.Sleep(60 * time.Millisecond)
  l.DrawIconAnywhere(20, 120, 16)
  time.Sleep(60 * time.Millisecond)
  l.DrawIconAnywhere(247, 199, 17)
  time.Sleep(60 * time.Millisecond)
  l.DrawIconAnywhere(120, 100, 18)
  time.Sleep(60 * time.Millisecond)
  l.DrawIconAnywhere(279, 56, 20)
  time.Sleep(60 * time.Millisecond)
  l.DrawIconAnywhere(50, 156, 8)
  time.Sleep(60 * time.Millisecond)
  l.SetTextBackgroundColor(lsi.BG_COLOR_RED)
  l.DrawTextAnywhere(75, 20, "1234567890!", lsi.FG_COLOR_YELLOW)
  time.Sleep(60 * time.Millisecond)
  l.DrawTextAnywhere(75, 120, "abcdefghijk!", lsi.FG_COLOR_YELLOWGREEN)
  l.SetTextBackgroundColor(lsi.BG_COLOR_BLACK)
}

func linesDemo(l *lsi.LCDSysInfo) {
  l.Clear(lsi.LINE_ALL, lsi.BG_COLOR_BLACK)
  time.Sleep(500 * time.Millisecond)

  l.DrawLine(lsi.Line{Y1: 22, X1: 10, Y2: 100, X2: 150, Color: lsi.BG_COLOR_BLUE})
  l.DrawLine(lsi.Line{Y1: 100, X1: 150, Y2: 200, X2: 10, Color: lsi.BG_COLOR_GREEN})
  l.DrawLine(lsi.Line{Y1: 35, X1: 100, Y2: 50, X2: 233, Color: lsi.BG_COLOR_YELLOW})
  l.DrawLine(lsi.Line{Y1: 152, X1: 2, Y2: 3, X2: 50, Color: lsi.BG_COLOR_RED})
}

func progressBarsDemo(l *lsi.LCDSysInfo) {
  l.Clear(lsi.LINE_ALL, lsi.BG_COLOR_BLACK)
  time.Sleep(500 * time.Millisecond)

  for count := 0; count < 101; count++ {
    l.DrawProgressBar(lsi.ProgressBar{
      Y1:                  30,
      X1:                  30,
      Y2:                  200,
      X2:                  40,
      Percent:             count,
      GradientTopColor:    lsi.BG_COLOR_YELLOW,
      GradientBottomColor: lsi.BG_COLOR_BLACK,
      BarColor:            lsi.Color24To16(0xFBC8F8F),
    })
    time.Sleep(5 * time.Millisecond)
  }
  for count := 100; count > 0; count-- {
    l.DrawProgressBar(lsi.ProgressBar{
      Y1:                  30,
      X1:                  30,
      Y2:                  200,
      X2:                  40,
      Percent:             count,
      GradientTopColor:    lsi.BG_COLOR_YELLOW,
      GradientBottomColor: lsi.BG_COLOR_BLACK,
      BarColor:            lsi.Color24To16(0xFBC8F8F),
    })
    time.Sleep(5 * time.Millisecond)
  }

  for count := 0; count < 101; count++ {
    l.DrawProgressBar(lsi.ProgressBar{
      Y1:                  60,
      X1:                  60,
      Y2:                  230,
      X2:                  70,
      Percent:             count,
      GradientTopColor:    lsi.BG_COLOR_BLUE,
      GradientBottomColor: lsi.BG_COLOR_BLACK,
      BarColor:            lsi.BG_COLOR_PURPLE,
    })
    time.Sleep(5 * time.Millisecond)
  }
  for count := 100; count > 0; count-- {
    l.DrawProgressBar(lsi.ProgressBar{
      Y1:                  60,
      X1:                  60,
      Y2:                  230,
      X2:                  70,
      Percent:             count,
      GradientTopColor:    lsi.BG_COLOR_BLUE,
      GradientBottomColor: lsi.BG_COLOR_BLACK,
      BarColor:            lsi.BG_COLOR_PURPLE,
    })
    time.Sleep(5 * time.Millisecond)
  }
}

func main() {
  l := &lsi.LCDSysInfo{}
  err := l.Open(0)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  defer l.Close()

  l.DimOnIdle(false)

  l.SetBrightness(155)

  // l.Clear(lsi.LINE_ALL, lsi.BG_COLOR_BLACK)
  // time.Sleep(500 * time.Millisecond)
  // l.DrawIconOnGrid(2*8+7, 15)

  fmt.Println("Draw Rectangles")
  rectDemo(l)
  time.Sleep(1 * time.Second)
  fmt.Println("Draw Text")
  textDemo(l)
  time.Sleep(1 * time.Second)
  fmt.Println("Draw text/icons anywhere")
  iconsTextAnywhereDemo(l)
  time.Sleep(1 * time.Second)
  fmt.Println("Draw lines")
  linesDemo(l)
  time.Sleep(1 * time.Second)
  fmt.Println("Draw progress bars")
  progressBarsDemo(l)
  time.Sleep(1 * time.Second)
}
