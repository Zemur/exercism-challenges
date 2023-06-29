class BoutiqueInventory
  def initialize(items)
    @items = items
  end

  def item_names
    @items.map{ |inv| inv[:name] }.flatten.sort
  end

  def cheap
    @items.select{ |inv| inv[:price] < 30.0 }
  end

  def out_of_stock
    @items.select { |inv| inv[:quantity_by_size].empty? } 
  end

  def stock_for_item(name)
    @items.select { |inv| inv[:name] == name }[0][:quantity_by_size]
  end

  def total_stock
    @items.sum { |inv| inv[:quantity_by_size].values.sum }
  end

  private
  attr_reader :items
end