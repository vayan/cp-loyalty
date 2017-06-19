class UsersController < ApplicationController
  before_action :set_user, only: [:show]
  before_action :build_user, only: [:new, :create]

  def show
    render json: @user,
      include: [:rides, :loyalty_rank],
      methods: [:rides_left_before_next_rank]
  end

  def new
  end

  def create
    if @user.update_attributes(user_params)
      redirect_to user_path(@user)
    else
      render :new, status: :unprocessable_entity
    end
  end

private
  def user_params
    params.require(:user).permit(:loyalty_points)
  end

  def build_user
    @user ||= User.new(
      loyalty_rank: LoyaltyRank.where(name: "bronze").first
    )
  end

  def set_user
    @user = User.find(params[:id])
  end
end
