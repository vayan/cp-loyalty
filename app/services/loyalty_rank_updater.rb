class LoyaltyRankUpdater
  def self.call(*args)
    new(*args).call
  end

  def initialize(user)
    @user = user
  end

  def call
    new_loyalty_rank = loyalty_rank_for_rides_count
    if @user.loyalty_rank != new_loyalty_rank
      @user.update_attributes(loyalty_rank: new_loyalty_rank)
    end
  end

private
  def loyalty_rank_for_rides_count
    LoyaltyRank
      .where("required_rides_count <= ?", @user.rides.count)
      .order(required_rides_count: :asc)
      .last
  end
end
