## Description

Reverse the DeepL client-side self-built API

<a href="https://hub.docker.com/r/zggsong/translate">
  <img alt="Docker pull" src="https://img.shields.io/docker/pulls/zggsong/translate">
</a>

## Usage

### Docker

```shell
docker pull zggsong/translate && docker run -itd --name deepl -p 4006:8000 zggsong/translate
```

### Binary files

Select your system version to download >> [Release](https://github.com/ZGGSONG/deepl-api/releases)

### Curl test

```shell
curl -X POST -d '{"text":"input your content","source_lang":"auto","target_lang":"ZH"}' "localhost:4006/translate"
```

## Languages

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
- `TR`: 土耳其语

> Supportted

```txt
EN
DE
FR
ES
PT
IT
NL
PL
RU
ZH
JA
BG
CS
DA
EL
ET
FI
HU
LT
LV
RO
SK
SL
SV
TR
ID
```


## Author

**deepl-api** © [zggsong](https://github.com/zggsong), Released under the [LGPL-3.0](https://github.com/ZGGSONG/deepl-api/blob/main/LICENSE) License.<br>

> Website [Blog](https://www.zggsong.com) · GitHub [@zggsong](https://github.com/zggsong)
