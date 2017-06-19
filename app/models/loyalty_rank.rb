class LoyaltyRank < ApplicationRecord
  has_many :users

  validates :multiplier, presence: true, numericality: { greater_than: 0 }
  validates :required_rides_count, presence: true, numericality: { greater_than_or_equal_to: 0 }
  validates :name, uniqueness: true
end
