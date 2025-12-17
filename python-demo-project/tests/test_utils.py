
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
工具函数模块测试
"""

import sys
import unittest
from typing import NoReturn

# 添加 src 目录到路径
sys.path.insert(0, '../src')

from utils import (
    format_result,
    validate_division,
    is_even,
    is_prime,
    fibonacci,
    calculate_average,
    find_max_min,
    celsius_to_fahrenheit,
    fahrenheit_to_celsius
)


class TestUtils(unittest.TestCase):
    """工具函数测试类"""
    
    def test_format_result_addition(self) -> None:
        """测试加法结果格式化"""
        result: str = format_result("加法", 5, 3, 8)
        self.assertEqual(result, "5 + 3 = 8")
    
    def test_format_result_subtraction(self) -> None:
        """测试减法结果格式化"""
        result: str = format_result("减法", 10, 4, 6)
        self.assertEqual(result, "10 - 4 = 6")
    
    def test_format_result_multiplication(self) -> None:
        """测试乘法结果格式化"""
        result: str = format_result("乘法", 6, 7, 42)
        self.assertEqual(result, "6 × 7 = 42")
    
    def test_format_result_division(self) -> None:
        """测试除法结果格式化"""
        result: str = format_result("除法", 10, 2, 5)
        self.assertEqual(result, "10 ÷ 2 = 5")
    
    def test_format_result_float_numbers(self) -> None:
        """测试浮点数格式化"""
        result: str = format_result("加法", 5.5, 2.5, 8.0)
        self.assertEqual(result, "5.5 + 2.5 = 8.0")
    
    def test_format_result_float_result(self) -> None:
        """测试浮点数结果格式化"""
        result: str = format_result("除法", 5, 2, 2.5)
        self.assertEqual(result, "5 ÷ 2 = 2.5")
    
    def test_validate_division_valid(self) -> None:
        """测试有效除数验证"""
        self.assertTrue(validate_division(5))
        self.assertTrue(validate_division(-3))
        self.assertTrue(validate_division(0.5))
    
    def test_validate_division_invalid(self) -> None:
        """测试无效除数验证"""
        self.assertFalse(validate_division(0))
        self.assertFalse(validate_division(0.0))
    
    def test_is_even_true(self) -> None:
        """测试偶数判断（真）"""
        self.assertTrue(is_even(0))
        self.assertTrue(is_even(2))
        self.assertTrue(is_even(10))
        self.assertTrue(is_even(-4))
    
    def test_is_even_false(self) -> None:
        """测试偶数判断（假）"""
        self.assertFalse(is_even(1))
        self.assertFalse(is_even(3))
        self.assertFalse(is_even(-5))
    
    def test_is_prime_true(self) -> None:
        """测试质数判断（真）"""
        self.assertTrue(is_prime(2))
        self.assertTrue(is_prime(3))
        self.assertTrue(is_prime(5))
        self.assertTrue(is_prime(7))
        self.assertTrue(is_prime(11))
        self.assertTrue(is_prime(13))
    
    def test_is_prime_false(self) -> None:
        """测试质数判断（假）"""
        self.assertFalse(is_prime(0))
        self.assertFalse(is_prime(1))
        self.assertFalse(is_prime(4))
        self.assertFalse(is_prime(6))
        self.assertFalse(is_prime(8))
        self.assertFalse(is_prime(9))
        self.assertFalse(is_prime(10))
    
    def test_fibonacci_single(self) -> None:
        """测试生成单个斐波那契数"""
        result: list[int] = fibonacci(1)
        self.assertEqual(result, [0])
    
    def test_fibonacci_double(self) -> None:
        """测试生成两个斐波那契数"""
        result: list[int] = fibonacci(2)
        self.assertEqual(result, [0, 1])
    
    def test_fibonacci_multiple(self) -> None:
        """测试生成多个斐波那契数"""
        result: list[int] = fibonacci(10)
        expected: list[int] = [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
        self.assertEqual(result, expected)
    
    def test_fibonacci_invalid(self) -> None:
        """测试无效输入"""
        with self.assertRaises(ValueError):
            fibonacci(0)
        
        with self.assertRaises(ValueError):
            fibonacci(-5)
    
    def test_calculate_average_valid(self) -> None:
        """测试计算平均值（有效输入）"""
        numbers: list[int] = [1, 2, 3, 4, 5]
        result: float = calculate_average(numbers)
        self.assertEqual(result, 3.0)
    
    def test_calculate_average_float(self) -> None:
        """测试计算浮点数平均值"""
        numbers: list[float] = [1.5, 2.5, 3.5]
        result: float = calculate_average(numbers)
        self.assertEqual(result, 2.5)
    
    def test_calculate_average_empty(self) -> None:
        """测试计算空列表的平均值"""
        result = calculate_average([])
        self.assertIsNone(result)
    
    def test_find_max_min_valid(self) -> None:
        """测试查找最大值和最小值（有效输入）"""
        numbers: list[int] = [5, 2, 8, 1, 9, 3]
        max_val, min_val = find_max_min(numbers)
        self.assertEqual(max_val, 9.0)
        self.assertEqual(min_val, 1.0)
    
    def test_find_max_min_float(self) -> None:
        """测试查找浮点数的最大值和最小值"""
        numbers: list[float] = [3.5, 1.2, 8.9, 2.1]
        max_val, min_val = find_max_min(numbers)
        self.assertEqual(max_val, 8.9)
        self.assertEqual(min_val, 1.2)
    
    def test_find_max_min_empty(self) -> None:
        """测试查找空列表的最大值和最小值"""
        max_val, min_val = find_max_min([])
        self.assertIsNone(max_val)
        self.assertIsNone(min_val)
    
    def test_celsius_to_fahrenheit(self) -> None:
        """测试摄氏度转华氏度"""
        # 水的冰点
        result: float = celsius_to_fahrenheit(0)
        self.assertEqual(result, 32.0)
        
        # 水的沸点
        result = celsius_to_fahrenheit(100)
        self.assertEqual(result, 212.0)
        
        # 室温
        result = celsius_to_fahrenheit(20)
        self.assertEqual(result, 68.0)
    
    def test_fahrenheit_to_celsius(self) -> None:
        """测试华氏度转摄氏度"""
        # 水的冰点
        result: float = fahrenheit_to_celsius(32)
        self.assertEqual(result, 0.0)
        
        # 水的沸点
        result = fahrenheit_to_celsius(212)
        self.assertEqual(result, 100.0)
        
        # 室温
        result = fahrenheit_to_celsius(68)
        self.assertEqual(result, 20.0)
    
    def test_temperature_conversion_round_trip(self) -> None:
        """测试温度转换往返"""
        celsius: float = 25.0
        fahrenheit: float = celsius_to_fahrenheit(celsius)
        celsius_back: float = fahrenheit_to_celsius(fahrenheit)
        
        # 允许小的浮点误差
        self.assertAlmostEqual(celsius, celsius_back, places=5)


def run_tests() -> NoReturn:
    """运行所有测试"""
    print("🧪 开始运行工具函数测试...")
    unittest.main(verbosity=2, exit=False)


if __name__ == "__main__":
    run_tests()