What
====

eve-greenspan is a financial advisor robot for the Eve Online game.


Roadmap
=======

* configuration
	* fetch Eve character list
	* (done) fetch Eve market orders
	* fetch Eve wallet
	* fetch Eve transaction journal
	* limit potential buys by price

* orders that need price bump
	* (done) fetch Eve market orders
	* fetch Eve-Central prices
	* indicate potential loss from outbidders
	* compare and output

* your profit per item type
	* fetch Eve transaction journal
	* calculate and output

* industry opportunities (Isk Per Hour can do this quite well, could use better UI)
	* fetch Eve-Central prices
	* calculate and output

* monitor market
	* fetch Eve-Central prices for cold start
	* connect to Eve Market Data Relay
	* alert on profitable routes
	* alert on cheap sell orders for items of interest

* monitor mining (ore/mineral orders)
	* fetch Eve-Central prices
	* predict mineral supply/demand
	* predict product prices

* monitor PvE (loot/salvage orders)
	* fetch Eve-Central prices
	* predict mineral/salvage prices
	* predict product prices

* monitor PvP (kill boards)
	* fetch kill boards history
	* predict ship/module demand
