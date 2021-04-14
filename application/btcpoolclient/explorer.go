package btcpoolclient

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Sign(params map[string]string) map[string]string {
	res := make(map[string]string)
	res["nonce"] = RandStringBytes(10)
	res["app_a"] = "6b2llbmFuMmUxb2lpaG9hMm4za2RodWZzZGRmYXMK"
	res["app_b"] = "Y2hhaW5hcHAK"
	res["timestamp"] = fmt.Sprintf("%v", time.Now().Unix())
	ks := make([]string, 0)
	for k, v := range params {
		res[k] = v
		ks = append(ks, k)
	}
	sort.Strings(ks)

	args := make([]string, 0)
	for _, v := range ks {
		args = append(args, res[v])
	}
	secretKey := "7f5a58ff3739aaaf45025599ae5a993269573dd2a215acfccd7cc30eec86efe1"

	hash := hmac.New(sha256.New, []byte(secretKey))
	msg := []byte(strings.Join(args, "|"))
	hash.Write(msg)

	// to base64
	sign := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	res["sign"] = sign
	return res
}

type PoolRankList []PoolRank

type PoolRank struct {
	PoolName               string `json:"pool_name"`
	IconLink               string `json:"icon_link"`
	RealtimeHashrate       string `json:"realtime_hashrate"`
	EstimateHashrate       string `json:"estimate_hashrate"`
	RealtimeCur2maxPercent string `json:"realtime_cur2max_percent"`
	EstimateCur2max        string `json:"estimate_cur2max"`
	HashSuffix             string `json:"hashrate_suffix"`
	RealtimeDiff24hPercent string `json:"realtime_diff_24h_percent"`
}

type RealtimeData struct {
	List []PoolRank `json:"list"`
}
type PoolRankData struct {
	Realtime RealtimeData `json:"realtime"`
}

// 获取矿池排名
func GetPoolRank(c *gin.Context) (map[string]PoolRankData, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]PoolRankData `json:"data"`
	}{}
	// res := make(map[string]string)
	// res["coins"] = "btc,bch"
	// res["show_unconfirm_info"] = "true"
	// fmt.Printf(">>>> sign1 = %v", Sign(res))
	_, err := doRequest(c, "explorer.poolRank", c.Params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error GetPoolRank: %v", err)
	}
	return dest.Data, nil
}

type LatestBlock struct {
	Timestamp string `json:"timestamp"`
	Reward    string `json:"reward"`
	Height    int    `json:"height"`
	PoolName  string `json:"relayed_by_text"`
	Hash      string `json:"hash"`
	Size      int    `json:"size"`
}
type LatestBlockList []LatestBlock
type LatestBlockData struct {
	List LatestBlockList `json:"list"`
}

// 获取最新出块列表
func GetLatestBlockList(c *gin.Context) (map[string]LatestBlockData, error) {
	var dest = struct {
		BtcpoolRescomm
		Data map[string]LatestBlockData `json:"data"`
	}{}

	// res := make(map[string]string)
	// res["coins"] = "btc,bch"
	// res["show_unconfirm_info"] = "true"
	// res["from"] = "ios"
	// fmt.Printf(">>>> sign = %v", Sign(res))
	_, err := doRequest(c, "explorer.blockList", c.Params, &dest)
	if err != nil {
		return nil, fmt.Errorf("error GetLatestBlockList: %v", err)
	}
	return dest.Data, nil
}
