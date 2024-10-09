package main

import "fmt"

type Piece struct {
    Color    rune  // 'W' for White, 'B' for Black
    Wall     bool  // Indicates if the piece is a wall (standing stone)
    Capstone bool  // Indicates if the piece is a capstone
}

type Board struct {
    Locations map[string][]*Piece
}

// Initialize an empty board with locations (3x3 to start with)
func initializeBoard() Board {
    return Board{
        Locations: map[string][]*Piece{
            "a1": {}, "a2": {}, "a3": {},
            "b1": {}, "b2": {}, "b3": {},
            "c1": {}, "c2": {}, "c3": {},
        },
    }
}

func printBoard(b Board) {
    for loc, stack := range b.Locations {
        fmt.Printf("%s: ", loc)

        if len(stack) == 0 {
            fmt.Print(". ")
        } else {
            for _, piece := range stack {
                fmt.Printf("%c", piece.Color) // Print piece colors in the stack
            }
        }

        fmt.Println()
    }

    fmt.Println()
}

// Function to move pieces between locations
// Todo: could add a direction (+ - < >) to follow notation instead of a "to" parameter
func movePiece(b *Board, from string, to string, numPieces int) {
    if len(b.Locations[from]) < numPieces {
        fmt.Println("Not enough pieces to move")
        return
    }

    // Move the top 'numPieces' pieces
    piecesToMove := b.Locations[from][len(b.Locations[from])-numPieces:]
    b.Locations[to] = append(b.Locations[to], piecesToMove...) // Add to destination
    b.Locations[from] = b.Locations[from][:len(b.Locations[from])-numPieces] // Remove from source
}

func main() {
    board := initializeBoard()

    // Example: Adding pieces to the board
    board.Locations["a2"] = append(board.Locations["a2"], &Piece{Color: 'W'})
    board.Locations["a2"] = append(board.Locations["a2"], &Piece{Color: 'B'})
    board.Locations["a2"] = append(board.Locations["a2"], &Piece{Color: 'B', Capstone: true})

    board.Locations["c3"] = append(board.Locations["c3"], &Piece{Color: 'W'})
    board.Locations["c3"] = append(board.Locations["c3"], &Piece{Color: 'B'})

    printBoard(board)

    // Example: Moving pieces from a2 to a3
    movePiece(&board, "a2", "a3", 2) // Move top 2 pieces from a2 to a3
    printBoard(board)
}
