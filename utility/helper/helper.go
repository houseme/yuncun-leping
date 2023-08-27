/*
 * Copyright yuncun-leping Author(https://houseme.github.io/yuncun-leping/). All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * You can obtain one at https://github.com/houseme/yuncun-leping.
 *
 */

package helper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unsafe"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/houseme/gocrypto"
	"github.com/houseme/gocrypto/aes"
	"github.com/houseme/snowflake"
	"golang.org/x/crypto/bcrypt"

	"github.com/houseme/yuncun-leping/utility/env"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	helperUtilSnowflake = "helper.util.snowflake"

	// userAgent .
	headerUserAgent = `Mozilla/5.0 (lanren; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36`
	iv              = "1011011101010100"

	cdnStaticAddress = ""
)

var (
	localInstances = gmap.NewStrAnyMap(true)
	src            = rand.NewSource(time.Now().UnixNano())
)

// Helper .
func Helper() *UtilHelper {
	return &UtilHelper{}
}

// UtilHelper .
type UtilHelper struct {
}

// UserAgent is a default http userAgent
func (h *UtilHelper) UserAgent(_ context.Context) string {
	return headerUserAgent
}

// InitTrxID .根据上下文以及账户标识获取交易订单号
func (h *UtilHelper) InitTrxID(ctx context.Context, ano uint64) uint64 {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-Helper-InitTrxID")
	defer span.End()

	var appEnv, err = env.NewSnowflakeEnv(ctx)
	if err != nil {
		return h.InitTrxID(ctx, ano)
	}
	g.Log(h.Logger(ctx)).Debug(ctx, "appEnv DatacenterID:", appEnv.Datacenter(ctx), " WorkerID:", appEnv.Worker(ctx))
	workerID := appEnv.Worker(ctx)
	if ano > 0 {
		workerID = int64(ano % 32)
	}
	return uint64(h.InitOrderID(ctx, appEnv.Datacenter(ctx), workerID))
}

// InitOrderID init64 order id
func (h *UtilHelper) InitOrderID(ctx context.Context, datacenterID, workerID int64) int64 {
	g.Log(h.Logger(ctx)).Debug(ctx, "InitOrderID DatacenterID:", datacenterID, " WorkerID:", workerID)
	if datacenterID < 0 || datacenterID > snowflake.GetDatacenterIDMax() {
		return 0
	}

	if workerID < 0 || workerID > snowflake.GetWorkerIDMax() {
		return 0
	}
	return int64(h.SnowflakeInstance(ctx, datacenterID, workerID).NextVal())
}

// SnowflakeInstance Get Client Instance
// datacenterID Datacenter ID must be greater than or equal to 0
// workerID Worker ID must be greater than or equal to 0
func (h *UtilHelper) SnowflakeInstance(ctx context.Context, datacenterID, workerID int64) *snowflake.Snowflake {
	instanceKey := fmt.Sprintf("%s.%02d.%02d", helperUtilSnowflake, datacenterID, workerID)
	g.Log(h.Logger(ctx)).Debug(ctx, "InitOrderID SnowflakeInstance ", instanceKey, workerID, datacenterID)
	return localInstances.GetOrSetFuncLock(instanceKey, func() interface{} {
		s, err := snowflake.NewSnowflake(datacenterID, workerID)
		if err != nil {
			panic(err)
		}
		return s
	}).(*snowflake.Snowflake)
}

// AuthToken user auth token
func (h *UtilHelper) AuthToken(ctx context.Context, accountNo uint64) string {
	return gconv.String(h.InitTrxID(ctx, accountNo%32)) + h.InitRandStr(64) + gtime.TimestampNanoStr()
}

// InitRandStr RandStringBytesMaskImprSrcUnsafe
func (h *UtilHelper) InitRandStr(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// UcFirst 首字母大些
func (h *UtilHelper) UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
func (h *UtilHelper) LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// GetOutBoundIP 获取本机 IP
func (h *UtilHelper) GetOutBoundIP(ctx context.Context) string {
	conn, err := net.Dial("udp", "119.29.29.29:80")
	if err != nil {
		g.Log(h.Logger(ctx)).Error(ctx, " GetOutBoundIP udp get Ip failed err: ", err)
		return ""
	}
	defer func() {
		_ = conn.Close()
	}()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// GetLocalIPV4 获取 IPV4 IP，没有则返回空
func (h *UtilHelper) GetLocalIPV4(ctx context.Context) string {
	inters, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, inter := range inters {
		// 判断网卡是否开启，过滤本地环回接口
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			// 获取网卡下所有的地址
			addrs, err := inter.Addrs()
			if err != nil {
				g.Log(h.Logger(ctx)).Error(ctx, " GetLocalIpV4 udp get Ip failed err: ", err)
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					// 判断是否存在 IPV4 IP 如果没有过滤
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

// Logger .获取上下文中的 logger
func (h *UtilHelper) Logger(ctx context.Context) string {
	return gconv.String(ctx.Value("logger"))
}

// EncryptSignData sign data
func (h *UtilHelper) EncryptSignData(_ context.Context, data interface{}, key []byte) ([]byte, error) {
	var byteInfo, err = gjson.Encode(data)
	if err != nil {
		return byteInfo, gerror.Wrap(err, "EncryptSignData json encode failed")
	}
	return aes.NewAESCrypt(key).Encrypt(byteInfo, gocrypto.ECB)
}

// Header .
func (h *UtilHelper) Header(_ context.Context) map[string]string {
	return g.MapStrStr{
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Accept":          "image/avif,image/webp,image/apng,image/*,*/*;q=0.8",
		"Connection":      "keep-alive",
		"User-Agent":      headerUserAgent,
	}
}

// HeaderToMap covert request headers to map.
func (h *UtilHelper) HeaderToMap(header http.Header) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range header {
		if len(v) > 1 {
			m[k] = v
		} else {
			m[k] = v[0]
		}
	}
	return m
}

// EncryptPass .加密处理
func (h *UtilHelper) EncryptPass(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

// CompareHashAndPassword 校验密码。
func (h *UtilHelper) CompareHashAndPassword(inputPass, authPass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(authPass), []byte(inputPass)); err != nil {
		return false
	}
	return true
}

// RequestTime .request time
func (h *UtilHelper) RequestTime(_ context.Context, ts string) *gtime.Time {
	return gtime.NewFromStrFormat(ts, "YmdHis")
}

// ConcatenateSignSource get sign url 排序并拼接签名的内容信息
func (h *UtilHelper) ConcatenateSignSource(ctx context.Context, data interface{}) string {
	ctx, span := gtrace.NewSpan(ctx, "tracing-enterprise-utility-ConcatenateSignSource")
	defer span.End()

	var (
		tt     = reflect.TypeOf(data)
		v      = reflect.ValueOf(data)
		count  = v.NumField()
		keys   = make([]string, 0, count)
		params = make(map[string]string)
		logger = h.Logger(ctx)
	)

	g.Log(logger).Info(ctx, "helper ConcatenateSignSource tt", tt, " v", v)
	for i := 0; i < count; i++ {
		if v.Field(i).CanInterface() { // 判断是否为可导出字段
			g.Log(logger).Printf(ctx, "%s %s = %v -tag:%s", tt.Field(i).Name, tt.Field(i).Type, v.Field(i).Interface(),
				tt.Field(i).Tag)
			keys = append(keys, h.LcFirst(tt.Field(i).Name))
			params[h.LcFirst(tt.Field(i).Name)] = gconv.String(v.Field(i).Interface())
		}
	}
	// sort params
	sort.Strings(keys)
	var buf bytes.Buffer
	for i := range keys {
		k := keys[i]
		if params[k] == "" || k == "sign" {
			continue
		}
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(params[k])
		buf.WriteString("&")
	}
	buf.Truncate(buf.Len() - 1)
	g.Log(logger).Info(ctx, "helper ConcatenateSignSource string end:", buf.String())
	return buf.String()
}

// DecryptSignDataInfo sign data 数据执行 aes 解密
func (h *UtilHelper) DecryptSignDataInfo(src []byte, key []byte) (dst []byte, err error) {
	return aes.NewAESCrypt(key).Decrypt(src, gocrypto.ECB)
}

// HexDecodeString .
func (h *UtilHelper) HexDecodeString(ctx context.Context, data string, key []byte) ([]byte, error) {
	signData, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return h.DecryptSignDataInfo(signData, key)
}

// Sha256Of returns the sha256 of the input string
func (h *UtilHelper) Sha256Of(input []byte) string {
	sum := sha256.Sum256(input)
	return hex.EncodeToString(sum[:])
}

// CheckFileExists .
func (h *UtilHelper) CheckFileExists(ctx context.Context, filePath string) (err error) {
	if !gfile.Exists(filePath) {
		if err = gfile.Mkdir(filePath); err != nil {
			return err
		}
	} else if !gfile.IsDir(filePath) {
		return gerror.NewCode(gcode.CodeInvalidParameter, `parameter "dirPath" should be a directory path`)
	}
	g.Log(h.Logger(ctx)).Info(ctx, "CheckFileExists filePath:", filePath)
	return nil
}

// UserAgentIPHash user agent ip hash
func (h *UtilHelper) UserAgentIPHash(useragent string, ip string) string {
	data, err := h.Sha256OfShort(fmt.Sprintf("%s-%s-%s-%d", useragent, ip, time.Now().String(), rand.Int()))
	if err != nil {
		return ""
	}
	str := h.Base58Encode(data)
	return str[:10]
}

// Sha256OfShort returns the sha256 of the input string
func (h *UtilHelper) Sha256OfShort(input string) ([]byte, error) {
	algorithm := sha256.New()
	if _, err := algorithm.Write([]byte(strings.TrimSpace(input))); err != nil {
		return nil, gerror.Wrap(err, "algorithm write failed")
	}
	return algorithm.Sum(nil), nil
}

// Base58Encode encodes the input byte array to base58 string
func (h *UtilHelper) Base58Encode(data []byte) string {
	return base58.Encode(data)
}

// PasswordBase58Hash password base58 hash
func (h *UtilHelper) PasswordBase58Hash(password string) (string, error) {
	data, err := h.Sha256OfShort(password)
	if err != nil {
		return "", gerror.Wrap(err, "password base58 hash sha256 short failed")
	}
	return h.Base58Encode(data), nil
}

// GenerateShortLink generate short link
func (h *UtilHelper) GenerateShortLink(_ context.Context, url string) (str string, err error) {
	var urlHash []byte
	if urlHash, err = h.Sha256OfShort(url); err != nil {
		return "", gerror.Wrap(err, "generate short link sha256 short failed")
	}
	// number := new(big.Int).SetBytes(urlHash).Uint64()
	str = h.Base58Encode(urlHash)
	return str[:8], nil
}

// AESEncrypt encrypts the input byte array with the given key
func (h *UtilHelper) AESEncrypt(_ context.Context, key, data []byte) (dst string, err error) {
	return aes.NewAESCrypt(key).EncryptToString(gocrypto.Base64, data, gocrypto.ECB)
}

// AESDecrypt decrypts the input byte array with the given key
func (h *UtilHelper) AESDecrypt(_ context.Context, key, data []byte) (dst string, err error) {
	return aes.NewAESCrypt(key).DecryptToString(gocrypto.Base64, data, gocrypto.ECB)
}

// CreateAccessToken create access token
func (h *UtilHelper) CreateAccessToken(ctx context.Context, accountNo uint64) (token string, err error) {
	var hash []byte
	if hash, err = h.Sha256OfShort(gconv.String(h.InitTrxID(ctx, accountNo))); err != nil {
		err = gerror.Wrap(err, "utilHelper CreateAccessToken Sha256OfShort failed")
		return
	}
	token = hex.EncodeToString(hash)
	return
}

// GetOriginPassword .
func (h *UtilHelper) GetOriginPassword(aesPasswd, salt string) (passwd string, err error) {
	decodePasswdByte, _ := hex.DecodeString(aesPasswd)
	aesByte, err := gaes.Decrypt(decodePasswdByte, []byte(salt), []byte(iv))
	if err != nil {
		return "", err
	}
	passwd = string(aesByte)
	return
}

// GetAESPassword .
func (h *UtilHelper) GetAESPassword(originPasswd string) (passwd string, salt string, err error) {
	rand.NewSource(time.Now().UnixNano())
	headSalt := rand.Intn(8) + 1
	salt = strconv.Itoa(headSalt) + grand.Digits(15)
	aesByte, err := gaes.Encrypt([]byte(originPasswd), []byte(salt), []byte(iv))
	if err != nil {
		return "", "", err
	}
	passwd = hex.EncodeToString(aesByte)
	return
}

// ReplaceCdnStaticAddress .
func (h *UtilHelper) ReplaceCdnStaticAddress(ctx context.Context, url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return cdnStaticAddress + url
}
