module QueenAttack

let create (x, y) = x >= 0 && x < 8 && y >= 0 && y < 8

// catAttack determines if two queens can attack each other
let canAttack queen1 queen2 =
    let (x1, y1) = queen1
    let (x2, y2) = queen2

    x1 = x2
    || y1 = y2
    || x2 - x1 = y2 - y1
    || x2 - x1 = y1 - y2
