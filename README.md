# af2f

![](https://github.com/shawn-bluce/af2f/actions/workflows/go.yml/badge.svg?event=release)
![](https://github.com/shawn-bluce/af2f/actions/workflows/go_test.yml/badge.svg)

You have too many reason to hidden a file. Use this tool you can append file A to file B, then open B looks like source B, file A is hiddened. And you can split file A from appended file B.

你有很多种理由想将一个文件隐藏起来。你可以用这个工具将文件 A 追加到文件 B 后面，打开文件 B 查看的时候与追加前保持一致，并且你可以将文件 A 从文件 B 中分离出来。

Forexample you append a.jpg to b.mp4, then open the b.mp4, it can still play normally, and not easy to detect changes in the file.

例如你将一个图片 a.jpg 追加到一个视频文件 b.mp4 之后，打开该视频文件，它是可以正常播放的，并不容易察觉该文件的变化。

# usage
- `af2f append -f avatar_2_4k_hdr.mp4 -a xxx.jpg -p YOUR_PASSWORD`
- `af2f append -f avatar_2_4k_hdr.mp4 -a xxx.jpg -p YOUR_PASSWORD -e aes-192`
- `af2f split -f avatar_2_4k_hdr.mp4 -o fff.jpg -p YOUR_PASSWORD`
- `af2f restore -f avatar_2_4k_hdr.mp4`


# data struct
[source file][appended file][8bytes offset][8bytes algorithm id]

# compile
go version = 1.20

```shell
go get
go build -o af2f
```
