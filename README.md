SDSCALC
=====

> 抑郁自评量表（Self-rating depression scale，SDS），是含有20个项目，分为4级评分的自评量表，原型是W.K.Zung编制的抑郁量表（1965）。其特点是使用简便，并能相当直观地反映抑郁患者的主观感受及其在治疗中的变化。主要适用于具有抑郁症状的成年人，包括门诊及住院患者。

本程序的题目以及计算方法均依据 [抑郁自评量表（Self-rating depression scale，SDS）的百度百科](https://baike.baidu.com/item/%E6%8A%91%E9%83%81%E8%87%AA%E8%AF%84%E9%87%8F%E8%A1%A8)。

## 运行

在文件创建成功的前提下，运行过程以及结果将会完整记录到文件 sds.out 中。

### Windows

```shell
sdscalc-windows-amd64.exe [-d datasource.yaml]
```

### Linux

```shell
./sdscalc-linux-amd64 [-d datasource.yaml]
```

### MacOS

Apple Chip

```shell
./sdscalc-darwin-arm64 [-d datasource.yaml]
```

X86

```shell
./sdscalc-darwin-amd64 [-d datasource.yaml]
```
