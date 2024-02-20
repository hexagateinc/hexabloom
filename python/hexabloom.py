import bitarray

# intdigest returns the result as an integer, which is easier to work with compared to a btyes object
from xxhash import xxh3_64_intdigest

class BloomFilterClient:
    def __init__(self, bitmap, number_of_hashes: int):
        self.bitmap = bitmap
        self.k = number_of_hashes
        self.M = len(bitmap)

    def contains_str(self, value: str) -> bool:
        return self.contains_bytes(value.encode("utf-8"))

    def contains_bytes(self, value: bytes) -> bool:
        # pay attention to the seed values and the hash function usef
        hash1 = xxh3_64_intdigest(value, seed=0) % self.M
        hash2 = xxh3_64_intdigest(value, seed=32) % self.M
        res = self.bitmap[hash1]
        for i in range(self.k):
            mo = (hash1 + i * hash2) % self.M
            if not self.bitmap[mo]:
                return False
        return res

    @classmethod
    def from_file_with_hashes(cls, file_path: str) -> "BloomFilterClient":
        with open(file_path, "rb") as f:
            hashes = int.from_bytes(f.read(4), "big")
            bitmap = bitarray.bitarray(endian="little")
            bitmap.fromfile(f)
            return cls(bitmap, hashes)