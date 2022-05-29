  def max_product(words)
    words = words.sort_by(&:length).reverse

    # 1. Pre-process each word into binary forms
    bins = words.map do |word|
      word.chars.reduce(0){|val, ch| val |= 1 << (ch.ord - 'a'.ord) }
    end

    max = 0
    words.each_with_index do |a, i|
      # 2. It'll be TLE without this pruning
      next if a.length * a.length <= max

      words[i+1..-1].each_with_index do |b, j|
        next if (bins[i] & bins[i+1+j]).nonzero?

        max = [ max, a.length * b.length ].max
      end
    end

    max
  end