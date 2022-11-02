func minTimeToType(word string) int {
    current_point := 'a'   // ascii a = 97
    type_count := len(word)
    move_count := 0
    
    for _, ch := range word {
        clockwise_move := math.Abs(float64(current_point - ch))
        counterclockwise_move := 26 - clockwise_move
        current_point = ch
        move_count += int(math.Min(clockwise_move, counterclockwise_move))
    }
    
    
    return type_count + move_count
}
