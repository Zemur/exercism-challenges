=begin
Write your code for the 'Diamond' exercise in this file. Make the tests in
`diamond_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/diamond` directory.
=end
class Diamond
  def self.make_diamond(letter)
    # Create diamond array
    diamond = []

    # Force letter to uppercase.
    letter.upcase!

    # Generate an array containing letters 'A' to the specified letter.
    letters = ('A'..letter).to_a

    # Set the width of the diamond.
    width = if letters.length > 1
              # Double the width and take away one to get the width of the full diamond.
              letters.length * 2 - 1
            else
              # If there is only one letter, just set width to 1.
              1
            end

    # The initial placement will be halfway through the width
    replacement_index = (width / 2)

    # Build half the diamond.
    letters.each_index do |i|
      # Create a line with the appropriate amount of spaces.
      line = ' '*width

      # Choose where to place the letters on the line.
      # If it is not the first iteration, use the replacement index specified above (halfway through the line).
      # Otherwise, place the letter at the specified replacement index.
      if i > 0
        [replacement_index-i, replacement_index+i].each { |rs| line[rs] = letters[i] }
      else
        line[replacement_index] = letters[i]
      end

      # Add line to build half the diamond.
      diamond.append("#{line}\n")
    end

    # Take half the diamond, except for the last element, flip it, merge it, and return as a string.
    diamond.concat(diamond[0...-1].reverse).join
  end
end