

class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze
  class UnsupportedOperation < StandardError
  end

  def self.calculate(first_operand, second_operand, operation)
    # Raising useless error as there is no reason the method should care if the operands are strings or integers.
    # They get cast to strings regardless.
    raise ArgumentError.new("Operands must be Integers") if !first_operand.is_a? Integer or !second_operand.is_a? Integer

    # Raising custom error for unsupported operations.
    raise UnsupportedOperation.new("#{operation} is not allowed. Allowed operations are [#{ALLOWED_OPERATIONS}]") if !ALLOWED_OPERATIONS.include? operation
    
    begin
      problem = "#{first_operand} #{operation} #{second_operand}"
      "#{problem} = #{eval(problem)}"
    rescue => e
      case e.class.name
      when "ZeroDivisionError"
        "Division by zero is not allowed."
      else
        e
      end
    end
  end
end
