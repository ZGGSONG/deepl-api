## deepl_api

逆向 DeepL 客户端自建 API

## docker

```shell
docker pull zggsong/translate && docker run -itd --name deepl -p 4006:8000 zggsong/translate
```

```shell
curl -X POST -d '{"text":"中文国际","source_lang":"auto","target_lang":"EN"}' "localhost:4006/translate"
```
