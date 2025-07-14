package snake

import (
    "time"

    "github.com/gdamore/tcell/v2"
)

// Game はゲームの状態を保持する
type Game struct {
    Snake      []Point
    Dir        Direction
    Food       Point
    Score      int
    GameOver   bool
    scr        tcell.Screen
    inputCh    chan Direction
    updateDone chan struct{}
}

// NewGame は新規ゲーム状態を初期化して返す
func NewGame(scr tcell.Screen) *Game {
    snakeBody := make([]Point, InitialLen)
    for i := 0; i < InitialLen; i++ {
        snakeBody[i] = Point{X: BoardWidth/2 + i, Y: BoardHeight / 2}
    }
    return &Game{
        Snake:      snakeBody,
        Dir:        Left,
        Food:       RandomPoint(snakeBody),
        scr:        scr,
        inputCh:    make(chan Direction, 1),
        updateDone: make(chan struct{}),
    }
}

// Loop はメインループ。tick ごとに更新→描画し、ゲームオーバー時には return する
func (g *Game) Loop() {
    ticker := time.NewTicker(TickMSec * time.Millisecond)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            g.update()
            g.Draw()
        case <-g.updateDone:
            return
        }
    }
}

// update はゲームの状態を更新する
func (g *Game) update() {
    // 入力反映
    select {
    case d := <-g.inputCh:
        g.Dir = d
    default:
    }

    // ヘッド移動
    head := Move(g.Snake, g.Dir)

    // 壁衝突
    if head.X < 0 || head.X >= BoardWidth || head.Y < 0 || head.Y >= BoardHeight {
        g.end()
        return
    }
    // 自己衝突
    if CollideSelf(g.Snake, head) {
        g.end()
        return
    }

    // 伸長 or 移動
    g.Snake = append([]Point{head}, g.Snake...)
    if head == g.Food {
        g.Score++
        g.Food = RandomPoint(g.Snake)
    } else {
        g.Snake = g.Snake[:len(g.Snake)-1]
    }
}

// end はゲームオーバー時の処理を行う
func (g *Game) end() {
    g.GameOver = true
    g.DrawGameOver()
    time.Sleep(1200 * time.Millisecond)
    g.scr.Clear()
    g.scr.Show()
    // 新しい Game で上書きして再スタート
    *g = *NewGame(g.scr)
    close(g.updateDone)
}
