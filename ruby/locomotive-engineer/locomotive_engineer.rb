class LocomotiveEngineer
  def self.generate_list_of_wagons(*args)
    args
  end

  def self.fix_list_of_wagons(each_wagons_id, missing_wagons)
    each_wagons_id.rotate!(2).insert(each_wagons_id.find_index(1)+1, missing_wagons).flatten
  end

  def self.add_missing_stops(routing_hash, **args)
    routing_hash[:stops] = args.to_h.values
    routing_hash
  end

  def self.extend_route_information(route, more_route_information)
    route.merge(more_route_information)
  end
end
