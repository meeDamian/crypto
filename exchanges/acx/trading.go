package acx

const openOrdersUrl = "https://acx.io/api/v2/orders.json"

//func ThisWillChange_OpenOrders(c crypto.Credentials) (o []crypto.Order, err error) {
//	resp, err := privateRequest(c, "GET", openOrdersUrl, map[string]string{
//		"market": "btcaud",
//	})
//	if err != nil {
//		return []crypto.Order{}, err
//	}
//
//	defer resp.Body.Close()
//
//	respBody, _ := ioutil.ReadAll(resp.Body)
//
//	log.Println(resp.Status)
//	log.Println(string(respBody))
//
//	return
//}
