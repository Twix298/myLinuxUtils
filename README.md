<h1>Утилита myFind</h1>
Рекурсивный поиск каталогов, симлинков и файлов. Поддерживает флаги "-sl", "-d" или "-f".

```
~$ ./myFind /foo
```

Так же реализована опция  "-ext"(работает только с -f)

```
  ~$ ./myFind -f -ext 'go' /go
```
```
/go/src/github.com//main.go
/go/src/github.com//magic.go
```

<h1>Утилита myWc</h1>

<h3>Подсчет слов в файле</h3>

```
~$ ./myWc -w input.txt
```
```
777 input.txt
```
<h3> Подсчет линий в файлах input2.txt and input3.txt </h3>

```
~$ ./myWc -l input2.txt input3.txt
```
```
42 input2.txt
53 input3.txt
```
 <h3>Подсчет символов в файлах input4.txt, input5.txt and input6.txt</h3>
 
 ```
~$ ./myWc -m input4.txt input5.txt input6.txt
```
```
1337 input4.txt
2664 input5.txt
3991 input6.txt
```

<h1>Утилита myXargs</h1>

https://losst.pro/komanda-xargs-linux

<h1>Утилита myRotate</h1>

Утилита ротации журналов. Старый файл журнала архивируется и помещается на хранение, чтобы журналы не накапливались в одном файле бесконечно.

```
# Будет создан файл /path/to/logs/some_application_1600785299.tag.gz
# где 1600785299 - UNIX timestamp создания `some_application.log`
~$ ./myRotate /path/to/logs/some_application.log
```

```
# Будут созданы 2 файла tar.gz
# и помещены в директорию /data/archive
~$ ./myRotate -a /data/archive /path/to/logs/some_application.log /path/to/logs/other_application.log
```

