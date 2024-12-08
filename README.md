# unitejson

```bash
go install github.com/goropikari/unitejson@latest
```


```bash
$ cat sample/hoge.json 
{
        // hoge
        "hoge": 123, // hogehoge
        "piyo": "piyopiyo"
}

$ cat sample/piyo.json 
{
        "piyo": "piyo-override",
        "fuga": "fugafuga"
}

$ unitejson sample/hoge.json sample/piyo.json 
{"fuga":"fugafuga","hoge":123,"piyo":"piyo-override"}
```
