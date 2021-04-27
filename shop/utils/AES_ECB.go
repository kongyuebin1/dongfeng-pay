/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2018-8-30 13:50
 ** @Author : Joker
 ** @File : ACE_ECB
 ** @Last Modified by : Joker
 ** @Last Modified time: 2018-8-30 13:50
 ** @Software: GoLand
****************************************************/
package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"github.com/beego/beego/v2/core/logs"
)

/*
	aes解码
	crypted:要加密的字符串
	key:用来加密的密钥 密钥长度可以是128bit、192bit、256bit中的任意一个
	16位key对应128bit
*/
func AesDecrypt(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		logs.Error("Joker: AesDecrypt failed to NewCipher")
		logs.Error(err)
	}
	blockMode := NewECBDecrypter(block)
	origData := make([]byte, len(src))
	blockMode.CryptBlocks(origData, src)
	origData = PKCS5UnPadding(origData)
	return origData
}

/*aes编码*/
func AesEncrypt(src []byte, key string) []byte {
	decodeString, err := hex.DecodeString(key)
	if err != nil {
		logs.Error("Joker: AesEncrypt failed to hex key")
		logs.Error(err)
	}
	block, err := aes.NewCipher(decodeString)
	if err != nil {
		logs.Error("Joker: AesEncrypt failed to NewCipher")
		logs.Error(err)
	}
	if len(src) < 0 {
		logs.Error("Joker: AesEncrypt`s input is null ")
	}
	ecb := NewECBEncrypter(block)
	src = PKCS5Padding(src, block.BlockSize())
	crypted := make([]byte, len(src))
	ecb.CryptBlocks(crypted, src)
	// 普通base64编码加密 区别于urlsafe base64
	//fmt.Println("base64 result:", base64.StdEncoding.EncodeToString(crypted))
	//fmt.Println("base64UrlSafe result:", Base64UrlSafeEncode(crypted))
	return crypted
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		logs.Error("Joker: CryptBlocks`s input not full blocks")
	}
	if len(dst) < len(src) {
		logs.Error("Joker: CryptBlocks`s output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		logs.Error("Joker: CryptBlocks`s input not full blocks")
	}
	if len(dst) < len(src) {
		logs.Error("Joker: CryptBlocks`s output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
