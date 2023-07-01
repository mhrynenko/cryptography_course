# Large Numbers Practice

## Task
1. Using one of the specific libraries (you can build it yourself, but it is better to use ready-made for time economy)
to output several key options that can be set to <i>8-, 16-, 32-, 64-, 128-, 256-, 512-, 1024 -, 2048-, 4096-</i>bit sequence.
   1. <b>Example:</b> If the length of the key is equal to 16 bits - the length of the key is equal to 65,536.   
      1. The key space is the number of unique keys that are in a given range.
2. For each option, it is necessary to generate a random key value that is recognized in the range from <i>0x00...0</i> to
<i>0xFF...F</i> depending on the reverse key.
3. Write a function to brute force the values from the range in order to find the key. Function purpose is to iterate 
over the value of the key from <i>0x00...0</i> until a value equal to the pre-generated key is set. The function should show 
the amount of time in milliseconds that was calculated for the key value.

## Solution

- Some notes:
  1. To find space size I just used 2^(bit size)
  2. To generate random key `ctypto/rand` library was used
  3. To brute force generated key was used a simple loop that adding 1 if the number is not we are looking for
- Screenshots:
  1. 8-bit
  2. 16-bit
  3. 32-bit
  4. 64-bit
  5. 128-bit
  6. 256-bit
  7. 512-bit
  8. 1024-bit 
  9. 2048-bit
  10. 4096-bit

## Note
1. As developing language was chosen `Golang`
2. To run the code, you need to have go installed 
3. Clone repo
    ```shell
    git clone https://github.com/mhrynenko/cryptography_course
    ```
4. Go to the `cryptography_course/large_numbers` repo
    ```shell
    cd cryptography_course/large_numbers
    ```
5. Run the code
    ```shell
    go run main.go
    ```