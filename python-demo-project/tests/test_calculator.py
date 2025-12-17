
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
计算器模块测试
"""

import sys
import unittest
from typing import NoReturn

# 添加 src 目录到路径
sys.path.insert(0, '../src')

from calculator import Calculator


class TestCalculator(unittest.TestCase):
    """计算器测试类"""
    
    def setUp(self) -> None:
        """在每个测试方法之前运行"""
        self.calc: Calculator = Calculator()
    
    def tearDown(self) -> None:
        """在每个测试方法之后运行"""
        pass
    
    def test_add_positive_numbers(self) -> None:
        """测试正数加法"""
        result: float = self.calc.add(5, 3)
        self.assertEqual(result, 8.0)
        self.assertEqual(self.calc.get_operation_count(), 1)
        self.assertEqual(self.calc.get_last_result(), 8.0)
    
    def test_add_negative_numbers(self) -> None:
        """测试负数加法"""
        result: float = self.calc.add(-5, -3)
        self.assertEqual(result, -8.0)
    
    def test_add_mixed_numbers(self) -> None:
        """测试混合数字加法"""
        result: float = self.calc.add(5, -3)
        self.assertEqual(result, 2.0)
    
    def test_add_float_numbers(self) -> None:
        """测试浮点数加法"""
        result: float = self.calc.add(5.5, 2.5)
        self.assertEqual(result, 8.0)
    
    def test_subtract_positive_numbers(self) -> None:
        """测试正数减法"""
        result: float = self.calc.subtract(10, 4)
        self.assertEqual(result, 6.0)
    
    def test_subtract_negative_numbers(self) -> None:
        """测试负数减法"""
        result: float = self.calc.subtract(-5, -3)
        self.assertEqual(result, -2.0)
    
    def test_subtract_mixed_numbers(self) -> None:
        """测试混合数字减法"""
        result: float = self.calc.subtract(5, -3)
        self.assertEqual(result, 8.0)
    
    def test_multiply_positive_numbers(self) -> None:
        """测试正数乘法"""
        result: float = self.calc.multiply(6, 7)
        self.assertEqual(result, 42.0)
    
    def test_multiply_by_zero(self) -> None:
        """测试乘以零"""
        result: float = self.calc.multiply(5, 0)
        self.assertEqual(result, 0.0)
    
    def test_multiply_negative_numbers(self) -> None:
        """测试负数乘法"""
        result: float = self.calc.multiply(-5, 3)
        self.assertEqual(result, -15.0)
    
    def test_divide_positive_numbers(self) -> None:
        """测试正数除法"""
        result: float = self.calc.divide(10, 2)
        self.assertEqual(result, 5.0)
    
    def test_divide_float_result(self) -> None:
        """测试浮点数结果除法"""
        result: float = self.calc.divide(5, 2)
        self.assertEqual(result, 2.5)
    
    def test_divide_by_zero(self) -> None:
        """测试除以零"""
        with self.assertRaises(ZeroDivisionError):
            self.calc.divide(10, 0)
    
    def test_operation_count(self) -> None:
        """测试操作计数"""
        self.assertEqual(self.calc.get_operation_count(), 0)
        
        self.calc.add(1, 2)
        self.assertEqual(self.calc.get_operation_count(), 1)
        
        self.calc.subtract(5, 3)
        self.assertEqual(self.calc.get_operation_count(), 2)
        
        self.calc.multiply(2, 3)
        self.assertEqual(self.calc.get_operation_count(), 3)
    
    def test_last_result(self) -> None:
        """测试最后结果"""
        self.assertIsNone(self.calc.get_last_result())
        
        result1: float = self.calc.add(5, 3)
        self.assertEqual(self.calc.get_last_result(), result1)
        
        result2: float = self.calc.subtract(10, 4)
        self.assertEqual(self.calc.get_last_result(), result2)
    
    def test_reset(self) -> None:
        """测试重置功能"""
        self.calc.add(5, 3)
        self.calc.subtract(10, 4)
        
        self.assertEqual(self.calc.get_operation_count(), 2)
        self.assertIsNotNone(self.calc.get_last_result())
        
        self.calc.reset()
        
        self.assertEqual(self.calc.get_operation_count(), 0)
        self.assertIsNone(self.calc.get_last_result())
    
    def test_string_representation(self) -> None:
        """测试字符串表示"""
        self.calc.add(5, 3)
        expected_str: str = "Calculator(operations=1, last_result=8.0)"
        self.assertEqual(str(self.calc), expected_str)
    
    def test_repr_representation(self) -> None:
        """测试详细表示"""
        self.calc.add(5, 3)
        expected_repr: str = "Calculator(operation_count=1, last_result=8.0)"
        self.assertEqual(repr(self.calc), expected_repr)


def run_tests() -> NoReturn:
    """运行所有测试"""
    print("🧪 开始运行计算器测试...")
    unittest.main(verbosity=2, exit=False)


if __name__ == "__main__":
    run_tests()