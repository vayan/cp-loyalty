class AddMultiplierToLoyaltyRanks < ActiveRecord::Migration[5.0]
  def change
    add_column :loyalty_ranks, :multiplier, :integer, null: false, default: 0
  end
end
