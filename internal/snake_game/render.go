package snake

import (
    "fmt"

    "github.com/gdamore/tcell/v2"
)

// Draw は枠・フード・スネーク・スコアを描画する
func (g *Game) Draw() {
    g.scr.Clear()
    styleSnake := tcell.StyleDefault.Foreground(tcell.ColorGreen)
    styleFood := tcell.StyleDefault.Foreground(tcell.ColorRed)
    styleBorder := tcell.StyleDefault.Foreground(tcell.ColorGray)

    // 枠
    for x := 0; x < BoardWidth+2; x++ {
        g.scr.SetContent(x, 0, '■', nil, styleBorder)
        g.scr.SetContent(x, BoardHeight+1, '■', nil, styleBorder)
    }
    for y := 0; y < BoardHeight+2; y++ {
        g.scr.SetContent(0, y, '■', nil, styleBorder)
        g.scr.SetContent(BoardWidth+1, y, '■', nil, styleBorder)
    }

    // フード
    g.scr.SetContent(g.Food.X+1, g.Food.Y+1, '●', nil, styleFood)

    // スネーク
    for i, p := range g.Snake {
        ch := '○'
        if i == 0 {
            ch = '◎'
        }
        g.scr.SetContent(p.X+1, p.Y+1, ch, nil, styleSnake)
    }

    // スコア
    for i, r := range fmt.Sprintf("Score: %d", g.Score) {
        g.scr.SetContent(BoardWidth+4+i, 1, r, nil, styleBorder)
    }

    g.scr.Show()
}

// DrawGameOver は GAME OVER メッセージを中央に表示する
func (g *Game) DrawGameOver() {
    msg := " GAME OVER "
    style := tcell.StyleDefault.Foreground(tcell.ColorWhite).
        Background(tcell.ColorMaroon).Bold(true)
    y := BoardHeight/2 + 1
    x := (BoardWidth+2-len(msg))/2 + 1
    for i, r := range msg {
        g.scr.SetContent(x+i, y, r, nil, style)
    }
    g.scr.Show()
}
