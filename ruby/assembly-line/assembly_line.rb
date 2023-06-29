class AssemblyLine
  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    base_rph = @speed * 221
    case 
    when @speed >= 1 && @speed <= 4
      base_rph
    when @speed >= 5 && @speed <= 8
      base_rph * 0.9
    when @speed == 9
      base_rph * 0.8
    when @speed == 10
      base_rph * 0.77
    end
  end

  def working_items_per_minute
    (self.production_rate_per_hour / 60).to_i
  end
end
