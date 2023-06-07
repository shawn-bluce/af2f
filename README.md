# af2f
Append file to file

go 1.20

# data struct
`[source_file][hidden_file][40bit -> offset][3bit -> version]`

40bit = 128G source file

3bit = 0/1/2/3/4/5/6/7 version

# usage
af2f -a target.jpg bigfile.mp4
af2f -s bigfile.mp4 target.jpg

af2f -k password -a target.jpg bigfile.mp4

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
