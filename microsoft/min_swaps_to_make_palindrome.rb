def min_swaps_to_make_palindrome(input)
  return -1 if input.length.zero? || !is_can_make_palindrome?(input)

  swap_count = 0
  start_index = 0
  end_index = input.length - 1

  while (start_index < end_index)
    if (input[start_index] == input[end_index])
      start_index += 1
      end_index -= 1
    else
      temp_end = end_index - 1

      while(temp_end > start_index && input[temp_end] != input[start_index] )
        temp_end -= 1
      end

      if temp_end == start_index
        input[start_index], input[start_index + 1] = input[start_index + 1], input[start_index]
        swap_count += 1
      else
        (end_index - temp_end).times do
          input[temp_end], input[temp_end + 1] = input[temp_end + 1], input[temp_end]
          swap_count += 1
          temp_end += 1
        end
      end
    end
  end

  swap_count
end

def is_can_make_palindrome?(str)
  h, count_odd = Hash.new(0), 0

  str.each_char { |c| h[c] += 1 }
  ('a'..'z').to_a.each do |c|
    count_odd += 1 if h[c] > 0 && h[c] % 2 == 1
  end

  count_odd <= 1
end

pp min_swaps_to_make_palindrome('mamad')
pp min_swaps_to_make_palindrome('asflkj')
pp min_swaps_to_make_palindrome('aabb')
pp min_swaps_to_make_palindrome('ntiin')
