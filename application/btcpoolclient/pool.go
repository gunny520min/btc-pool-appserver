package btcpoolclient

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// type ShareHashrates []Ticket

// type Ticket interface{} //[]float64

// type Ticket struct {
// 	x float64
// 	y float64
// }

func GetPoolShareHashrate(c *gin.Context, params interface{}) (interface{}, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]interface{} `json:"data"`
	}{}

	_, err := doRequest(c, "pool.hashrateHistory", params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error getting banner list: %v", err)
	}

	fmt.Printf("ticket= %v", dest.Data["tickers"])
	return dest.Data, nil
	// if tarr, ok := dest.Data["tickers"].([]Ticket); ok {
	// 	// fmt.Printf("ticket data=%v", tarr)

	// 	// res := make([]Ticket, len(tarr))
	// 	// for _, t := range tarr {
	// 	// 	tck := Ticket{t[0], t[1]}
	// 	// 	res = append(res, tck)
	// 	// }
	// 	return tarr, nil
	// } else {
	// 	return []Ticket{}, nil
	// }
}
