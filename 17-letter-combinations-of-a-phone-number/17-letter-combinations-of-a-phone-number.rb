def letter_combinations(digits)
  @result = []
  @digit_to_letters = {
    '2' => %w[a b c],
    '3' => %w[d e f],
    '4' => %w[g h i],
    '5' => %w[j k l],
    '6' => %w[m n o],
    '7' => %w[p q r s],
    '8' => %w[t u v],
    '9' => %w[w x y z],
  }
  combination(digits, "", 0) unless digits.empty?
  @result
end

def combination(digits, curr, index)
  @result << curr and return if curr.length == digits.length 
  @digit_to_letters[digits[index]].each {|c|
    combination(digits, curr+c, index+1)
  }
end