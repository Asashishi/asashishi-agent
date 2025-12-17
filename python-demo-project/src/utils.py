
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
工具函数模块
包含各种辅助函数
"""

from typing import Union, Optional


def format_result(
    operation: str,
    num1: Union[int, float],
    num2: Union[int, float],
    result: Union[int, float]
) -> str:
    """
    格式化运算结果
    
    Args:
        operation: 操作名称
        num1: 第一个数字
        num2: 第二个数字
        result: 运算结果
        
    Returns:
        str: 格式化的结果字符串
    """
    # 根据操作类型选择符号
    symbols: dict[str, str] = {
        "加法": "+",
        "减法": "-",
        "乘法": "×",
        "除法": "÷"
    }
    
    symbol: str = symbols.get(operation, "?")
    
    # 格式化数字，如果是整数则显示为整数，否则保留两位小数
    def format_number(num: Union[int, float]) -> str:
        if isinstance(num, int) or num.is_integer():
            return str(int(num))
        return f"{num:.2f}"
    
    num1_str: str = format_number(num1)
    num2_str: str = format_number(num2)
    result_str: str = format_number(result)
    
    return f"{num1_str} {symbol} {num2_str} = {result_str}"


def validate_division(divisor: Union[int, float]) -> bool:
    """
    验证除法运算的除数是否有效
    
    Args:
        divisor: 除数
        
    Returns:
        bool: 如果除数不为零则返回 True，否则返回 False
    """
    return divisor != 0


def is_even(number: int) -> bool:
    """
    判断数字是否为偶数
    
    Args:
        number: 要判断的数字
        
    Returns:
        bool: 如果是偶数返回 True，否则返回 False
    """
    return number % 2 == 0


def is_prime(number: int) -> bool:
    """
    判断数字是否为质数
    
    Args:
        number: 要判断的数字
        
    Returns:
        bool: 如果是质数返回 True，否则返回 False
    """
    if number < 2:
        return False
    if number == 2:
        return True
    if number % 2 == 0:
        return False
    
    # 检查从 3 到 sqrt(number) 的奇数
    i: int = 3
    while i * i <= number:
        if number % i == 0:
            return False
        i += 2
    
    return True


def fibonacci(n: int) -> list[int]:
    """
    生成斐波那契数列的前 n 项
    
    Args:
        n: 要生成的项数
        
    Returns:
        list[int]: 斐波那契数列的前 n 项
        
    Raises:
        ValueError: 当 n 小于 1 时抛出
    """
    if n < 1:
        raise ValueError("n 必须大于等于 1")
    
    if n == 1:
        return [0]
    elif n == 2:
        return [0, 1]
    
    sequence: list[int] = [0, 1]
    for i in range(2, n):
        next_value: int = sequence[i-1] + sequence[i-2]
        sequence.append(next_value)
    
    return sequence


def calculate_average(numbers: list[Union[int, float]]) -> Optional[float]:
    """
    计算数字列表的平均值
    
    Args:
        numbers: 数字列表
        
    Returns:
        Optional[float]: 平均值，如果列表为空则返回 None
    """
    if not numbers:
        return None
    
    total: float = sum(numbers)
    return total / len(numbers)


def find_max_min(numbers: list[Union[int, float]]) -> tuple[Optional[float], Optional[float]]:
    """
    查找数字列表的最大值和最小值
    
    Args:
        numbers: 数字列表
        
    Returns:
        tuple[Optional[float], Optional[float]]: (最大值, 最小值)，如果列表为空则返回 (None, None)
    """
    if not numbers:
        return None, None
    
    max_value: float = max(numbers)
    min_value: float = min(numbers)
    return max_value, min_value


def celsius_to_fahrenheit(celsius: float) -> float:
    """
    摄氏度转华氏度
    
    Args:
        celsius: 摄氏度
        
    Returns:
        float: 华氏度
    """
    return (celsius * 9/5) + 32


def fahrenheit_to_celsius(fahrenheit: float) -> float:
    """
    华氏度转摄氏度
    
    Args:
        fahrenheit: 华氏度
        
    Returns:
        float: 摄氏度
    """
    return (fahrenheit - 32) * 5/9