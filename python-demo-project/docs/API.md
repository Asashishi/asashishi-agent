
# Python 演示项目 API 文档 📚

## 概述

这是一个简单的 Python 演示项目，包含计算器功能和工具函数。

## 模块说明

### 1. 计算器模块 (calculator.py)

#### Calculator 类

提供基本的数学运算功能。

##### 构造函数
```python
def __init__(self) -> None
```
初始化一个新的计算器实例。

##### 方法

**add(a, b)**
```python
def add(self, a: Union[int, float], b: Union[int, float]) -> float
```
执行加法运算。

**参数:**
- `a`: 第一个数字
- `b`: 第二个数字

**返回值:** 两个数字的和

**示例:**
```python
calc = Calculator()
result = calc.add(5, 3)  # 返回 8.0
```

**subtract(a, b)**
```python
def subtract(self, a: Union[int, float], b: Union[int, float]) -> float
```
执行减法运算。

**参数:**
- `a`: 被减数
- `b`: 减数

**返回值:** 两个数字的差

**示例:**
```python
result = calc.subtract(10, 4)  # 返回 6.0
```

**multiply(a, b)**
```python
def multiply(self, a: Union[int, float], b: Union[int, float]) -> float
```
执行乘法运算。

**参数:**
- `a`: 第一个数字
- `b`: 第二个数字

**返回值:** 两个数字的积

**示例:**
```python
result = calc.multiply(6, 7)  # 返回 42.0
```

**divide(a, b)**
```python
def divide(self, a: Union[int, float], b: Union[int, float]) -> float
```
执行除法运算。

**参数:**
- `a`: 被除数
- `b`: 除数

**返回值:** 两个数字的商

**异常:**
- `ZeroDivisionError`: 当除数为零时抛出

**示例:**
```python
result = calc.divide(10, 2)  # 返回 5.0
```

**get_operation_count()**
```python
def get_operation_count(self) -> int
```
获取执行过的操作总数。

**返回值:** 操作次数

**示例:**
```python
count = calc.get_operation_count()  # 返回执行过的操作数量
```

**get_last_result()**
```python
def get_last_result(self) -> Optional[float]
```
获取上一次操作的结果。

**返回值:** 上一次操作的结果，如果没有操作则返回 `None`

**示例:**
```python
last_result = calc.get_last_result()
```

**reset()**
```python
def reset(self) -> None
```
重置计算器状态（操作计数和最后结果）。

**示例:**
```python
calc.reset()  # 重置计算器
```

### 2. 工具函数模块 (utils.py)

#### 格式化函数

**format_result(operation, num1, num2, result)**
```python
def format_result(
    operation: str,
    num1: Union[int, float],
    num2: Union[int, float],
    result: Union[int, float]
) -> str
```
格式化运算结果为可读的字符串。

**参数:**
- `operation`: 操作名称（"加法"、"减法"、"乘法"、"除法"）
- `num1`: 第一个数字
- `num2`: 第二个数字
- `result`: 运算结果

**返回值:** 格式化的字符串

**示例:**
```python
formatted = format_result("加法", 5, 3, 8)
# 返回 "5 + 3 = 8"
```

#### 验证函数

**validate_division(divisor)**
```python
def validate_division(divisor: Union[int, float]) -> bool
```
验证除法运算的除数是否有效。

**参数:**
- `divisor`: 除数

**返回值:** 如果除数不为零则返回 `True`，否则返回 `False`

**示例:**
```python
is_valid = validate_division(5)  # 返回 True
is_valid = validate_division(0)  # 返回 False
```

#### 数学函数

**is_even(number)**
```python
def is_even(number: int) -> bool
```
判断数字是否为偶数。

**参数:**
- `number`: 要判断的数字

**返回值:** 如果是偶数返回 `True`，否则返回 `False`

**示例:**
```python
is_even(4)  # 返回 True
is_even(5)  # 返回 False
```

**is_prime(number)**
```python
def is_prime(number: int) -> bool
```
判断数字是否为质数。

**参数:**
- `number`: 要判断的数字

**返回值:** 如果是质数返回 `True`，否则返回 `False`

**示例:**
```python
is_prime(7)  # 返回 True
is_prime(8)  # 返回 False
```

**fibonacci(n)**
```python
def fibonacci(n: int) -> list[int]
```
生成斐波那契数列的前 n 项。

**参数:**
- `n`: 要生成的项数

**返回值:** 斐波那契数列的前 n 项

**异常:**
- `ValueError`: 当 n 小于 1 时抛出

**示例:**
```python
sequence = fibonacci(5)  # 返回 [0, 1, 1, 2, 3]
```

**calculate_average(numbers)**
```python
def calculate_average(numbers: list[Union[int, float]]) -> Optional[float]
```
计算数字列表的平均值。

**参数:**
- `numbers`: 数字列表

**返回值:** 平均值，如果列表为空则返回 `None`

**示例:**
```python
avg = calculate_average([1, 2, 3, 4, 5])  # 返回 3.0
```

**find_max_min(numbers)**
```python
def find_max_min(numbers: list[Union[int, float]]) -> tuple[Optional[float], Optional[float]]
```
查找数字列表的最大值和最小值。

**参数:**
- `numbers`: 数字列表

**返回值:** (最大值, 最小值)，如果列表为空则返回 `(None, None)`

**示例:**
```python
max_val, min_val = find_max_min([5, 2, 8, 1, 9])  # 返回 (9.0, 1.0)
```

#### 温度转换函数

**celsius_to_fahrenheit(celsius)**
```python
def celsius_to_fahrenheit(celsius: float) -> float
```
摄氏度转华氏度。

**参数:**
- `celsius`: 摄氏度

**返回值:** 华氏度

**示例:**
```python
fahrenheit = celsius_to_fahrenheit(0)  # 返回 32.0
```

**fahrenheit_to_celsius(fahrenheit)**
```python
def fahrenheit_to_celsius(fahrenheit: float) -> float
```
华氏度转摄氏度。

**参数:**
- `fahrenheit`: 华氏度

**返回值:** 摄氏度

**示例:**
```python
celsius = fahrenheit_to_celsius(32)  # 返回 0.0
```

## 主程序 (main.py)

### 功能
- 交互式计算器界面
- 支持四则运算
- 自动日志记录
- 错误处理

### 运行方式
```bash
python main.py
```

### 菜单选项
1. 加法运算
2. 减法运算
3. 乘法运算
4. 除法运算
5. 显示帮助
0. 退出程序

## 测试

### 运行所有测试
```bash
python -m unittest discover tests
```

### 运行特定测试
```bash
python -m unittest tests.test_calculator
python -m unittest tests.test_utils
```

## 日志系统

程序会自动在 `log/` 目录下创建日志文件，格式为 `YYYYMMDDHHMMSS.md`。

日志内容包括：
- 程序启动信息
- 每次操作记录
- 错误信息
- 程序退出信息