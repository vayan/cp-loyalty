class CreateUsers < ActiveRecord::Migration[5.0]
  def change
    create_table :users do |t|
      t.integer :loyalty_points, null: false, default: 0

      t.timestamps
    end
  end
end
