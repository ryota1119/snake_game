package snake

// 盤面サイズや初期設定
const (
    BoardWidth  = 30
    BoardHeight = 20
    InitialLen  = 4
    TickMSec    = 120
)

// Point は座標を表す
type Point struct {
    X, Y int
}

// Direction は方向を表す
type Direction int

// Direction の定義
const (
    Up Direction = iota
    Down
    Left
    Right
)
