# af2f
Append file to file

go 1.20

# data struct
`[source_file][hidden_file][8byte source_file_size][8byte algorithm][8byte version]`


3bit = 0/1/2/3/4/5/6/7 version

# usage
af2f append -f bigfile.mp4 -a target.jpg -k password
af2f split -f bigfile.mp4 -o target.jpg -k password
af2f check -f bigfile.mp4

# fake code
append:
```c
offset = size(source_file)
// maybe encrypt
result_file = source_file + hidden_file + offset
done.
```

split:
```c
offset = get_offset() // last 40 bit
read with seed offset
stop = size(bigfile) - 40bit
read ... to stop and write
// maybe decrypt
done.
```


# roadmap
1. simple append/split without stream and encrypt
2. add encrypt with aes-256(default)
3. add encrypt algorithm arguments
4. add stream, don't write twice when encrypt
5. 
