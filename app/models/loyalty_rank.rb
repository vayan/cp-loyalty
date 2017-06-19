class LoyaltyRank < ApplicationRecord
  has_many :users

  validates :required_rides_count, presence: true, numericality: { greater_than_or_equal_to: 0 }
  validates :name, uniqueness: true
end
