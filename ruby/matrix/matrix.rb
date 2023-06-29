=begin
Write your code for the 'Matrix' exercise in this file. Make the tests in
`matrix_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/matrix` directory.
=end

class Matrix
  def initialize(m)
    @matrix = m.split("\n").map!(&:split).each { |i| i.map!(&:to_i) }
  end

  def row(row_num)
    @matrix[row_num-1]
  end

  def column(col_num)
    @matrix.transpose()[col_num-1]
  end
end