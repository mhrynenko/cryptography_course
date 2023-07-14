# ECDSA

## Task
1. Implement digital signature algorithm

## Solution

- Some notes:
    1. <b>ECDSA</b> was chosen.
    2. <i>Why?</i> It is widely used in big systems (e.g. Bitcoin) and I'd like to have some knowledge of how it
    is working under the hood.
    3. As documentation, I used this [paper](https://datatracker.ietf.org/doc/html/rfc6979#appendix-A.2.5) 
    and document from task
    4. Test data can be found in `ecdsa_test.go` file
    5. As hash function for message I used SHA256 that I create in previous task



### Note
1. As developing language was chosen `Golang`
2. To run the code, you need to have go installed
3. Clone repo
    ```shell
    git clone https://github.com/mhrynenko/cryptography_course
    ```
4. Go to the `cryptography_course/ecdsa` repo
    ```shell
    cd cryptography_course/ecdsa
    ```
5. Run tests
    ```shell
    go test
    ```
