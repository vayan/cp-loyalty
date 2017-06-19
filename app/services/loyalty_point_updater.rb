class LoyaltyPointUpdater
  def self.call(*args)
    new(*args).call
  end

  def initialize(user, ride)
    @user = user
    @ride = ride
  end

  def call
    new_points = @ride.price
    @user.update_attributes(loyalty_points: @user.loyalty_points + new_points)
  end
end
