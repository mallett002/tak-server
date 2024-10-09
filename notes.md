# Server side for TAK mobile

### Notation
move: 3a1+11
"move 3 pieces from a1 up and drop 2 pieces along the way"

ex visual:
        a1      a2      a3
before: |||  
after:          |       ||


4b2+112
5b5-221


### Board

```go
type Piece struct {
  Color: rune // 'W' | 'B'
  wall: bool //
  capstone: bool, // is capstone
}

type Board [][]*Piece  // Pointer to Piece for efficiency

// Initialize a board with pieces
// Todo: Think each location should be a slice to allow for pieces stacked on each other
func initializeBoard() Board {
    return Board{
        {nil, &Piece{Color: 'W', Capstone: true}, nil, &Piece{Color: 'B'}, &Piece{Color: 'B'}},
        {nil, nil, nil, &Piece{Color: 'B'}, nil},
        {nil, nil, nil, nil, nil},
        {&Piece{Color: 'W'}, nil, nil, &Piece{Color: 'B'}, nil},
        {nil, &Piece{Color: 'B'}, nil, nil, &Piece{Color: 'W'}},
    }
}

func printBoard(b Board) {
    for _, row := range b {
        for _, piece := range row {
            if piece == nil {
                fmt.Print(". ") // Empty space
            } else {
                fmt.Printf("%c ", piece.Color) // Print piece color
            }
        }
        fmt.Println()
    }
    fmt.Println()
}

func main() {
    board := initializeBoard()
    printBoard(board)
}

```

### Wondering if this is better for the board (I think so)
{ 
    "a1": [],
    "a2: [&Piece{Color: 'W'}, &Piece{Color: 'B'}, &Piece{Color: 'B', Capstone: true}] // stack with capstone on top
    "a3: [],
    "b1": [],
    ...
    "c3": [&Piece{Color: 'W'}, &Piece{Color: 'B'}] // black piece on top
}






