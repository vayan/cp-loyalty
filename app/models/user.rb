class User < ApplicationRecord
  has_many    :rides
  belongs_to  :loyalty_rank

  validates :loyalty_points, presence: true, numericality: { greater_than_or_equal_to: 0 }
  validates :loyalty_rank, presence: true

  def rides_left_before_next_rank
    LoyaltyRank
      .where("required_rides_count > ?", rides.count)
      .order(required_rides_count: :asc)
      .first
      .required_rides_count - rides.count
  end
end
