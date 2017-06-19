class User < ApplicationRecord
  validates :loyalty_points, presence: true, numericality: { greater_than_or_equal_to: 0 }
end
