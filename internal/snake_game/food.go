package snake

import "math/rand"

// RandomPoint は board 内ランダムな Point を生成し、
// forbid スライス内の点と重複しないようにする
func RandomPoint(forbid []Point) Point {
out:
    for {
        p := Point{
            X: rand.Intn(BoardWidth),
            Y: rand.Intn(BoardHeight),
        }
        for _, f := range forbid {
            if p == f {
                continue out
            }
        }
        return p
    }
}
