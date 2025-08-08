# Please implement your solution to acronym in this file

module Acronym
    def self.abbreviate(s : String)
        s.split(/[\s-_]/).reject(&.empty?).map(&.char_at(0)).join.upcase
    end
end
        