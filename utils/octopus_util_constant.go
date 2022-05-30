package utils

//定义hash和地址长度byte
const (
	// hash长度
	HashLength = 32
	// 地址值长度
	AddressLength = 20
	//合同创建 gas
	TxGasContractCreation uint64 = 53000
	//交易 gas
	TxGas uint64 = 21000
	//交易数据非零gas限制
	TxDataNonZeroGasFrontier uint64 = 68
	//eip2028类型gas限制
	TxDataNonZeroGasEIP2028 uint64 = 16

	TxDataZeroGas uint64 = 4
)

//定义hash字节类型
type Hash [HashLength]byte

//定义地址字节类型
type Address [AddressLength]byte

// 十六进制将哈希转换为十六进制字符串。
func (h Hash) Hex() string { return Encode(h[:]) }

//BytesToAddress返回值为b的地址。
func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}
