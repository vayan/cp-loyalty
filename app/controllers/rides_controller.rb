class RidesController < ApplicationController
  before_action :set_user

  def new
    @ride ||= @user.rides.new
  end

  def create
    @ride = @user.rides.create(ride_params)
    if @ride.persisted?
      LoyaltyRankUpdater.call(@user)
      LoyaltyPointUpdater.call(@user, @ride)
      redirect_to user_path(@user)
    else
      render :new, status: :unprocessable_entity
    end
  end

private
  def ride_params
    params.require(:ride).permit(:price)
  end

  def set_user
    @user = User.find(params[:user_id])
  end
end
