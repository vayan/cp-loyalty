# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the rails db:seed command (or created alongside the database with db:setup).

LoyaltyRank.create(name: 'bronze', multiplier: 1, required_rides_count: 0)
LoyaltyRank.create(name: 'silver', multiplier: 3, required_rides_count: 5)
LoyaltyRank.create(name: 'gold', multiplier: 5, required_rides_count: 15)
LoyaltyRank.create(name: 'platinum', multiplier: 10, required_rides_count: 30)

