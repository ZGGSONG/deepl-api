## deepl_api

逆向 DeepL 客户端自建 API

## docker

```shell
docker pull zggsong/translate && docker run -itd --name deepl -p 4006:8000 zggsong/translate
```

```shell
curl -X POST -d '{"text":"input your content","source_lang":"auto","target_lang":"ZH"}' "localhost:4006/translate"
```

> language
- `DE`: 德语
- `EN`: 英语
- `ES`: 西班牙语
- `FR`: 法语
- `IT`: 意大利语
- `JA`: 日语
- `NL`: 荷兰语
- `PL`: 波兰语
- `PT`: 葡萄牙语
- `RU`: 俄语
- `ZH`: 中文
- `BG`: 保加利亚语
- `CS`: 捷克语
- `DA`: 丹麦语
- `EL`: 希腊语
- `ET`: 爱沙尼亚语
- `FI`: 芬兰语
- `HU`: 匈牙利语
- `LT`: 立陶宛语
- `LV`: 拉脱维亚语
- `RO`: 罗马尼亚语
- `SK`: 斯洛伐克语
- `SL`: 斯洛文尼亚语
- `SV`: 瑞典语