
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
计算器模块
提供基本的数学运算功能
"""

from typing import Union, Optional


class Calculator:
    """计算器类，提供基本的数学运算"""
    
    def __init__(self) -> None:
        """初始化计算器"""
        self._operation_count: int = 0
        self._last_result: Optional[float] = None
    
    def add(self, a: Union[int, float], b: Union[int, float]) -> float:
        """
        加法运算
        
        Args:
            a: 第一个数字
            b: 第二个数字
            
        Returns:
            float: 两个数字的和
        """
        self._operation_count += 1
        result: float = float(a + b)
        self._last_result = result
        return result
    
    def subtract(self, a: Union[int, float], b: Union[int, float]) -> float:
        """
        减法运算
        
        Args:
            a: 被减数
            b: 减数
            
        Returns:
            float: 两个数字的差
        """
        self._operation_count += 1
        result: float = float(a - b)
        self._last_result = result
        return result
    
    def multiply(self, a: Union[int, float], b: Union[int, float]) -> float:
        """
        乘法运算
        
        Args:
            a: 第一个数字
            b: 第二个数字
            
        Returns:
            float: 两个数字的积
        """
        self._operation_count += 1
        result: float = float(a * b)
        self._last_result = result
        return result
    
    def divide(self, a: Union[int, float], b: Union[int, float]) -> float:
        """
        除法运算
        
        Args:
            a: 被除数
            b: 除数
            
        Returns:
            float: 两个数字的商
            
        Raises:
            ZeroDivisionError: 当除数为零时抛出
        """
        if b == 0:
            raise ZeroDivisionError("除数不能为零")
        
        self._operation_count += 1
        result: float = float(a / b)
        self._last_result = result
        return result
    
    def get_operation_count(self) -> int:
        """
        获取操作次数
        
        Returns:
            int: 执行过的操作总数
        """
        return self._operation_count
    
    def get_last_result(self) -> Optional[float]:
        """
        获取上一次操作的结果
        
        Returns:
            Optional[float]: 上一次操作的结果，如果没有操作则返回 None
        """
        return self._last_result
    
    def reset(self) -> None:
        """重置计算器状态"""
        self._operation_count = 0
        self._last_result = None
    
    def __str__(self) -> str:
        """返回计算器的字符串表示"""
        return f"Calculator(operations={self._operation_count}, last_result={self._last_result})"
    
    def __repr__(self) -> str:
        """返回计算器的详细表示"""
        return f"Calculator(operation_count={self._operation_count}, last_result={self._last_result})"