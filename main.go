package main

import (
	"fmt"
	"bitflyer_buy_sell/bitflyer"
)

func main() {
  //自分のbitflyerのxにAPIKey・yにシークレットキー
	apiClient := bitflyer.New(xxxxxxxx, yyyyyyyy)
	
	order := &bitflyer.Order{
		ProductCode:     config.Config.ProductCode,
		ChildOrderType:  "MARKET",//販売所から買う/売る 指値の場合は「LIMIT」
		Side:            "BUY",//買いの時は「BUY」売りの時は「SELL」
//     price: 5000000, //「LIMIT」の場合は1BTCがいくらになったら売り買いするか
		Size:            0.001,//いくらbitcoinを売買するか
		MinuteToExpires: 1,
		TimeInForce:     "GTC",
	}
  
  response, _ := apiClient.SendOrder(order)
	fmt.Println(response.ChildOrderAcceptanceID)
  
}
