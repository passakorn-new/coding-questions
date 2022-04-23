require 'base64'

@h = {}

def encode(long_url)
  short_url = Base64.encode64(long_url)[0..5]
  @h[short_url] = long_url
  short_url
end

def decode(short_url)
  @h[short_url]
end