package snake

// Move は body の先頭ヘッドを dir の向きに進めた新しい Point を返す
func Move(body []Point, dir Direction) Point {
    head := body[0]
    switch dir {
    case Up:
        head.Y--
    case Down:
        head.Y++
    case Left:
        head.X--
    case Right:
        head.X++
    }
    return head
}

// CollideSelf は head が body のいずれかと一致するかを判定する
func CollideSelf(body []Point, head Point) bool {
    for _, p := range body {
        if p == head {
            return true
        }
    }
    return false
}
