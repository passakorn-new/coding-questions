# @param {Integer[][]} board
# @return {Void} Do not return anything, modify board in-place instead.
def game_of_life(board)
    @rows = board.length
    @cols = board[0].length
    
    @temp_board = Marshal.load(Marshal.dump(board))
    @rows.times do |n|
        @cols.times do |m|
            live_neighbors = live_neighbors_cnt(n, m)
            if board[n][m] == 1
                board[n][m] = 0 if live_neighbors < 2 || live_neighbors > 3
            else
                board[n][m] = 1 if live_neighbors == 3
            end
        end
    end
    
    board
end

def live_neighbors_cnt(n, m)
    # check 8 neighbors
    cnt = 0
    
    #  vertical     
    cnt += 1 if n - 1 >= 0    && @temp_board[n - 1][m] == 1
    cnt += 1 if n + 1 < @rows && @temp_board[n + 1][m] == 1
    
    # horizontal
    cnt += 1 if m - 1 >= 0    && @temp_board[n][m - 1] == 1
    cnt += 1 if m + 1 < @cols && @temp_board[n][m + 1] == 1
    
    # diagonal
    cnt += 1 if m - 1 >= 0    && n - 1 >= 0    && @temp_board[n - 1][m - 1] == 1
    cnt += 1 if m + 1 < @cols && n + 1 < @rows && @temp_board[n + 1][m + 1] == 1
    cnt += 1 if m + 1 < @cols && n - 1 >= 0    && @temp_board[n - 1][m + 1] == 1
    cnt += 1 if m - 1 >= 0    && n + 1 < @rows && @temp_board[n + 1][m - 1] == 1
    cnt
end