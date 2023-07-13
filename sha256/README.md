# SHA256

## Task
1. Implement hashing algorithm

## Solution

- Some notes:
    1. <b>SHA256</b> was chosen.
    2. Why? The choice was between SHA256 and Keccak256, after some research and reading white papers,
  I understood that SHA256 is slightly easier to implement compared to Keccak, it has more clear documentation
  and at the same time it is still widely used and considered to be secure.
    3. As documentation, I used this [paper](https://csrc.nist.gov/csrc/media/publications/fips/180/2/archive/2002-08-01/documents/fips180-2.pdf)
    4. Test data can be found in `sha256_test.go` file 



### Note
1. As developing language was chosen `Golang`
2. To run the code, you need to have go installed
3. Clone repo
    ```shell
    git clone https://github.com/mhrynenko/cryptography_course
    ```
4. Go to the `cryptography_course/sha256` repo
    ```shell
    cd cryptography_course/sha256
    ```
5. Run tests
    ```shell
    go test
    ```
