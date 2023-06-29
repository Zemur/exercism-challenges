=begin
Write your code for the 'Resistor Color Duo' exercise in this file. Make the tests in
`resistor_color_duo_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/resistor-color-duo` directory.
=end

module ResistorColorDuo
  def self.value(band_colors)
    colors = ['Black', 'Brown', 'Red', 'Orange', 'Yellow', 'Green', 'Blue', 'Violet', 'Grey', 'White']
    band_colors[0..1].map(&:capitalize!).map { |bc| colors.find_index(bc) }.join.to_i
  end
end