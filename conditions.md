# 4433筛选条件

## 基金类型
- 指数型-股票
- 股票型
- 混合型-偏股
- 混合型-偏债
- 混合型-灵活
- 混合型-平衡

## 同类排名百分比

```json
{
    "pageNum":"10",
    "bType":"01,04,21,22,23,24",
    "pageIndex":"1",
    "stageRanking":"4_0-33,5_0-33,6_0-25,7_0-25,8_0-25,9_0-25,10_0-25"
}
```

# 高夏普比率

## 基金类型
- 股票型
- 混合型-偏股
- 混合型-灵活

## 风险指标
近1年/近3年夏普比率前10%

```json
{
    "isBuy":"1",
    "sharpRanking":"8_0-10,6_0-10",
    "pageNum":"10",
    "bType":"01_0-100,24_0-100,21_0-100",
    "pageIndex":"1",
    "isSale":"1",
    "orderField":"5_6_-1"
}
```

# 低回撤

## 基金类型
- 指数型-股票
- 股票型
- 混合型-偏股
- 混合型-灵活
- 混合型-平衡

## 风险指标
收益回撤比近3年排名前20%
阶段回撤近3年小于10%


```json
{
    "isBuy":"1",
    "stageMaxReturn":"8_0-10",
    "establishPeriod":"4",
    "pageNum":"10",
    "cmRanking":"8_0-20,6_0-20",
    "bType":"01_0-100,04_0-100,24_0-100,22_0-100,21_0-100",
    "pageIndex":"1",
    "isSale":"1",
    "orderField":"5_6_-1"
}
```

# 固收+进取

## 风险指标
近3年收益率（35%~253%）
近3年最大回撤10%内

```json
{
    "isBuy":"1",
    "stageMaxReturn":"8_0-10",
    "pageNum":"10",
    "stageSyl":"8_35-253",
    "pageIndex":"1",
    "isSale":"1",
    "orderField":"5_6_-1"
}
```

# "小而美"主动股基

## 基金类型
- 股票型
- 混合型-偏股
- 混合型-灵活

## 风险指标
近3月/近6月/近1年阶段排名前20%
季度超额收益胜率>80%
基金经理管理大于1年
成立大于1年



```json
{
    "isBuy":"1",
    "fundSize":"4,3,2,1",
    "establishPeriod":"2",
    "pageNum":"10",
    "bType":"01_0-100,24_60-100,21_0-100",
    "manageFundPeriod":"1",
    "pageIndex":"1",
    "quarterExcessWinRate":"6_80-100",
    "isSale":"1",
    "orderField":"5_6_-1",
    "stageRanking":"6_0-20,5_0-20,4_0-20"
}
```

## 参数解释
|参数|解释|参数值|
|---|---|---|
|establishPeriod|成立时间|1: >6月<br>2: >1年<br>3: >2年<br>4: >3年<br>5: >4年<br>6: >5年<br>7: <6月<br>8: <1年|
|fundSize|基金规模|1: 0~1亿元<br>2: 1~5亿元<br>3: 5~10亿元<br>4: 10~50亿元<br>5: >50亿元|
|fundLevel|晨星评级|1: 一星<br>2: 二星<br>3: 三星<br>4: 四星<br>5: 五星|
|bType|基金类型|04: 指数型-股票<br>01: 股票<br>21: 混合型-偏股<br>23: 混合型-偏债<br>24: 混合型-灵活<br>22: 混合型-平衡<br>31: 债券型-长债<br>32: 债券型-中短债<br>33: 债券型-混合债<br>34: 债券型-可转债<br>05: 货币型<br>07: QDII<br>25: 商品|
|riskLevel|风险等级|1: 低风险<br>2: 中低风险<br>3: 中风险<br>4: 中高风险<br>5: 高风险|
|stageSyl|阶段收益率|近3月: 4_<br>近6月: 5_<br>近1年: 6_<br>近2年: 7_<br>近3年: 8_<br>近5年: 9_<br>今年以来: 10_<br>(4_10-35，表示近3月收益率在10%-35%；4_10-max，表示近3月收益率在10%到最大值)|
|byYearSyl|每年收益率|2014-2022(表示xxxx_)|
|stageExcessSyl|阶段超额收益率|同阶段收益率|
|buYearExcessSyl|每年超额收益率|同每年收益率|
|stageRanking|阶段排名|同阶段收益率|
|byYearRanking|每年排名|同每年收益率|
|quarterAvgRanking|季度平均排名|过去4季度: 6_<br>过去8季度: 7_<br>过去12季度: 8_<br>过去20季度: 9_|
|yearAvgRanking|年度平均排名|过去2年: 7_<br>过去3年: 8_<br>过去5年: 9_|
|stageMaxReturn|阶段最大回撤|同阶段收益率|
|byYearMaxReturn|每年最大回撤|同每年收益率|
|annualizedVolatility|年化波动率|近1年: 6_<br>近2年: 7_<br>近3年: 8_<br>近5年: 9_|
|sharpRanking|夏普比率同类排名百分比|同年化波动率|
|cmRanking|年收益回撤比同类排名百分比|同年化波动率|
|monthWinRate|月度胜率|同年化波动率|
|quarterWinRate|季度胜率|同年化波动率|
|monthExcessWinRate|月度超额收益胜率|同年化波动率|
|quarterExcessWinRate|季度超额收益胜率|同年化波动率|
|goldBullReward|金牛奖|1或空|
|workPeriod|从业年限|1: >1年<br>2: >2年<br>3: >3年<br>5: >5年<br>10: >10年|
|manageFundPeriod|管理该基金时长|1: >1年<br>2: >2年<br>3: >3年<br>5: >5年|
|paybackSyl|基金经理年化回报率|0-47|
|isSale|代销基金|1或0|
|instituteHeavy|机构重仓|50-100|
|orderField|排序|5_阶段收益对应的前缀_-1降序/1升序|