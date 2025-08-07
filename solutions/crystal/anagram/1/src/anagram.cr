# Please implement your solution to anagram in this file

module Anagram
    def self.find(word : String, candidates : Array(String))
        permutations = word.downcase.chars.permutations.map(&.join).reject(word.downcase)
        candidates.select {|c| permutations.includes?(c.downcase)}
    end
end