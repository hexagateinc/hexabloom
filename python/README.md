## Hexabloom - Python

An example code would be:
```python
def main():
    file_path = "bloom.bin"

    bloom_filter_client = BloomFilterClient.from_file_with_hashes(file_path)
    assert bloom_filter_client.contains_str("0x910cbd523d972eb0a6f4cae4618ad62622b39dbf")

if __name__ == "__main__":
    main()
```
