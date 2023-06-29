=begin
Write your code for the 'Food Chain' exercise in this file. Make the tests in
`food_chain_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/food-chain` directory.
=end

module FoodChain
  def self.song
    lyrics = []
    
    animals = [
      {name: 'fly'},
      {name: 'spider', action: 'wriggled and jiggled and tickled inside her.'},
      {name: 'bird', action: "How absurd to swallow a bird!\n"},
      {name: 'cat', action: "Imagine that, to swallow a cat!\n"}, 
      {name: 'dog', action: "What a hog, to swallow a dog!\n"}, 
      {name: 'goat', action: "Just opened her throat and swallowed a goat!\n"}, 
      {name: 'cow', action: "I don't know how she swallowed a cow!\n"}, 
      {name: 'horse', action: "She's dead, of course!\n"}
    ]
    
    first_line = "I know an old lady who swallowed a %s.\n"
    swallowed_line = "She swallowed the %s to catch the %s.\n"
    last_line = "I don't know why she swallowed the fly. Perhaps she'll die.\n\n"

    animals.each_index do |i|
      lyrics.append(sprintf(first_line, animals[i][:name]))
      if animals[i][:name] == 'horse'
        lyrics.append(animals[i][:action])
        break
      end
    
      if i == 1
        lyrics.append(sprintf("It %s\n", animals[i][:action]))
      elsif i >= 2 
        lyrics.append(animals[i][:action])
      end
      (1..i).reverse_each do |j|
        if animals[j-1][:name] == 'spider'
          prev_animal = "#{animals[j-1][:name]} that wriggled and jiggled and tickled inside her"
        else
          prev_animal = animals[j-1][:name]
        end
        lyrics.append(sprintf(swallowed_line, animals[j][:name], prev_animal))
      end
      lyrics.append(last_line)
    end

    lyrics.join()
  end
end
