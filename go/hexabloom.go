package hexabloom

import (
	"math/big"
	"os"

	"github.com/zeebo/xxh3"
)

type BloomFilterClient struct {
	bitmap *big.Int
	k      int
	M      int
}

func (bfc *BloomFilterClient) ContainsStr(value string) bool {
	return bfc.ContainsBytes([]byte(value))
}

func (bfc *BloomFilterClient) ContainsBytes(value []byte) bool {
	// pay attention to the seed values and the hash function used
	hash1 := int(xxh3.HashSeed(value, 0) % uint64(bfc.M))
	hash2 := int(xxh3.HashSeed(value, 32) % uint64(bfc.M))
	for i := 0; i < bfc.k; i++ {
		mo := int((hash1 + i*hash2) % bfc.M)
		bit := bfc.bitmap.Bit(mo)
		if bit == 0 {
			return false
		}
	}
	return true
}

func NewBloomFilterClientFromFile(filePath string) (*BloomFilterClient, error) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	numberOfHashesBytes := fileBytes[:4]
	bitmap := fileBytes[4:]

	// convert byte endianness
	numberOfHashes := int(numberOfHashesBytes[3]) | int(numberOfHashesBytes[2])<<8 | int(numberOfHashesBytes[1])<<16 | int(numberOfHashesBytes[0])<<24
	for i, j := 0, len(bitmap)-1; i < j; i, j = i+1, j-1 {
		bitmap[i], bitmap[j] = bitmap[j], bitmap[i]
	}

	// create bloom filter client
	return &BloomFilterClient{
		bitmap: new(big.Int).SetBytes(bitmap),
		k:      numberOfHashes,
		M:      len(bitmap) * 8,
	}, nil

}
