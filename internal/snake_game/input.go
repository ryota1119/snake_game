package snake

import (
    "log"

    "github.com/gdamore/tcell/v2"
)

// RunInputLoop はキー入力を受け取って inputCh に流し込むゴルーチン
func (g *Game) RunInputLoop() {
    for {
        ev := g.scr.PollEvent()
        if ev == nil {
            continue
        }
        if key, ok := ev.(*tcell.EventKey); ok {
            switch key.Key() {
            case tcell.KeyUp:
                g.trySendDir(Up)
            case tcell.KeyDown:
                g.trySendDir(Down)
            case tcell.KeyLeft:
                g.trySendDir(Left)
            case tcell.KeyRight:
                g.trySendDir(Right)
            case tcell.KeyRune:
                if key.Rune() == 'q' {
                    g.scr.Fini()
                    log.Fatal("Quit")
                }
            }
        }
    }
}

// trySendDir は方向キーが押されたときに inputCh に流し込む
func (g *Game) trySendDir(d Direction) {
    // 逆走を防止
    if (g.Dir == Up && d == Down) ||
        (g.Dir == Down && d == Up) ||
        (g.Dir == Left && d == Right) ||
        (g.Dir == Right && d == Left) {
        return
    }
    select {
    case g.inputCh <- d:
    default:
    }
}
