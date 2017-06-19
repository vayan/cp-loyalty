class AddLoyaltyRankToUsers < ActiveRecord::Migration[5.0]
  def change
    add_reference :users, :loyalty_rank, index: true
  end
end
